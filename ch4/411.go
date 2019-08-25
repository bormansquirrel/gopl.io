package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strconv"
	//"net/url"
	//"os"
	//"io"
	//"time"
)

const (
	BaseURL             = "https://api.github.com"
	Owner               = "bormansquirrel"
	Repo                = "devops-test"
	PersonalAccessToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxx"
)

type Issue struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Number int    `json:"number,omitempty"`
}

// GetIssue in GitHub issue tracker.
func GetIssue(owner string, repo string, issueNumber int) (*Issue, error) {
	u := fmt.Sprintf("%s/repos/%s/%s/issues/%d", BaseURL, owner, repo, issueNumber)

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Authorization", "token "+PersonalAccessToken)

	i := new(Issue)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	err1 := json.Unmarshal(body, i)
	if err1 != nil {
		return nil, err
	}

	return i, nil
}

// CreateIssue in GitHub issue tracker.
func CreateIssue(issue *Issue, owner string, repo string) (int, error) {
	var respIssue Issue
	u := fmt.Sprintf("%s/repos/%s/%s/issues", BaseURL, owner, repo)

	jsonStr, err := json.Marshal(issue)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+PersonalAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	err1 := json.Unmarshal(body, &respIssue)
	if err1 != nil {
		return -1, err
	}

	//TODO: not working with NewDecoder ...
	//dec := json.NewDecoder(req.Body)
	//for {
	//	if err := dec.Decode(&respIssue); err == io.EOF {
	//		break
	//	} else if err != nil {
	//		return -1, err
	//	}
	//}

	return respIssue.Number, nil
}

func main() {
	issue := &Issue{
		Title: "myissue6",
		Body:  "mybody6",
	}

	//create an issue
	issueNumber, err := CreateIssue(issue, Owner, Repo)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("########################")
	fmt.Printf("Issue Number: %d\n", issueNumber)
	fmt.Println("########################")

	//get the issue previously created
	issue, err1 := GetIssue(Owner, Repo, issueNumber)
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Println("########################")
	fmt.Printf("Title: %s, Body: %s\n", issue.Title, issue.Body)
	fmt.Println("########################")
}
