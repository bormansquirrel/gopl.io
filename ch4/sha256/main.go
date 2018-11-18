// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import "crypto/sha256"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%b\n%b\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	count := 0
	for i, _ := range c1 {
		for c1[i] > 0 { // until all bits are zero
			if c1[i]&1 != c2[i]&1 { // check lower bit
				count++
			}
			c1[i] >>= 1 // shift bits, removing lower bit
			c2[i] >>= 1
		}
	}

	fmt.Printf("Number of different bits: %d\n", count)
}

//!-
