// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bormansquirrel/gopl.io/ch4/github"
)

const (
	UnixMonth = 2678400
	UnixYear  = 31536000
)

func main() {
	unixTimeNow := time.Now().Unix()
	ageCategories := make(map[string][]*github.Issue)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("%d issues:\n", len(result.Items))
	for _, item := range result.Items {
		unixTime := item.CreatedAt.Unix()
		if unixTimeNow-unixTime < UnixMonth {
			ageCategories["<month"] = append(ageCategories["<month"], item)
			continue
		}
		if unixTimeNow-unixTime < UnixYear {
			ageCategories["<year"] = append(ageCategories["<year"], item)
			continue
		}
		if unixTimeNow-unixTime >= UnixYear {
			ageCategories[">=year"] = append(ageCategories["<year"], item)
			continue
		}
		fmt.Println("yeah")
	}
	fmt.Println("#### Less than a month")
	for _, element := range ageCategories["<month"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			element.Number, element.User.Login, element.Title)
	}
	fmt.Println("#### Less than a year")
	for _, element := range ageCategories["<year"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			element.Number, element.User.Login, element.Title)
	}
	fmt.Println("#### More than a year")
	for _, element := range ageCategories[">=year"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			element.Number, element.User.Login, element.Title)
	}
}
