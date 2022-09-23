package godirstat

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

func GetDirInfo(dirString string, size int) DirInfo {
	items, err := ioutil.ReadDir(dirString)
	if err != nil {
		message := dirString + " - " + err.Error()
		fmt.Println(message)
		return DirInfo{}
	}

	dir := DirInfo{
		Name:  getNameFromDir(dirString),
		Size:  size,
		Files: getDirFiles(items),
		Dirs:  getDirDirs(items, dirString)}

	if dir.Name == "." {
		wd, err := os.Getwd()

		if err == nil {
			dir.Name = getNameFromDir(wd)
		}
	}

	dirsSize, filesAmount, dirsAmount := calculateTotals(dir)
	dir.Size += dirsSize
	dir.TotalFiles = filesAmount
	dir.TotalDirs = dirsAmount

	return dir
}

func getDirDirs(items []fs.FileInfo, dirStringPrev string) []DirInfo {
	var dirs []DirInfo

	for _, item := range items {
		if item.IsDir() {
			dirString := dirStringPrev + "/" + item.Name()
			dirs = append(dirs, GetDirInfo(dirString, int(item.Size())))
		}
	}

	return dirs
}

func getDirFiles(items []fs.FileInfo) []FileInfo {
	var files []FileInfo

	for _, item := range items {
		if !item.IsDir() {
			files = append(files, FileInfo{item.Name(), int(item.Size())})
		}
	}

	return files
}

func calculateTotals(dir DirInfo) (int, int, int) {
	totalSize := 0
	totalFiles := 0
	totalDirs := 0

	for _, dir := range dir.Dirs {
		totalSize += int(dir.Size)
		totalDirs += dir.TotalDirs + 1
		totalFiles += dir.TotalFiles
	}

	for _, file := range dir.Files {
		totalSize += int(file.Size)
		totalFiles += 1
	}

	return totalSize, totalFiles, totalDirs
}

func getNameFromDir(dir string) string {
	dirParts := strings.FieldsFunc(dir, func(r rune) bool {
		return r == '/' || r == '\\'
	})
	return dirParts[len(dirParts)-1]
}
