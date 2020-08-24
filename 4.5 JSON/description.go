package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 將 GO 資料結構 轉換成 JSON 被稱為 marshaling, 透過 json.Marshal 完成

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}
// omitempty 若欄位為空就不輸出

func main()  {
	// Marshaling
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Lake", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	// MarshalIndent 產生整齊的縮排
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	// Unmarshaling
	var titles []struct{Title string}

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println(titles)
}
