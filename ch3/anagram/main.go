// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 74 Exercise 3.12.

// Anagram reports weather two strings are anagrams of each other, that is, they contain
// the same characters but in different order
//
// Example:
// 	$ go build gopl.io/ch3/anagram
//	$ ./anagram chocolate etalchoco
//	  true
//
package main

import (
	"fmt"
	"os"
)

func main() {
	if anagram(os.Args[1], os.Args[2]) {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
}

//!+
func anagram(s1, s2 string) bool {
	m := map[byte]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}

	lenS1 := len(s1)
	lenS2 := len(s2)

	if lenS1 != lenS2 {
		return false
	}
	for i := 0; i <= lenS1-1; i++ {
		m[s1[i]]++
	}
	for i := 0; i <= lenS2-1; i++ {
		if m[s2[i]] != 1 {
			return false
		}
	}

	return true
}

//!-
