package tests

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/anth2o/refugenavigator/internal/scrapper"
)

func checkExport(featureCollection *scrapper.FeatureCollection, t *testing.T, expectedFile string) {
	var f *os.File
	var err error
	expectedFile = "../data/" + expectedFile
	update := false
	if update {
		f, err = os.Create(expectedFile)
		t.Errorf("Don't let update to true, and check the differences to %s before committing", expectedFile)
	} else {
		f, err = os.CreateTemp("", "sample")
	}
	defer f.Close()
	if err != nil {
		panic(err)
	}
	guard := monkey.Patch(time.Now, func() time.Time {
		return time.Date(2025, 6, 29, 12, 13, 24, 0, time.UTC)
	})

	defer guard.Unpatch() // Make sure to unpatch after the test
	exportedGPX, err := scrapper.ExportFeatureCollection(featureCollection)
	checkError(err)
	exportedGPXStr := string(exportedGPX)

	expectedExportedGPX, err := os.ReadFile(expectedFile)
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
func TestExportFeatureCollection(t *testing.T) {
	featureCollection := getFeatureCollectionTest()
	checkExport(featureCollection, t, "exported.gpx")
}
func TestExportFeatureCollectionEnriched(t *testing.T) {
	featureCollection := getFeatureCollectionEnrichedTest()
	checkExport(featureCollection, t, "exported_enriched.gpx")
}
