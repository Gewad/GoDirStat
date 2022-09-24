package godirstat

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type Config struct {
	FnReadDir func(string) ([]fs.DirEntry, error)
}

var defaultConfig = Config{
	FnReadDir: os.ReadDir,
}

func GetDirInfo(dirString string, size int) DirInfo {
	return GetDirInfoWithConfig(dirString, size, defaultConfig)
}

func GetDirInfoWithConfig(dirString string, size int, config Config) DirInfo {
	entries, err := config.FnReadDir(dirString)

	if err != nil {
		message := dirString + " - " + err.Error()
		fmt.Println(message)
		return DirInfo{}
	}

	items := convertToFileInfo(entries)

	dir := DirInfo{
		Name:  getNameFromDir(dirString),
		Size:  size,
		Files: getDirFiles(items),
		Dirs:  getDirDirs(items, dirString, config)}

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

func convertToFileInfo(items []fs.DirEntry) []fs.FileInfo {
	converted := make([]fs.FileInfo, len(items))
	for i, v := range items {
		converted[i], _ = v.Info()
	}
	return converted
}

func getDirDirs(items []fs.FileInfo, dirStringPrev string, config Config) []DirInfo {
	var dirs []DirInfo

	for _, item := range items {
		if item.IsDir() {
			dirString := dirStringPrev + "/" + item.Name()
			dirs = append(dirs, GetDirInfoWithConfig(dirString, int(item.Size()), config))
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
