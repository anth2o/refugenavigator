package tests

import (
	"os"
	"reflect"
	"testing"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

type TestQuerier struct {
	t *testing.T
}

func (q TestQuerier) QueryUrl(url string) []byte {
	data, err := os.ReadFile("../data/example.json")
	checkError(err)
	return data
}

func TestGetFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	querier := TestQuerier{t: t}
	featureCollection := scrapper.GetFeatureCollection(bbox, querier)
	expectedFeatureCollection := getFeatureCollectionTest()
	if len(featureCollection.Features) != len(expectedFeatureCollection.Features) {
		t.Errorf("GetFeatureCollection() has length %v, want %v", len(featureCollection.Features), len(expectedFeatureCollection.Features))
	}
	for i := range featureCollection.Features {
		if !reflect.DeepEqual(featureCollection.Features[i], expectedFeatureCollection.Features[i]) {
			t.Errorf("GetFeatureCollection() = %v, want %v", featureCollection.Features[i], expectedFeatureCollection.Features[i])
		}
	}
}
