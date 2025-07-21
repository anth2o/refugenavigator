// Source code to fetch data from refuges.info API
// Here is the documentation about it https://www.refuges.info/api/doc/#/api

package scrapper

import (
	"encoding/json"
	"fmt"
)

var baseUrl string

func SetBaseUrl(url string) {
	baseUrl = url
}

func GetBaseUrl() string {
	if baseUrl == "" {
		baseUrl = "https://www.refuges.info/"
	}
	return baseUrl
}

func parseFeatureCollection(body []byte) *FeatureCollection {
	var featureCollection FeatureCollection
	json.Unmarshal(body, &featureCollection)
	return &featureCollection
}

func GetFeatureCollection(bbox BoundingBox) *FeatureCollection {
	body := QueryUrl(GetBaseUrl() + "api/bbox?nb_points=all&" + bbox.String())
	return parseFeatureCollection(body)
}

func GetFeature(featureId int) *Feature {
	body := QueryUrl(GetBaseUrl() + "api/point?id=" + fmt.Sprint(featureId) + "&detail=complet&format=geojson&format_texte=markdown")
	featureCollection := parseFeatureCollection(body)
	feature := featureCollection.Features[0]
	if feature.Properties.Type.Valeur == "point d'eau" {
		html := QueryUrl(GetBaseUrl() + "point/" + fmt.Sprint(featureId))
		comments, err := ScrapeComments(string(html))
		if err != nil {
			fmt.Println("Error scraping comments for feature", featureId, err)
		} else {
			feature.CommentData.Comments = comments
			feature.CommentData.Prompt = prompt
			summary, restored := feature.GetCommentsSummaryFromDb()
			if restored {
				fmt.Printf("Comments summary restored for feature %v\n", featureId)
				feature.CommentData.Summary = summary
			} else {
				feature.CommentData.Summary = feature.GetCommentsSummaryFromLlm()
				fmt.Printf("Comments summary not restored for feature %v. Summary from LLM is %v\n", featureId, feature.CommentData.Summary)
				feature.StoreCommentsSummary()
			}
		}
	}
	return &feature
}

func EnrichFeatureCollection(featureCollection *FeatureCollection) {
	syncChannel := make(chan int, len(featureCollection.Features))
	for i := range featureCollection.Features {
		go func(i int) {
			featureCollection.Features[i] = *GetFeature(featureCollection.Features[i].Id)
			syncChannel <- 1
		}(i)
	}
	for i := 0; i < len(featureCollection.Features); i++ {
		<-syncChannel
	}
}
