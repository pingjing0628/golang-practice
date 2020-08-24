package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pingjing0628/golang-practice/4.5_JSON/github"
)

// 輸出符合搜尋條件的 Github 記錄表格
func main()  {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d Issues:\n", result.TotalCount)

	now := time.Now()

	beforeMonth := now.AddDate(0, -1, 0)
	beforeYear := now.AddDate(-1, 0, 0)

	fmt.Println("\n-- Created at less than a month --")

	for _, item := range result.Items {
		if item.CreatedAt.After(beforeMonth) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("\n-- Created at less than a year --")

	for _, item := range result.Items {
		if (item.CreatedAt.Before(beforeMonth) || item.CreatedAt.Equal(beforeMonth)) && item.CreatedAt.After(beforeYear) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("\n-- Created at more than a year --")

	for _, item := range result.Items {
		if item.CreatedAt.Before(beforeYear) || item.CreatedAt.Equal(beforeYear) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

}