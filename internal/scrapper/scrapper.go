package scrapper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
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

func ExportFeatureCollection(featureCollection *FeatureCollection, outputFile string) {
	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(xml.Header); err != nil {
		fmt.Printf("err == %v", err)
		return
	}
	if err := featureCollection.ToGpx().WriteIndent(f, "", "  "); err != nil {
		fmt.Printf("err == %v", err)
		return
	}
	fmt.Printf("GPX was successfully exported to %s\n", outputFile)
}
