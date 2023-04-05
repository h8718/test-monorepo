package psql_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/direktiv/direktiv/pkg/refactor/utils"

	"github.com/direktiv/direktiv/pkg/refactor/filestore"
	"github.com/direktiv/direktiv/pkg/refactor/filestore/psql"
	"github.com/google/uuid"
)

func TestRoot_CreateFile(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := psql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.UUID{})
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	tests := []struct {
		path    string
		typ     string
		payload string
	}{
		{"/example.text", "text", "abcd"},
		{"/example1.text", "text", "abcd"},
		{"/example2.text", "text", "abcd"},
	}
	for _, tt := range tests {
		t.Run("valid", func(t *testing.T) {
			assertRootCorrectFileCreation(t, fs, root, tt.path)
		})
	}
}

func assertRootCorrectFileCreation(t *testing.T, fs filestore.FileStore, root *filestore.Root, path string) {
	t.Helper()

	var data []byte
	typ := "directory"
	if strings.Contains(path, ".text") {
		data = []byte("some data")
		typ = "text"
	}

	assertRootCorrectFileCreationWithContent(t, fs, root, path, typ, data)
}

func assertRootCorrectFileCreationWithContent(t *testing.T, fs filestore.FileStore, root *filestore.Root, path string, typ string, data []byte) {
	t.Helper()

	file, _, err := fs.ForRootID(root.ID).CreateFile(context.Background(), path, filestore.FileType(typ), bytes.NewReader(data))
	if err != nil {
		t.Errorf("unexpected CreateFile() error: %v", err)
		return
	}
	if file == nil {
		t.Errorf("unexpected nil file CreateFile()")
		return
	}
	if file.Path != path {
		t.Errorf("unexpected file.Path, got: >%s<, want: >%s<", file.Path, path)
		return
	}

	if typ != "directory" {
		reader, _ := fs.ForFile(file).GetData(context.Background())
		createdData, _ := io.ReadAll(reader)
		if string(createdData) != string(data) {
			t.Errorf("unexpected GetPath(), got: >%s<, want: >%s<", createdData, data)
			return
		}
	}

	file, err = fs.ForRootID(root.ID).GetFile(context.Background(), path)
	if err != nil {
		t.Errorf("unexpected GetFile() error: %v", err)
		return
	}
	if file == nil {
		t.Errorf("unexpected nil file GetFile()")
		return
	}
	if file.Path != path {
		t.Errorf("unexpected file.Path, got: >%s<, want: >%s<", file.Path, path)
		return
	}
}

func TestRoot_CorrectReadDirectory(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := psql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.New())
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	// Test root directory:
	{
		assertRootCorrectFileCreation(t, fs, root, "/file1.text")
		assertRootCorrectFileCreation(t, fs, root, "/file2.text")

		assertRootFilesInPath(t, fs, root, "/",
			"/file1.text",
			"/file2.text",
		)
	}

	// Add /dir1 directory:
	{
		assertRootCorrectFileCreation(t, fs, root, "/dir1")
		assertRootCorrectFileCreation(t, fs, root, "/dir1/file3.text")
		assertRootCorrectFileCreation(t, fs, root, "/dir1/file4.text")

		assertRootFilesInPath(t, fs, root, "/dir1",
			"/dir1/file3.text",
			"/dir1/file4.text",
		)
		assertRootFilesInPath(t, fs, root, "/",
			"/file1.text",
			"/file2.text",
			"/dir1",
		)
	}

	// Add /dir1/dir2 directory:
	{
		assertRootCorrectFileCreation(t, fs, root, "/dir1/dir2")
		assertRootCorrectFileCreation(t, fs, root, "/dir1/dir2/file5.text")
		assertRootCorrectFileCreation(t, fs, root, "/dir1/dir2/file6.text")

		assertRootFilesInPath(t, fs, root, "/dir1/dir2",
			"/dir1/dir2/file5.text",
			"/dir1/dir2/file6.text",
		)
		assertRootFilesInPath(t, fs, root, "/dir1",
			"/dir1/file3.text",
			"/dir1/file4.text",
			"/dir1/dir2",
		)
		assertRootFilesInPath(t, fs, root, "/",
			"/file1.text",
			"/file2.text",
			"/dir1",
		)
	}
}

func TestRoot_CalculateChecksumDirectory(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := psql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.New())
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	filestore.DefaultCalculateChecksum = func(data []byte) []byte {
		return []byte(fmt.Sprintf("---%s---", data))
	}

	// Test root directory:
	{
		assertRootCorrectFileCreationWithContent(t, fs, root, "/file1.text", "text", []byte("content1"))
		assertRootCorrectFileCreationWithContent(t, fs, root, "/file2.text", "text", []byte("content2"))
		assertRootCorrectFileCreationWithContent(t, fs, root, "/dir1", "directory", nil)
		assertRootCorrectFileCreationWithContent(t, fs, root, "/empty_dir", "directory", nil)
		assertRootCorrectFileCreationWithContent(t, fs, root, "/dir1/file3.text", "text", []byte("content3"))
		assertRootCorrectFileCreationWithContent(t, fs, root, "/dir1/file4.text", "text", []byte("content4"))

		assertChecksums(t, fs, root,
			"/file1.text", "---content1---",
			"/file2.text", "---content2---",
			"/dir1", "",
			"/empty_dir", "",
			"/dir1/file3.text", "---content3---",
			"/dir1/file4.text", "---content4---",
		)
	}
}

func assertRootFilesInPath(t *testing.T, fs filestore.FileStore, root *filestore.Root, searchPath string, paths ...string) {
	t.Helper()

	files, err := fs.ForRootID(root.ID).ReadDirectory(context.Background(), searchPath)
	if err != nil {
		t.Errorf("unepxected ReadDirectory() error = %v", err)
		return
	}
	if len(files) != len(paths) {
		t.Errorf("unexpected ReadDirectory() length, got: %d, want: %d", len(files), len(paths))
		return
	}

	for i := range paths {
		if files[i].Path != paths[i] {
			t.Errorf("unexpected files[%d].Path , got: >%s<, want: >%s<", i, files[i].Path, paths[i])
			return
		}
	}
}

func assertChecksums(t *testing.T, fs filestore.FileStore, root *filestore.Root, paths ...string) {
	t.Helper()

	checksumsMap, err := fs.ForRootID(root.ID).CalculateChecksumsMap(context.Background())
	if err != nil {
		t.Errorf("unepxected CalculateChecksumsMap() error = %v", err)
	}
	if len(checksumsMap)*2 != len(paths) {
		t.Errorf("unexpected CalculateChecksumsMap() length, got: %d, want: %d", len(checksumsMap), len(paths)/2)
	}

	wantChecksumsMap := make(map[string]string)

	for i := 0; i < len(paths)-1; i += 2 {
		wantChecksumsMap[paths[i]] = paths[i+1]
	}

	if !reflect.DeepEqual(checksumsMap, wantChecksumsMap) {
		t.Errorf("unexpected CalculateChecksumsMap() result, got: %v, want: %v", checksumsMap, wantChecksumsMap)
	}
}
