// Dup4 prints the count, text, and source files of lines that appear
// more than once in the input. Reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup4() {
	counts := make(map[string]int)
	fileNames := make(map[string]map[string]bool)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var fileList []string
			for f := range fileNames[line] {
				fileList = append(fileList, f)
			}
			fmt.Printf("%d\t%s\t%v\n", n, line, fileList)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if fileNames[line] == nil {
			fileNames[line] = make(map[string]bool)
		}
		fileNames[line][filename] = true
	}
}
