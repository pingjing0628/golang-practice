package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main()  {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		fmt.Printf("words: %d\n", words)
		fmt.Printf("images: %d\n", images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	response, err := http.Get(url)

	if err != nil {
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("getting %s: %s", url, err)
		return
	}

	doc, err := html.Parse(response.Body)

	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}

	words, images = countWordsAndImages(doc)

	// 僅回傳(代表回傳全部)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return countWordsAndImages(n.NextSibling)
	} else if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	cwords, cimages := countWordsAndImages(n.FirstChild)
	swords, simages := countWordsAndImages(n.NextSibling)

	words += cwords + swords
	images += cimages + simages

	return
}