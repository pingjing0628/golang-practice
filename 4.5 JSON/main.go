package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pingjing0628/golang-practice/4.5 JSON/github"
)

// 輸出符合搜尋條件的 Github 記錄表格
func main()  {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}