package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pingjing0628/golang-practice/4.5_JSON/xkcd"
	"io/ioutil"
	"os"
)

const (
	maxComicId = 1833
)

var fetchFlag = flag.Bool("fetch", false, "fetch all comics")

func fetch()  {
	comicIndex := xkcd.NewComicIndex()

	for comicId := 1; comicId <= maxComicId; comicId++ {
		comic, err := xkcd.GetComic(comicId)

		if err != nil {
			fmt.Fprintf(os.Stderr, "get %d: %v\n", comicId, err)
		} else {
			comicIndex.Comic = append(comicIndex.Comic, comic)
		}
	}

	result, err := json.MarshalIndent(comicIndex, "", "   ")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", result)
	}
}

func search(terms []string)  {
	index, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	comics := xkcd.NewComicIndex()

	err = json.Unmarshal(index, &comics)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	searchResult := xkcd.SearchComic(comics, terms)

	fmt.Printf("%d comics:\n", len(searchResult))

	for _, comic := range searchResult{
		printComic(comic)
	}
}

func printComic(comic *xkcd.Comic)  {
	fmt.Printf("\n-- Comic %d --\n", comic.Num)
	fmt.Printf("\nImage URL:\n%s\n", comic.Img)
	fmt.Printf("\nTranscript:\n%s\n", comic.Transcript)
}

func main()  {
	flag.Parse()

	if *fetchFlag {
		fetch()
	} else {
		if len(flag.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "must have at least 1 query")
			os.Exit(1)
		}

		search(flag.Args())
	}
}
