package tests

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"net/url"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

type UnifiedQuerier struct {
	t *testing.T
}

func (q UnifiedQuerier) QueryUrl(url string) []byte {
	// Check if this is a bbox query
	if strings.Contains(url, "/bbox") {
		expectedUrl := scrapper.GetBaseUrl() + "/bbox?" + getBoundingBoxStringTest()
		if url != expectedUrl {
			q.t.Errorf("QueryUrl() = %s, want %s", url, expectedUrl)
		}
		data, err := os.ReadFile("../data/bbox.json")
		checkError(err)
		return data
	}

	// Check if this is a point query
	if strings.Contains(url, "/point") {
		id, err := extractId(url)
		if err != nil {
			q.t.Errorf("QueryUrl() = %s, want to be able to extract an id", url)
		}
		expectedUrl := scrapper.GetBaseUrl() + "/point?id=" + strconv.Itoa(id) + "&detail=complet&format=geojson&format_texte=markdown"
		if url != expectedUrl {
			q.t.Errorf("QueryUrl() = %s, want %s", url, expectedUrl)
		}
		data, err := os.ReadFile("../data/" + strconv.Itoa(id) + ".json")
		checkError(err)
		return data
	}

	q.t.Errorf("QueryUrl() = %s, want either bbox or point query", url)
	return nil
}

func extractId(rawurl string) (int, error) {
	parsedUrl, err := url.Parse(rawurl)
	if err != nil {
		return 0, fmt.Errorf("url %s is not valid", rawurl)
	}
	idStr, ok := parsedUrl.Query()["id"]
	if !ok {
		return 0, fmt.Errorf("url %s does not contain an id", rawurl)
	}
	id, err := strconv.Atoi(idStr[0])
	if err != nil {
		return 0, fmt.Errorf("cannot convert %s to int", idStr[0])
	}
	return id, nil
}

func TestExtractId(t *testing.T) {
	tests := []struct {
		name    string
		rawurl  string
		want    int
		wantErr bool
	}{
		{
			name:   "valid url",
			rawurl: "https://example.com/point?id=1234&detail=complet&format=geojson",
			want:   1234,
		},
		{
			name:    "url without id",
			rawurl:  "https://example.com/point?detail=complet&format=geojson",
			want:    0,
			wantErr: true,
		},
		{
			name:    "url with id not convertible to int",
			rawurl:  "https://example.com/point?id=abc&detail=complet&format=geojson",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractId(tt.rawurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	querier := UnifiedQuerier{t: t}
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

func TestGetFeature(t *testing.T) {
	querier := UnifiedQuerier{t: t}
	feature := *scrapper.GetFeature(1198, querier)
	expectedFeature := getFeatureCollectionEnrichedTest().Features[1]
	if !reflect.DeepEqual(feature, expectedFeature) {
		t.Errorf("GetFeature() = %v, want %v", feature, expectedFeature)
	}
}

func TestEnrichFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	querier := UnifiedQuerier{t: t}
	featureCollection := scrapper.GetFeatureCollection(bbox, querier)
	scrapper.EnrichFeatureCollection(featureCollection, querier)
	expectedFeatureCollection := getFeatureCollectionEnrichedTest()
	if len(featureCollection.Features) != len(expectedFeatureCollection.Features) {
		t.Errorf("EnrichFeatureCollection() has length %v, want %v", len(featureCollection.Features), len(expectedFeatureCollection.Features))
	}
	for i := range featureCollection.Features {
		if !reflect.DeepEqual(featureCollection.Features[i], expectedFeatureCollection.Features[i]) {
			t.Errorf("EnrichFeatureCollection() = %v, want %v", featureCollection.Features[i], expectedFeatureCollection.Features[i])
		}
	}
}
