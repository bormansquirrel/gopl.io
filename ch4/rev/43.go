// Rev reverses an array of ints.
package main

import (
	//"bufio"
	//"os"
	//"strconv"
	//"strings"
	"fmt"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array
}

//!+rev
// reverse reverses an array of ints by getting an array pointer.
func reverse(s *[6]int) {
	for i := 0; i < len(s)/2; i++ {
		end := len(s) - i - 1
		s[i], s[end] = s[end], s[i]
	}
}

//!-rev
