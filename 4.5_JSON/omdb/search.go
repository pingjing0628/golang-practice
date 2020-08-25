package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Search from API
func search(terms []string) (*Movie, error) {
	var result Movie 
	response, err := http.Get(searchURL(terms))

	if err != nil {
		return nil, err
	}
	
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		
		return nil, fmt.Errorf("get poster failed: %s", response.Status)
	}
	
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		response.Body.Close()

		return nil, err 
	}
	
	response.Body.Close()

	return &result, nil
}

// Get poster from search
func GetPoster(w io.Writer, terms []string) error {
	movie, err := search(terms)

	if err != nil {
		return err
	}

	if movie.Response != "True" {
		return fmt.Errorf("movie not found: %v", terms)
	}

	response , err := http.Get(movie.Poster)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return fmt.Errorf("get poster failed: %s", response.Status)
	}

	_, err = io.Copy(w, response.Body)

	response.Body.Close()

	return err
}
