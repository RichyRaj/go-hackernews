package hn

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const maxCount = 10

var hnUrlBase = "https://hacker-news.firebaseio.com/v0/%s.json"

var hnItemBase = "https://hacker-news.firebaseio.com/v0/item/%d.json"

type Story struct {
	Title string
	Score int
	By    string
	Url   string
}

func makeError(message string) error {
	return errors.New(message)
}

func GetLatest() ([]*Story, error) {
	var items []int
	stories := make([]*Story, 0)
	url := fmt.Sprintf(hnUrlBase, "newstories")
	resp, err := http.Get(url)
	if err != nil {
		return nil, makeError("Error fetching the news")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, makeError("Error fetching the news")
	}
	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return nil, makeError("JSON parse error")
	}
	for i := range items {
		story := new(Story)
		url = fmt.Sprintf(hnItemBase, items[i])
		resp, err = http.Get(url)
		if err == nil {
			err = json.NewDecoder(resp.Body).Decode(story)
			if err == nil {
				stories = append(stories, story)
			}
		}
		if len(stories) > maxCount {
			break
		}
	}
	return stories, nil

}
