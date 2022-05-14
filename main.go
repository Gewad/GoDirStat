package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

type FileInfo struct {
	name string
	size int
}

type DirInfo struct {
	name  string
	size  int
	files []FileInfo
	dirs  []DirInfo
}

type RootInfo struct {
	rootDir string
	size    int
	files   []FileInfo
	dirs    []DirInfo
}

func GetDirInfo(dirString string, size int) DirInfo {
	items, err := ioutil.ReadDir(dirString)
	if err != nil {
		log.Fatal(err)
	}

	dir := DirInfo{
		name:  GetNameFromDir(dirString),
		size:  size,
		files: GetDirFiles(items),
		dirs:  GetDirDirs(items, dirString)}

	dir.size += CalculateTotalSize(dir)
	return dir
}

func GetDirDirs(items []fs.FileInfo, dirStringPrev string) []DirInfo {
	var dirs []DirInfo

	for _, item := range items {
		if item.IsDir() {
			dirString := dirStringPrev + "/" + item.Name()
			dirs = append(dirs, GetDirInfo(dirString, int(item.Size())))
		}
	}

	return dirs
}

func GetDirFiles(items []fs.FileInfo) []FileInfo {
	var files []FileInfo

	for _, item := range items {
		if !item.IsDir() {
			files = append(files, FileInfo{item.Name(), int(item.Size())})
		}
	}

	return files
}

func CalculateTotalSize(dir DirInfo) int {
	totalSize := 0

	return totalSize
}

func GetNameFromDir(dir string) string {
	dirParts := strings.Split(dir, "/")
	return dirParts[len(dirParts)-1]
}

func GetRootInfo(dir string) RootInfo {
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var dirs []DirInfo
	var files []FileInfo

	for _, item := range items {
		if item.IsDir() {
			dirString := dir + "/" + item.Name()
			dirs = append(dirs, GetDirInfo(dirString, int(item.Size())))
		} else {
			files = append(files, FileInfo{item.Name(), int(item.Size())})
		}
	}

	totalSize := 0
	for _, dir := range dirs {
		totalSize += int(dir.size)
	}

	for _, file := range files {
		totalSize += int(file.size)
	}

	return RootInfo{rootDir: dir, size: totalSize, files: files, dirs: dirs}
}

func main() {
	root := GetRootInfo("/home/g/")
	fmt.Println(root.size)
}
