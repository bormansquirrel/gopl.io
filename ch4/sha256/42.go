package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var hashType string

	flag.StringVar(&hashType, "hash", "SHA256", "the type of digest of a string")

	flag.Parse()

	// Wrapping the unbuffered `os.Stdin` with a buffered
	// scanner gives us a convenient `Scan` method that
	// advances the scanner to the next token; which is
	// the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` returns the current token, here the next line,
		// from the input.
		if hashType == "SHA384" {
			h := sha512.New384()
			h.Write([]byte(scanner.Text()))
			fmt.Printf("%x\n", h.Sum(nil))
		} else if hashType == "SHA512" {
			h := sha512.New()
			h.Write([]byte(scanner.Text()))
			fmt.Printf("%x\n", h.Sum(nil))
		} else {
			h := sha256.New()
			h.Write([]byte(scanner.Text()))
			fmt.Printf("%x\n", h.Sum(nil))
		}
	}

	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
