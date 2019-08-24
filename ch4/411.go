package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	//"os"
	//"strings"
	//"time"
)

const (
	URL                 = "https://api.github.com/repos"
	Owner               = "bormansquirrel"
	Repo                = "devops-test"
	PersonalAccessToken = "xxxxxxxxxxxxxx"
)

type Issue struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

// CreateIssue in GitHub issue tracker.
func CreateIssue(issue *Issue, owner string, repo string) error {
	jsonStr, err := json.Marshal(issue)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", jsonStr)
	fmt.Println(URL + "/" + owner + "/" + repo + "/issues")
	req, err := http.NewRequest("POST", URL+"/"+owner+"/"+repo+"/issues", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+PersonalAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}

func main() {
	issue := &Issue{
		Title:  "myissue1",
		Body:   "mybody1",
		Labels: []string{"mylabel1", "mylabel2"},
	}

	err := CreateIssue(issue, Owner, Repo)
	if err != nil {
		log.Fatalln(err)
	}
}
