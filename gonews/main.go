package main

import (
	"fmt"

	"github.com/RichyRaj/go-hackernews"
)

func main() {
	fmt.Println("vim-go")

	// Get the new stories
	fmt.Println(hn.GetLatest())

}
