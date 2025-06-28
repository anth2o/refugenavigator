package refugenavigator

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

var baseUrl string = "https://www.refuges.info/api/bbox?"

func getFeatureCollection(bbox BoundingBox) *FeatureCollection {
	url := baseUrl + bbox.String()
	// Make a GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return nil
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}

	var featureCollection FeatureCollection
	json.Unmarshal(body, &featureCollection)
	return &featureCollection
}

func exportFeatureCollection(featureCollection *FeatureCollection) {
	f, err := os.Create("data/example.gpx")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(xml.Header); err != nil {
		fmt.Printf("err == %v", err)
	}
	if err := featureCollection.ToGpx().WriteIndent(f, "", "  "); err != nil {
		fmt.Printf("err == %v", err)
	}
}

func main() {
	bbox := BoundingBox{
		northEast: Point{5.52315, 44.9159},
		southWest: Point{5.49826, 44.8983},
	}
	featureCollection := getFeatureCollection(bbox)
	if featureCollection == nil {
		return
	}
	exportFeatureCollection(featureCollection)
}
