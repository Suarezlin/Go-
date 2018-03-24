package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			if f, err := os.Open(file); err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			} else {
				countLines(f, counts)
				f.Close()
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
