package scrapper

import (
	"fmt"
	"io"
	"net/http"
)

func QueryUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}
	return body
}
