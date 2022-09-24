package godirstat_test

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Gewad/GoDirStat/godirstat"
	"github.com/Gewad/GoDirStat/godirstat/godirstat_test_utility"
	"github.com/pkg/errors"
)

func TestGetDirInfo(t *testing.T) {
	config := godirstat.Config{
		FnReadDir: godirstat_test_utility.ReadDir,
	}

	root := godirstat.GetDirInfoWithConfig("1", 0, config)

	expRoot := godirstat.DirInfo{
		Name: "1",
		Size: 3,
		Files: []godirstat.FileInfo{
			{
				Name: "f1",
				Size: 1,
			},
		},
		TotalFiles: 2,
		Dirs: []godirstat.DirInfo{
			{
				Name: "2",
				Size: 2,
				Files: []godirstat.FileInfo{
					{
						Name: "f2",
						Size: 1,
					},
				},
				TotalFiles: 1,
			},
		},
		TotalDirs: 1,
	}

	if !reflect.DeepEqual(root, expRoot) {
		t.Fatalf("expected root to match %+v, got %+v", expRoot, root)
	}
}

func TestReplaceRootName(t *testing.T) {
	config := godirstat.Config{
		FnReadDir: godirstat_test_utility.ReadDir,
	}

	root := godirstat.GetDirInfoWithConfig(".", 0, config)

	wd, err := os.Getwd()
	if err != nil {
		t.Error(errors.Wrap(err, "could not retrieve name of current working directory"))
	}
	if !strings.Contains(wd, root.Name) {
		t.Fatalf("name of root did not match working directory: %s", root.Name)
	}
}

func TestNonExistentFile(t *testing.T) {
	config := godirstat.Config{
		FnReadDir: godirstat_test_utility.ReadDir,
	}

	root := godirstat.GetDirInfoWithConfig("non-existent", 0, config)

	if !reflect.DeepEqual(root, godirstat.DirInfo{}) {
		t.Fatalf("expected empty DirInfo to be returned due to error, got %+v", root)
	}
}
