// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)     // counts of Unicode characters
	runeType := make(map[string]int) // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int  // count of lengths of UTF-8 encodings
	invalid := 0                     // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsControl(r) {
			runeType["control"]++
		}
		if unicode.IsDigit(r) {
			runeType["digit"]++
		}
		if unicode.IsGraphic(r) {
			runeType["graphic"]++
		}
		if unicode.IsLetter(r) {
			runeType["letter"]++
		}
		if unicode.IsLower(r) {
			runeType["lower"]++
		}
		if unicode.IsMark(r) {
			runeType["mark"]++
		}
		if unicode.IsNumber(r) {
			runeType["number"]++
		}
		if unicode.IsPrint(r) {
			runeType["print"]++
		}
		if unicode.IsPunct(r) {
			runeType["punct"]++
		}
		if unicode.IsSpace(r) {
			runeType["space"]++
		}
		if unicode.IsSymbol(r) {
			runeType["symbol"]++
		}
		if unicode.IsTitle(r) {
			runeType["title"]++
		}
		if unicode.IsUpper(r) {
			runeType["upper"]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Printf("type\tcount\n")
	for c, n := range runeType {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

//!-
