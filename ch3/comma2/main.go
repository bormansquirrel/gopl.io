// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 74, Exercise 3.10

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma2
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
func comma(s string) string {
	var buffer bytes.Buffer

	for i := 0; i <= len(s) - 1; i++ {
		buffer.WriteByte(s[i])
		if i != len(s) - 1 && len(s) > 3 && i%3 == 0 {
			buffer.WriteString(",")
		}
	}

	return buffer.String()
}
//!-

