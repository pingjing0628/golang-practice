package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIssues(owner, repo string) ([]Issue, error) {
	var issues []Issue
	req, err := http.NewRequest("GET", getIssueURL(owner, repo), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed: %s", response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(&issues); err != nil {
		return nil, err
	}

	return issues, nil
}
