package main

import "flag"

const (
	maxComicId = 1833
)

var fetchFlag = flag.Bool("fetch", false, "fetch all comics")

func fetch()  {
	comicIndex := xkcd.NewComicIndex()
}

func main()  {
	flag.Parse()

	if *fetchFlag {
		fetch()
	}
}
