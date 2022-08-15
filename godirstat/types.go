package godirstat

import (
	"fmt"
	"sort"
)

type FileInfo struct {
	name string
	size int
}

type DirInfo struct {
	name       string
	size       int
	files      []FileInfo
	totalFiles int
	dirs       []DirInfo
	totalDirs  int
}

func (dir *DirInfo) SetName(name string) {
	dir.name = name
}

func (dir *DirInfo) SortDirsAndFiles() {
	sort.Slice(dir.dirs, func(i, j int) bool {
		return dir.dirs[i].size > dir.dirs[j].size
	})

	sort.Slice(dir.files, func(i, j int) bool {
		return dir.files[i].size > dir.files[j].size
	})

	for _, dir := range dir.dirs {
		dir.SortDirsAndFiles()
	}
}

func (dir *DirInfo) PrintContents(level int, max_depth int) {
	prefix_dir := string(make([]rune, (3*level))) + " - "
	prefix_files := string(make([]rune, (3*(level+1)))) + " - "

	fmt.Printf(prefix_dir+dir.name+" = %d bytes\n", dir.size)

	// Return small summary when max depth is reached
	// Max depth of -1 or smaller means infinite depth
	if level > max_depth && max_depth > -1 {
		prefix_summary := string(make([]rune, (3*(level+1)))) + " + "
		fmt.Printf(prefix_summary+"%d more dirs and %d more files\n", dir.totalDirs, dir.totalFiles)
		return
	}

	for _, dir := range dir.dirs {
		dir.PrintContents(level+1, max_depth)
	}

	// Create prefix with 3 spaces per level
	for _, file := range dir.files {
		fmt.Printf(prefix_files+file.name+" = %d bytes\n", file.size)
	}
}
