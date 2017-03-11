package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var hnUrlBase = "https://hacker-news.firebaseio.com/v0/%s.json"

func GetLatest() string {
	url := fmt.Sprintf(hnUrlBase, "newstories")
	var items []int
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return "error"
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return resp.Status + "Error"
	}
	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		fmt.Println(err)
		return "JSON Parse Error"
	}
	fmt.Println(items)
	return "Done"

}
