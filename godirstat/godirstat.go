package godirstat

import (
	"fmt"
	"io/fs"
	"io/ioutil"
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
		name:  getNameFromDir(dirString),
		size:  size,
		files: getDirFiles(items),
		dirs:  getDirDirs(items, dirString)}

	dirsSize, filesAmount, dirsAmount := calculateTotals(dir)
	dir.size += dirsSize
	dir.totalFiles = filesAmount
	dir.totalDirs = dirsAmount

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

	for _, dir := range dir.dirs {
		totalSize += int(dir.size)
		totalDirs += dir.totalDirs + 1
		totalFiles += dir.totalFiles
	}

	for _, file := range dir.files {
		totalSize += int(file.size)
		totalFiles += 1
	}

	return totalSize, totalFiles, totalDirs
}

func getNameFromDir(dir string) string {
	dirParts := strings.Split(dir, "/")
	return dirParts[len(dirParts)-1]
}
