package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) map[string]int {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts
}

func main() {

	counts := make(map[string]int)

	fmt.Printf("Variable counts detailed: %+v\n", counts)
	fmt.Printf("Variable counts value: %#v\n", counts)
	fmt.Printf("Variable counts type: %T\n", counts)
	fmt.Printf("Variable counts len: %d\n", len(counts))

	files := os.Args[1:]

	if len(files) == 0 {
		counts = countLines(os.Stdin, counts)
	} else {

		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts = countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
