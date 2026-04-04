// Source code to fetch data from refuges.info API
// Here is the documentation about it https://www.refuges.info/api/doc/#/api

package scrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var baseUrl string = "https://www.refuges.info/api"

func GetBaseUrl() string {
	return baseUrl
}

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

func query(url string, querier Querier) []byte {
	if querier == nil {
		querier = DefaultQuerier{}
	}
	return querier.QueryUrl(url)
}

func parseFeatureCollection(body []byte) *FeatureCollection {
	var featureCollection FeatureCollection
	json.Unmarshal(body, &featureCollection)
	return &featureCollection
}

func GetFeatureCollection(bbox BoundingBox, querier Querier) *FeatureCollection {
	body := query(baseUrl+"/bbox?nb_points=all&"+bbox.String(), querier)
	return parseFeatureCollection(body)
}

func GetFeature(featureId int, querier Querier) *Feature {
	body := query(baseUrl+"/point?id="+fmt.Sprint(featureId)+"&detail=complet&format=geojson&format_texte=markdown", querier)
	featureCollection := parseFeatureCollection(body)
	return &featureCollection.Features[0]
}

func EnrichFeatureCollection(featureCollection *FeatureCollection, querier Querier) {
	for i := range featureCollection.Features {
		featureCollection.Features[i] = *GetFeature(featureCollection.Features[i].Id, querier)
	}
}
