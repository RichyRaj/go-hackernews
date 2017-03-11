package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var hnUrlBase = "https://hacker-news.firebaseio.com/v0/%s.json"

var hnItemBase = "https://hacker-news.firebaseio.com/v0/item/%d.json"

type Story struct {
	Title string
	Score int
	By    string
	Url   string
}

func GetLatest() string {
	var items []int
	url := fmt.Sprintf(hnUrlBase, "newstories")
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return "error"
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return resp.Status + "Error"
	}
	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		fmt.Println(err)
		return "JSON Parse Error"
	}
	fmt.Println(items)
	items = items[:11]
	for i := range items {
		story := new(Story)
		url = fmt.Sprintf(hnItemBase, items[i])
		resp, err = http.Get(url)
		// TODO : Error Handling
		err = json.NewDecoder(resp.Body).Decode(story)
		fmt.Println(story)
	}
	return "Done"

}
