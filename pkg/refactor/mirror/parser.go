package mirror

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/direktiv/direktiv/pkg/model"
	"github.com/direktiv/direktiv/pkg/refactor/api"
	"github.com/direktiv/direktiv/pkg/refactor/core"
	"github.com/direktiv/direktiv/pkg/refactor/filestore"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

type Parser struct {
	log     FormatLogger
	matcher gitignore.Matcher
	src     Source
	tempDir string

	Filters   map[string]*api.Filter
	Services  map[string]*api.Service
	Workflows map[string][]byte

	DeprecatedNamespaceVars map[string][]byte
	DeprecatedWorkflowVars  map[string]map[string][]byte
}

func NewParser(log FormatLogger, src Source) (*Parser, error) {
	tempDir, err := os.MkdirTemp("", "direktiv_sync_*")
	if err != nil {
		return nil, err
	}
	log.Debugf("Processing repository in temporary directory: %s", tempDir)

	p := &Parser{
		log:     log,
		matcher: gitignore.NewMatcher(nil),
		src:     src,
		tempDir: tempDir,

		Filters:   make(map[string]*api.Filter),
		Services:  make(map[string]*api.Service),
		Workflows: make(map[string][]byte),

		DeprecatedNamespaceVars: make(map[string][]byte),
		DeprecatedWorkflowVars:  make(map[string]map[string][]byte),
	}

	err = p.parse()
	if err != nil {
		log.Errorf("Error processing repository: %v", err)
		_ = p.Close()
		return nil, err
	}

	return p, nil
}

func (p *Parser) Close() error {
	return os.RemoveAll(p.tempDir)
}

func (p *Parser) parse() error {
	err := p.loadIgnores()
	if err != nil {
		return err
	}

	err = p.filterCopySource()
	if err != nil {
		return err
	}

	err = p.scanAndPruneDirektivResourceFiles()
	if err != nil {
		return err
	}

	err = p.scanAndPruneAmbiguousDirektivWorkflowFiles()
	if err != nil {
		return err
	}

	err = p.parseDeprecatedVariableFiles()
	if err != nil {
		return err
	}

	err = p.logUnprunedFiles()
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) loadIgnores() error {
	f, err := p.src.FS().Open(".direktivignore")
	if errors.Is(err, os.ErrNotExist) {
		p.log.Infof("No .direktivignore file detected")
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to open direktivignore file: %w", err)
	}
	defer f.Close()

	var ps []gitignore.Pattern
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if !strings.HasPrefix(s, "#") && len(strings.TrimSpace(s)) > 0 {
			fmt.Println("~", s)
			ps = append(ps, gitignore.ParsePattern(s, nil))
		}
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("failed to close direktivignore file: %w", err)
	}

	p.matcher = gitignore.NewMatcher(ps)

	return nil
}

func (p *Parser) filterCopySourceWalkFunc(path string, d fs.DirEntry, err error) error {
	isMatch := p.matcher.Match(strings.Split(path, "/"), d.IsDir())
	if isMatch {
		if d.IsDir() {
			p.log.Infof("Skipping directory '%s': excluded by .direktivignore patterns", path)
			return fs.SkipDir
		}

		p.log.Infof("Skipping file '%s': excluded by .direktivignore patterns", path)
		return nil
	}

	base := filepath.Base(path)
	_, err = filestore.SanitizePath(base)
	if err != nil {
		if d.IsDir() {
			p.log.Infof("Skipping directory '%s': filename contains illegal characters", path)
			return fs.SkipDir
		}

		p.log.Infof("Skipping file '%s': filename contains illegal characters", path)
		return nil
	}

	tpath := filepath.Join(p.tempDir, path)

	if d.IsDir() {
		err := os.MkdirAll(tpath, 0755)
		if err != nil {
			return err
		}

		p.log.Debugf("Created directory '%s'", path)
		return nil
	}

	// NOTE: duplicating the file here isn't strictly necessary and could cause problems,
	// 	but large file sizes are a problem anyway.
	src, err := p.src.FS().Open(path)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(tpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	err = src.Close()
	if err != nil {
		return err
	}

	err = dst.Close()
	if err != nil {
		return err
	}

	p.log.Debugf("Created file '%s'", path)

	return nil
}

func (p *Parser) filterCopySource() error {
	err := fs.WalkDir(p.src.FS(), ".", p.filterCopySourceWalkFunc)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) listYAMLFiles() ([]string, error) {
	var paths []string

	tfs := os.DirFS(p.tempDir)

	err := fs.WalkDir(tfs, ".", func(path string, d fs.DirEntry, err error) error {
		if !d.Type().IsRegular() {
			return nil
		}

		if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
			paths = append(paths, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (p *Parser) scanAndPruneDirektivResourceFile(path string) error {
	tpath := filepath.Join(p.tempDir, path)

	data, err := os.ReadFile(tpath)
	if err != nil {
		return err
	}

	resource, err := api.LoadResource(data)
	if errors.Is(err, api.ErrNotDirektivAPIResource) {
		return nil
	}
	if err != nil {
		p.log.Warnf("Error loading possible Direktiv resource definition '%s': %v", path, err)
		return nil
	}

	switch typ := resource.(type) {
	case *api.Filters:
		filters := resource.(*api.Filters)
		err = p.handleFilters(path, filters)
		if err != nil {
			return err
		}
	case *api.Services:
		services := resource.(*api.Services)
		err = p.handleServices(path, services)
		if err != nil {
			return err
		}
	case *model.Workflow:
		err = p.handleWorkflow(path, data)
		if err != nil {
			return err
		}
	default:
		panic(typ)
	}

	err = os.Remove(tpath)
	if err != nil {
		return err
	}

	p.log.Debugf("Pruned Direktiv resource file '%s'", path)

	return nil
}

func (p *Parser) scanAndPruneDirektivResourceFiles() error {
	paths, err := p.listYAMLFiles()
	if err != nil {
		return err
	}

	for _, path := range paths {
		err = p.scanAndPruneDirektivResourceFile(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) scanAndPruneAmbiguousDirektivWorkflowFile(path string) error {
	tpath := filepath.Join(p.tempDir, path)

	data, err := os.ReadFile(tpath)
	if err != nil {
		return err
	}

	wf := new(model.Workflow)
	err = wf.Load(data)
	if err != nil {
		p.log.Warnf("Error loading possible Direktiv workflow definition (ambiguous) '%s': %v", path, err)
		return nil
	}

	err = p.handleWorkflow(path, data)
	if err != nil {
		return err
	}

	err = os.Remove(tpath)
	if err != nil {
		return err
	}

	p.log.Debugf("Pruned Direktiv workflow definition file '%s'", path)

	return nil
}

func (p *Parser) scanAndPruneAmbiguousDirektivWorkflowFiles() error {
	paths, err := p.listYAMLFiles()
	if err != nil {
		return err
	}

	for _, path := range paths {
		err = p.scanAndPruneAmbiguousDirektivWorkflowFile(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) handleWorkflow(path string, data []byte) error {
	p.log.Infof("Direktiv resource file containing a workflow definition found at '%s'", path)

	p.Workflows[path] = data

	return nil
}

func (p *Parser) handleFilters(path string, filters *api.Filters) error {
	p.log.Infof("Direktiv resource file containing %d filter definitions found at '%s'", len(filters.Filters), path)

	for idx, filter := range filters.Filters {
		if _, exists := p.Filters[filter.Name]; exists {
			return fmt.Errorf("duplicate definition detected for filter '%s'", filter.Name)
		}
		p.Filters[filter.Name] = &filters.Filters[idx]
		p.log.Infof("Filter '%s' loaded.", filter.Name)
	}

	return nil
}

func (p *Parser) handleServices(path string, services *api.Services) error {
	p.log.Infof("Direktiv resource file containing %d service definitions found at '%s'", len(services.Services), path)

	for idx, service := range services.Services {
		if _, exists := p.Services[service.Name]; exists {
			return fmt.Errorf("duplicate definition detected for service '%s'", service.Name)
		}
		p.Services[service.Name] = &services.Services[idx]
		p.log.Infof("Servcice '%s' loaded.", service.Name)
	}

	return nil
}

func (p *Parser) logUnprunedFiles() error {
	tfs := os.DirFS(p.tempDir)

	err := fs.WalkDir(tfs, ".", func(path string, d fs.DirEntry, err error) error {
		if !d.Type().IsRegular() {
			return nil
		}

		p.log.Infof("File '%s' loaded.", path)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) ListFiles() ([]string, error) {
	var paths []string

	tfs := os.DirFS(p.tempDir)

	err := fs.WalkDir(tfs, ".", func(path string, d fs.DirEntry, err error) error {
		paths = append(paths, path)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (p *Parser) parseDeprecatedVariableFiles() error {
	regex := regexp.MustCompile(core.RuntimeVariableNameRegexPattern)

	allFiles, err := p.ListFiles()
	if err != nil {
		return err
	}

	var trimmed []string
	for _, fpath := range allFiles {
		actual := filepath.Join(p.tempDir, fpath)
		fi, err := os.Stat(actual)
		if err != nil {
			return err
		}

		if fi.Mode().IsRegular() {
			trimmed = append(trimmed, fpath)
		}
	}

	allFiles = trimmed

	for _, fpath := range allFiles {
		base := filepath.Base(fpath)
		prefix := "var."
		vname := strings.TrimPrefix(base, prefix)

		if !regex.MatchString(vname) {
			p.log.Warnf("Detected a possible deprecated namespace variable definition with an invalid name at: %s", fpath)

			continue
		}

		if strings.HasPrefix(base, prefix) {
			actual := filepath.Join(p.tempDir, fpath)

			data, err := os.ReadFile(actual)
			if err != nil {
				return err
			}

			p.DeprecatedNamespaceVars[vname] = data

			p.log.Warnf("Detected deprecated namespace variable definition at: %s", fpath)
		}
	}

	var allWorkflows []string
	for k := range p.Workflows {
		allWorkflows = append(allWorkflows, k)
	}

	for _, fpath := range allFiles {
		for _, wpath := range allWorkflows {
			prefix := wpath + "."
			vname := strings.TrimPrefix(fpath, prefix)
			if !regex.MatchString(vname) {
				p.log.Warnf("Detected a possible deprecated workflow variable definition with an invalid name at: %s", fpath)

				continue
			}

			if strings.HasPrefix(fpath, prefix) {
				actual := filepath.Join(p.tempDir, fpath)

				data, err := os.ReadFile(actual)
				if err != nil {
					return err
				}

				m, exists := p.DeprecatedWorkflowVars[wpath]
				if !exists {
					m = make(map[string][]byte)
					p.DeprecatedWorkflowVars[wpath] = m
				}

				m[vname] = data

				p.log.Warnf("Detected deprecated workflow variable definition at: %s", fpath)
			}
		}
	}

	return nil
}
