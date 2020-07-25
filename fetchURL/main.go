package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main()  {
	for _, url := range os.Args[1:] {
		// Practice 1.8 add prefix
		hasPrefix := strings.HasPrefix(url, "https://")

		if ! hasPrefix {
			url = "https://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// Practice 1.9 Print resp.Status HTTP status code
		fmt.Printf("http-get: %v\n", resp.Status)

		//b, err := ioutil.ReadAll(resp.Body)
		// Practice 1.7
        _, err = io.Copy(os.Stdout, resp.Body) // 從第二個參數讀取並寫入至第一個參數

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		//fmt.Printf("%s", b)
	}
}
