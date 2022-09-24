package godirstat_test_utility

import (
	"errors"
	"io/fs"
	"time"
)

type CustomFileInfo struct {
	name  string
	isDir bool
}

func (fileInfo CustomFileInfo) Name() string {
	return fileInfo.name
}

func (fileInfo CustomFileInfo) Size() int64 {
	return 1
}

func (fileInfo CustomFileInfo) Mode() fs.FileMode {
	return 0
}

func (fileInfo CustomFileInfo) ModTime() time.Time {
	return time.Now()
}

func (fileInfo CustomFileInfo) IsDir() bool {
	return fileInfo.isDir
}

func (fileInfo CustomFileInfo) Sys() any {
	return nil
}

type CustomDirEntry struct {
	fileInfo CustomFileInfo
}

func (dirEntry CustomDirEntry) Name() string {
	return dirEntry.fileInfo.name
}

func (dirEntry CustomDirEntry) IsDir() bool {
	return dirEntry.fileInfo.isDir
}

func (dirEntry CustomDirEntry) Type() fs.FileMode {
	return 0
}

func (dirEntry CustomDirEntry) Info() (fs.FileInfo, error) {
	return dirEntry.fileInfo, nil
}

var dirMap = map[string][]fs.DirEntry{
	"1": {
		CustomDirEntry{
			fileInfo: CustomFileInfo{
				name:  "2",
				isDir: true,
			},
		},
		CustomDirEntry{
			fileInfo: CustomFileInfo{
				name:  "f1",
				isDir: false,
			},
		},
	},
	"1/2": {
		CustomDirEntry{
			fileInfo: CustomFileInfo{
				name:  "f2",
				isDir: false,
			},
		},
	},
	".": {},
}

func ReadDir(dir string) ([]fs.DirEntry, error) {
	dirs := dirMap[dir]

	if dirs == nil {
		return nil, errors.New("can't read dir")
	}

	return dirs, nil
}
