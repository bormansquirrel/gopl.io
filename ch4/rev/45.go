// ex4.5 in-place function to eliminated adjacent duplicates
// in a []string slice
package main

import (
	"fmt"
)

func eliminate_adjacents(strings []string) []string {
	murders := 0
	for i := 0; i+murders < len(strings)-1; i++ {
		for strings[i] == strings[i+1] {
			if i+murders == len(strings)-1 {
				break
			}
			copy(strings[i:], strings[i+1:])
			murders++
		}
	}
	return strings[:len(strings)-murders]
}

func main() {
	r := []string{"hello", "world", ":)"}
	fmt.Println(eliminate_adjacents(r))
	s := []string{"world", "world", "world"}
	fmt.Println(eliminate_adjacents(s))
	t := []string{"hello", "world", "world", "and", "universe"}
	fmt.Println(eliminate_adjacents(t))
	u := []string{"hello", "hello", "world", "world", "and", "and", "universe", "universe"}
	fmt.Println(eliminate_adjacents(u))
}
