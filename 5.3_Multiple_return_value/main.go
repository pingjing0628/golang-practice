package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main()  {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "findLinks2: %v\n", err)
			continue
		}

		for _, link := range links{
			fmt.Println(link)
		}
	}
}

// 對 url 執行 HTTP 的 GET 請求，解析 HTML 回應，然後取出與回傳連結
func findLinks(url string) ([]string, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		// 確保再發生錯誤時也能釋放網路資源
		// Go 的垃圾回收會將未使用的記憶體循環
		response.Body.Close()

		return nil, fmt.Errorf("getting %s: %s", url, response.Status)
	}

	doc, err := html.Parse(response.Body)

	response.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// Ex5.1 修改程式以使用 visit 遞迴呼叫替換迴圈遍歷 n.FirstChild
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}

func findLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)

	return findLinks(url)
}