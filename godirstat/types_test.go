package godirstat_test

import (
	"reflect"
	"testing"

	"github.com/Gewad/GoDirStat/godirstat"
)

func TestSortDirsAndFiles(t *testing.T) {
	t.Parallel()

	dir := godirstat.DirInfo{
		Dirs: []godirstat.DirInfo{
			{
				Size: 2,
			},
			{
				Size: 1,
			},
			{
				Size: 3,
			},
		},
		Files: []godirstat.FileInfo{
			{
				Size: 2,
			},
			{
				Size: 1,
			},
			{
				Size: 3,
			},
		},
	}

	expDir := godirstat.DirInfo{
		Dirs: []godirstat.DirInfo{
			{
				Size: 3,
			},
			{
				Size: 2,
			},
			{
				Size: 1,
			},
		},
		Files: []godirstat.FileInfo{
			{
				Size: 3,
			},
			{
				Size: 2,
			},
			{
				Size: 1,
			},
		},
	}

	dir.SortDirsAndFiles()

	if !reflect.DeepEqual(dir, expDir) {
		t.Fatalf("expected correctly sorted DirInfo, got: %+v", dir)
	}
}
