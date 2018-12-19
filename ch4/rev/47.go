package main

import (
	"fmt"
	"unicode/utf8"
)

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

// Reverse all the runes, and then the entire slice. The runes' bytes end up in
// the right order. DecodeRune unpacks the first UTF-8 encoding
// and returns the rune and its width in bytes.
func revUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	rev(b)
	return b
}

func main() {
	s := []byte("Päivää")
	fmt.Println(string(revUTF8(s)))
}
