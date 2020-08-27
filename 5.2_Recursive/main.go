package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// Ex5.4 新增取出圖像等等
var linkAttrs = map[string][]string {
	"a":      []string{"href"},
	"link":   []string{"href"},
	"img":    []string{"src"},
	"script": []string{"src"},
	"iframe": []string{"src"},
	"form":   []string{"action"},
	"html":   []string{"manifest"},
	"video":  []string{"src", "poster"},
}

func main()  {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	//for _, link := range visit(nil, doc) {
	//	fmt.Println(link)
	//}

	// Ex5.2
	//for key, value := range count(make(map[string]int), doc) {
	//	fmt.Printf("%s: %d\n", key, value)
	//}

	// Ex5.3
	//notRecord(os.Stdout, doc)

	// Ex5.4
	for _, link := range Others(nil, doc) {
		fmt.Println(link)
	}
}

//func visit(links []string, n *html.Node) []string {
//	if n == nil {
//		return links
//	}
//
//	if n.Type == html.ElementNode && n.Data == "a" {
//		for _, a := range n.Attr {
//			if a.Key == "href" {
//				links = append(links, a.Val)
//			}
//		}
//	}
//
//	// Ex5.1 修改程式以使用 visit 遞迴呼叫替換迴圈遍歷 n.FirstChild
//	links = visit(links, n.FirstChild)
//	links = visit(links, n.NextSibling)
//
//	return links
//}

// Ex5.2 計算HTML文件中 p div span 的數量
//func count(counts map[string]int, n *html.Node) map[string]int {
//	if n == nil {
//		return counts
//	}
//
//	if n.Type == html.ElementNode {
//		counts[n.Data]++
//	}
//
//	count(counts, n.FirstChild)
//	count(counts, n.NextSibling)
//
//	return counts
//}

// Ex5.3 不要紀錄含有 script 和 style 的元素
//func notRecord(w io.Writer, n *html.Node) {
//	if n == nil {
//		return
//	}
//
//	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
//		notRecord(w, n.NextSibling)
//		return
//	} else if n.Type == html.TextNode {
//		w.Write([]byte(n.Data))
//	}
//
//	notRecord(w, n.FirstChild)
//	notRecord(w, n.NextSibling)
//}

// Ex5.4
func Others(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		for k, v := range linkAttrs {
			if n.Data == k {
				for _, attr := range v {
					for _, a := range n.Attr {
						if a.Key == attr {
							links = append(links, a.Val)
						}
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Others(links, c)
	}

	return links
}
