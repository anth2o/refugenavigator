package scrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var baseUrl string = "https://www.refuges.info/api/bbox?"

type Querier interface {
	QueryUrl(url string) []byte
}

type DefaultQuerier struct{}

func (d DefaultQuerier) QueryUrl(url string) []byte {
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

func GetFeatureCollection(bbox BoundingBox, querier Querier) *FeatureCollection {
	url := baseUrl + bbox.String()
	if querier == nil {
		querier = DefaultQuerier{}
	}
	body := querier.QueryUrl(url)
	if body == nil {
		return nil
	}
	var featureCollection FeatureCollection
	json.Unmarshal(body, &featureCollection)
	return &featureCollection
}
