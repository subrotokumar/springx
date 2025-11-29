package spring

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	SpringInitializrMetadataURL = "https://start.spring.io/metadata/client"
)

func Run() (SpringInitializrResponse, error) {
	res, err := http.Get(SpringInitializrMetadataURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching metadata: %v\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Unexpected status code: %d\n", res.StatusCode)
		os.Exit(1)
	}

	var response SpringInitializrResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return SpringInitializrResponse{}, fmt.Errorf("failed to decode metadata: %w", err)
	}

	return response, nil
}
