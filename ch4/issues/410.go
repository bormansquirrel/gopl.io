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

	"gopl.io/ch4/github"
)

//!+

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
	fmt.Printf("%d issues:\n", result.TotalCount)
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

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
