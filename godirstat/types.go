package godirstat

import (
	"fmt"
	"sort"
)

type FileInfo struct {
	Name string
	Size int
}

type DirInfo struct {
	Name       string
	Size       int
	Files      []FileInfo
	TotalFiles int
	Dirs       []DirInfo
	TotalDirs  int
}

func (dir *DirInfo) SortDirsAndFiles() {
	sort.Slice(dir.Dirs, func(i, j int) bool {
		return dir.Dirs[i].Size > dir.Dirs[j].Size
	})

	sort.Slice(dir.Files, func(i, j int) bool {
		return dir.Files[i].Size > dir.Files[j].Size
	})

	for _, dir := range dir.Dirs {
		dir.SortDirsAndFiles()
	}
}

func (dir *DirInfo) PrintContents(level int, max_depth int) {
	prefix_dir := string(make([]rune, (3*level))) + " - "
	prefix_files := string(make([]rune, (3*(level+1)))) + " - "

	fmt.Printf(prefix_dir+dir.Name+" = %d bytes\n", dir.Size)

	// Return small summary when max depth is reached
	// Max depth of -1 or smaller means infinite depth
	if level > max_depth && max_depth > -1 {
		prefix_summary := string(make([]rune, (3*(level+1)))) + " + "
		fmt.Printf(prefix_summary+"%d more dirs and %d more files\n", dir.TotalDirs, dir.TotalFiles)
		return
	}

	for _, dir := range dir.Dirs {
		dir.PrintContents(level+1, max_depth)
	}

	// Create prefix with 3 spaces per level
	for _, file := range dir.Files {
		fmt.Printf(prefix_files+file.Name+" = %d bytes\n", file.Size)
	}
}
