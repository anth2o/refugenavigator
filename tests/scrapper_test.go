package tests

import (
	"os"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
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

func TestExportFeatureCollection(t *testing.T) {
	featureCollection := getFeatureCollectionTest()
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	guard := monkey.Patch(time.Now, func() time.Time {
		return time.Date(2025, 6, 29, 12, 13, 24, 0, time.UTC)
	})

	defer guard.Unpatch() // Make sure to unpatch after the test
	scrapper.ExportFeatureCollection(&featureCollection, f.Name())

	exportedGPX, err := os.ReadFile(f.Name())
	checkError(err)
	exportedGPXStr := string(exportedGPX)
	expectedExportedGPX, err := os.ReadFile("../data/example.gpx")
	checkError(err)
	expectedExportedGPXStr := string(expectedExportedGPX)
	if exportedGPXStr != expectedExportedGPXStr {
		diff, err := diffLines(expectedExportedGPXStr, exportedGPXStr)
		if err != nil {
			panic(err)
		}
		t.Errorf("ExportFeatureCollection() did not produce the expected GPX file:\n%s", diff)
	}
}
