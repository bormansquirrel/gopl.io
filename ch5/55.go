// ex5.5 counts the number of words and images at a url.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	var w, i int
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return 0, 0, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	words, images = countWordsAndImages(doc, &w, &i)
	return
}

func countWordsAndImages(n *html.Node, words, images *int) (int, int) {
	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		*images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		*words, *images = countWordsAndImages(c, words, images)
	}
	return *words, *images
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: 55 url")
	}
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}
