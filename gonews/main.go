package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/RichyRaj/go-hackernews"
)

func main() {
	fmt.Println("######## 10 Latest Hacker News Stories ########")
	// Get the new stories
	latest, err := hn.GetLatest()
	if err != nil {
		log.Fatal(err)
	}
	for _, story := range latest {
		fmt.Println()
		fmt.Println(story.Title + " (" + strconv.Itoa(story.Score) + " point(s))")
		fmt.Println("By " + story.By)
		fmt.Println(story.Url)
	}
}
