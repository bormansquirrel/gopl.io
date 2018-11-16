// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 74, Exercise 3.11

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
//	$ ./comma 1 12 123. .123 1.234 123456.7890
// 	1
// 	12
// 	123
// 	123
// 	1.,234
// 	1,234,56.7,890
//	$ ./comma 1 12 -123 -12.34 123456.7890
// 	1
// 	12
// 	123
// 	123
// 	-1,2.34
// 	1,234,56.7,890
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
	var reversedBuffer bytes.Buffer

	j := 1
	for i := len(s) - 1; i >= 0; i--{
		buffer.WriteByte(s[i])
		if s[i] == '.' || s[i] == '-' || s[i] == '+'{
			continue;
		}
		if j%3 == 0{
			buffer.WriteString(",")
		}
		j++
	}

	s = buffer.String()
	for i := len(s) - 1; i >= 0; i--{
		reversedBuffer.WriteByte(s[i])
	}

	return reversedBuffer.String()
}
//!-

