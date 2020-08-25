package xkcd

import (
	"fmt"
	"strconv"
)

type ComicIndex struct {
	Comic []*Comic
}

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

// Call API
func getComicURL(comicId int) string {
	return fmt.Sprintf("https://xkcd.com/%s/info.0.json", strconv.Itoa(comicId))
}

func NewComicIndex() ComicIndex {
	return ComicIndex{[]*Comic{}}
}