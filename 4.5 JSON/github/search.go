package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Search Github record
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// url.QueryEscape 確保特殊字元會以字面處理
	q := url.QueryEscape(strings.Join(terms, " "))
	
	resp, err := http.Get(IssueURL + "?q=" + q)
	
	if err != nil {
		return nil, err
	}
	
	// 必須在所有執行路徑上關閉 resp.Body(以後會使用 defer)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()

		return nil, err
	}
	
	resp.Body.Close()

	return &result, nil
}
