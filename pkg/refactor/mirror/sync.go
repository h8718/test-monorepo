package mirror

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/direktiv/direktiv/pkg/refactor/filestore"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	ignore "github.com/sabhiram/go-gitignore"
	"go.uber.org/zap"
)

var WorkflowConfigHook func(context.Context, filestore.FileStore, Store, uuid.UUID, *filestore.File) error

// TODO: implement parsing direktiv variables.
// TODO: check %w verb on errors.
// TODO: fix errors and add logs.
// TODO: implement a mechanism to clean dangling processes and cleaning them up.
// TODO: implement synchronizing jobs.

type mirroringJob struct {
	// job parameters.

	//nolint:containedctx
	ctx context.Context
	lg  *zap.SugaredLogger

	// job artifacts.
	err            error
	distDirectory  string
	sourcedPaths   []string
	direktivIgnore *ignore.GitIgnore
	rootChecksums  map[string]string

	workflowIDs map[string]uuid.UUID
}

func (j *mirroringJob) SetProcessStatus(store Store, process *Process, status string) *mirroringJob {
	if j.err != nil {
		return j
	}
	var err error

	process.Status = status
	_, err = store.UpdateProcess(j.ctx, process)

	if err != nil {
		j.err = fmt.Errorf("updating process state, err: %w", err)
	}

	return j
}

func (j *mirroringJob) CreateTempDirectory() *mirroringJob {
	if j.err != nil {
		return j
	}
	var err error

	j.distDirectory, j.err = os.MkdirTemp("", "direktiv_mirrors")
	if err != nil {
		j.err = fmt.Errorf("create mirror dst directory, err: %w", err)
	}

	j.lg.Infow("creating mirroring temp dir", "dir", j.distDirectory)

	return j
}

func (j *mirroringJob) DeleteTempDirectory() *mirroringJob {
	if j.err != nil {
		return j
	}

	err := os.RemoveAll(j.distDirectory)
	if err != nil {
		j.err = fmt.Errorf("os remove dist directory, dir: %s, err: %w", j.distDirectory, err)
	}

	j.lg.Infow("deleting mirroring temp dir", "dir", j.distDirectory)

	return j
}

func (j *mirroringJob) PullSourceInPath(source Source, config *Config) *mirroringJob {
	if j.err != nil {
		return j
	}

	err := source.PullInPath(config, j.distDirectory)
	if err != nil {
		j.err = fmt.Errorf("pulling source in path, path:%s, err: %w", j.distDirectory, err)
	}

	return j
}

func (j *mirroringJob) CreateSourceFilesList() *mirroringJob {
	if j.err != nil {
		return j
	}
	var err error

	paths := []string{}

	err = filepath.WalkDir(j.distDirectory, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		relativePath := strings.TrimPrefix(path, j.distDirectory)

		if strings.Contains(relativePath, ".git") {
			return nil
		}
		paths = append(paths, relativePath)

		return nil
	})
	if err != nil {
		j.err = fmt.Errorf("walking dist directory, path: %s, err: %w", j.distDirectory, err)

		return j
	}
	j.sourcedPaths = paths

	for _, p := range j.sourcedPaths {
		j.lg.Infow("source path", "path", p)
	}

	return j
}

func (j *mirroringJob) ParseIgnoreFile(ignorePath string) *mirroringJob {
	if j.err != nil {
		return j
	}
	var err error

	absoulteIgnoreFilePath := filepath.Join(j.distDirectory, ignorePath)

	// first we check if ignore file exists.
	ignoreFileExist := false
	for _, path := range j.sourcedPaths {
		if path == absoulteIgnoreFilePath {
			ignoreFileExist = true

			break
		}
	}
	if !ignoreFileExist {
		// don't parse as there is not ignore file.
		return j
	}

	j.direktivIgnore, err = ignore.CompileIgnoreFile(absoulteIgnoreFilePath)
	if err != nil {
		j.err = fmt.Errorf("parsing ignore file, path: %s, err: %w", absoulteIgnoreFilePath, err)
	}

	return j
}

func (j *mirroringJob) FilterIgnoredFiles() *mirroringJob {
	if j.err != nil {
		return j
	}

	// Ignore was not parsed because ignore file is not present.
	if j.direktivIgnore == nil {
		return j
	}

	skippedList := []string{}
	for _, path := range j.sourcedPaths {
		if !j.direktivIgnore.MatchesPath(path) {
			skippedList = append(skippedList, path)
		}
	}
	j.sourcedPaths = skippedList

	return j
}

func (j *mirroringJob) ParseDirektivVars(store Store, namespaceID uuid.UUID) *mirroringJob {
	if j.err != nil {
		return j
	}

	namespaceVarKeys, workflowVarKeys := parseDirektivVars(j.sourcedPaths)

	for _, pk := range namespaceVarKeys {
		path := j.distDirectory + pk[0] + "." + pk[1]
		data, err := os.ReadFile(path)
		if err != nil {
			j.err = fmt.Errorf("read os file, path: %s, err: %w", path, err)

			return j
		}
		mType := mimetype.Detect(data)
		hash, err := computeHash(data)
		if err != nil {
			j.err = fmt.Errorf("calculate hash string, path: %s, err: %w", path, err)

			return j
		}
		err = store.SetNamespaceVariable(j.ctx, namespaceID, pk[1], data, hash, mType.String())
		if err != nil {
			j.err = fmt.Errorf("save namespace variable, path: %s, err: %w", path, err)

			return j
		}
	}

	for _, pk := range workflowVarKeys {
		path := j.distDirectory + pk[0] + "." + pk[1]
		workflowID, ok := j.workflowIDs[pk[0]]
		if !ok {
			continue
		}

		data, err := os.ReadFile(path)
		if err != nil {
			j.err = fmt.Errorf("read os file, path: %s, err: %w", path, err)

			return j
		}
		mType := mimetype.Detect(data)
		hash, err := computeHash(data)
		if err != nil {
			j.err = fmt.Errorf("calculate hash string, path: %s, err: %w", path, err)

			return j
		}
		err = store.SetWorkflowVariable(j.ctx, workflowID, pk[1], data, hash, mType.String())
		if err != nil {
			j.err = fmt.Errorf("save namespace variable, path: %s, err: %w", path, err)

			return j
		}
	}

	return j
}

func (j *mirroringJob) CopyFilesToRoot(fStore filestore.FileStore, store Store, namespaceID uuid.UUID) *mirroringJob {
	if j.err != nil {
		return j
	}

	for _, path := range j.sourcedPaths {
		j.lg.Infow("trying to copy", "path", path)
		data, err := os.ReadFile(j.distDirectory + path)
		if err != nil {
			j.err = fmt.Errorf("read os file, path: %s, err: %w", path, err)

			return j
		}
		checksum := string(filestore.DefaultCalculateChecksum(data))
		fileChecksum, pathDoesExist := j.rootChecksums[path]
		isEqualChecksum := checksum == fileChecksum

		// TODO: yassir, check for workflowID.

		if pathDoesExist && isEqualChecksum {
			j.lg.Infow("checksum skipped to root", "path", path)

			continue
		}

		fileReader := bytes.NewReader(data)

		if !pathDoesExist {
			typ := filestore.FileTypeFile
			if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
				typ = filestore.FileTypeWorkflow
			}
			newFile, _, err := fStore.ForRootID(namespaceID).CreateFile(j.ctx, path, typ, fileReader)
			if err != nil {
				j.err = fmt.Errorf("filestore create file, path: %s, err: %w", path, err)

				return j
			}
			j.lg.Infow("copied to root", "path", path)

			if typ == filestore.FileTypeWorkflow {
				j.workflowIDs[path] = newFile.ID
			}

			err = WorkflowConfigHook(j.ctx, fStore, store, namespaceID, newFile)
			if err != nil {
				j.err = fmt.Errorf("filestore configure workflow, path: %s, err: %w", path, err)
				return j
			}

			continue
		}

		file, err := fStore.ForRootID(namespaceID).GetFile(j.ctx, path)
		if err != nil {
			j.err = fmt.Errorf("get file from root, path: %s, err: %w", path, err)

			return j
		}

		_, err = fStore.ForFile(file).CreateRevision(j.ctx, "", fileReader)
		if err != nil {
			j.err = fmt.Errorf("filestore create revision, path: %s, err: %w", path, err)
			return j
		}
		j.lg.Infow("revision to root", "path", path)

		err = WorkflowConfigHook(j.ctx, fStore, store, namespaceID, file)
		if err != nil {
			j.err = fmt.Errorf("filestore configure workflow, path: %s, err: %w", path, err)
			return j
		}
	}

	return j
}

func (j *mirroringJob) CropFilesAndDirectoriesInRoot(fStore filestore.FileStore, store Store, namespaceID uuid.UUID) *mirroringJob {
	if j.err != nil {
		return j
	}

	err := fStore.ForRootID(namespaceID).CropFilesAndDirectories(j.ctx, j.sourcedPaths)
	if err != nil {
		j.err = fmt.Errorf("filestore crop to paths, err: %w", err)

		return j
	}

	return j
}

func (j *mirroringJob) ReadRootFilesChecksums(fStore filestore.FileStore, namespaceID uuid.UUID) *mirroringJob {
	if j.err != nil {
		return j
	}

	checksums, err := fStore.ForRootID(namespaceID).CalculateChecksumsMap(j.ctx)
	if err != nil {
		j.err = fmt.Errorf("filestore calculate checksums map, err: %w", err)

		return j
	}

	j.rootChecksums = checksums

	return j
}

func (j *mirroringJob) CreateAllDirectories(fStore filestore.FileStore, namespaceID uuid.UUID) *mirroringJob {
	if j.err != nil {
		return j
	}

	createdDirs := map[string]bool{}

	for _, path := range j.sourcedPaths {
		dir := filepath.Dir(path)
		allParentDirs := splitPathToDirectories(dir)
		for _, d := range allParentDirs {
			if _, isExists := j.rootChecksums[d]; isExists {
				continue
			}

			if _, isCreated := createdDirs[d]; isCreated {
				continue
			}

			_, _, err := fStore.ForRootID(namespaceID).CreateFile(j.ctx, d, filestore.FileTypeDirectory, nil)

			// check if it is a fatal error.
			if err != nil && !errors.Is(err, filestore.ErrPathAlreadyExists) {
				j.err = fmt.Errorf("filestore create dir, path: %s, err: %w", d, err)

				return j
			}

			createdDirs[d] = true
		}
	}

	return j
}

func (j *mirroringJob) Error() interface{} {
	return j.err
}

func splitPathToDirectories(dir string) []string {
	list := []string{}

	dir = strings.TrimSpace(dir)
	dir = strings.TrimPrefix(dir, "/")

	parts := strings.Split(dir, "/")

	for i := range parts {
		list = append(list, "/"+strings.Join(parts[:i+1], "/"))
	}

	return list
}

func parseDirektivVars(paths []string) ([][]string, [][]string) {
	pathsMap := map[string]bool{}
	for _, p := range paths {
		pathsMap[p] = true
	}

	namespaceVarPathsKeys := [][]string{}
	workflowVarPathsKeys := [][]string{}

	for _, p := range paths {
		base := filepath.Base(p)
		dir := filepath.Dir(p)

		if strings.Contains(base, "var.") && len(base) > len("var.") {
			if strings.HasPrefix(strings.TrimPrefix(base, "var."), "_") {
				continue
			}
			namespaceVarPathsKeys = append(namespaceVarPathsKeys, []string{filepath.Clean(dir + "/var"), strings.TrimPrefix(base, "var.")})

			continue
		}

		if strings.Contains(base, ".yaml.") {
			if _, ok := pathsMap[p]; !ok {
				continue
			}
			parts := strings.Split(base, ".yaml.")
			//nolint:gomnd
			if len(parts) == 2 {
				workflowVarPathsKeys = append(workflowVarPathsKeys, []string{dir + "/" + parts[0] + ".yaml", parts[1]})
			}
		}
		if strings.Contains(base, ".yml.") {
			if _, ok := pathsMap[p]; !ok {
				continue
			}
			parts := strings.Split(base, ".yml.")
			//nolint:gomnd
			if len(parts) == 2 {
				workflowVarPathsKeys = append(workflowVarPathsKeys, []string{dir + "/" + parts[0] + ".yml", parts[1]})
			}
		}
	}

	return namespaceVarPathsKeys, workflowVarPathsKeys
}

// computeHash is a shortcut to calculate a hash for a byte slice.
func computeHash(data []byte) (string, error) {
	hasher := sha256.New()
	_, err := io.Copy(hasher, bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
