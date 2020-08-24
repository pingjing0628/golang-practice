package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetComic(comicId int) (*Comic, error) {
	var result Comic
	resp, err := http.Get(getComicURL(comicId))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()

		return nil, fmt.Errorf("get comic failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()

		return nil, err
	}

	resp.Body.Close()

	return &result, nil
}
