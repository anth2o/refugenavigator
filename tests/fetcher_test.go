package tests

import (
	"os"
	"reflect"
	"testing"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

type FeatureCollectionQuerier struct {
	t *testing.T
}

func (q FeatureCollectionQuerier) QueryUrl(url string) []byte {
	expectedUrl := scrapper.GetBaseUrl() + "/bbox?" + getBoundingBoxStringTest()
	if url != expectedUrl {
		q.t.Errorf("QueryUrl() = %s, want %s", url, expectedUrl)
	}
	data, err := os.ReadFile("../data/bbox.json")
	checkError(err)
	return data
}

func TestGetFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	querier := FeatureCollectionQuerier{t: t}
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

type FeatureQuerier struct {
	t *testing.T
}

func (q FeatureQuerier) QueryUrl(url string) []byte {
	expectedUrl := scrapper.GetBaseUrl() + "/point?id=1198&detail=complet&format=geojson"
	if url != expectedUrl {
		q.t.Errorf("QueryUrl() = %s, want %s", url, expectedUrl)
	}
	data, err := os.ReadFile("../data/point.json")
	checkError(err)
	return data
}

func TestGetFeature(t *testing.T) {
	querier := FeatureQuerier{t: t}
	feature := scrapper.GetFeature(1198, querier)
	expectedFeature := getFeatureTest()
	if !reflect.DeepEqual(feature, expectedFeature) {
		t.Errorf("GetFeature() = %v, want %v", feature, expectedFeature)
	}
}
