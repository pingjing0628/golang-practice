package omdb

import (
	"fmt"
	"net/url"
	"strings"
)

type Movie struct {
	Poster   string
	Response string
}

// Call API
func searchURL(terms []string) string {
	return fmt.Sprintf("https://www.omdbapi.com/?t=%s", url.QueryEscape(strings.Join(terms, " ")))
}

