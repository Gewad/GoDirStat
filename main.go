package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Gewad/GoDirStat/godirstat"
)

func main() {
	dirPtr := flag.String("dir", ".", "Directory to scan")
	print_depth := flag.Int("print_depth", 0, "Max depth to print in terminal")
	flag.Parse()

	begin := time.Now()
	root := godirstat.GetDirInfo(*dirPtr, 0)
	root.SortDirsAndFiles()
	end := time.Now()

	fmt.Printf("Scan complete for folder: %s\n", *dirPtr)
	fmt.Printf("Time in milliseconds:     %d\n", end.Sub(begin).Milliseconds())
	root.PrintContents(0, *print_depth)
}
