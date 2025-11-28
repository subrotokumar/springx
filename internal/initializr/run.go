package initializr

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	SpringInitializrMetadataURL = "https://start.spring.io/metadata/client"
)

func Run() {
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

	// save reponsonse to file
	outFile, err := os.Create("res.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving response to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Metadata saved to res.json")
}
