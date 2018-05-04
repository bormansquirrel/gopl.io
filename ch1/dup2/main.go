// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type match struct {
	n int
	file string
}

func main() {
	counts := make(map[string]*match)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.n > 1 {
			fmt.Printf("%s\t%d\t%s\n", n.file, n.n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*match) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
      counts[input.Text()] = &match{}
		}
		counts[input.Text()].n++
		if ! strings.Contains(counts[input.Text()].file, f.Name()) {
			counts[input.Text()].file += " " + f.Name()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
