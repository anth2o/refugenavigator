package tests

import (
	"testing"

	"github.com/anth2o/refugenavigator/internal/scrapper"
	"github.com/google/go-cmp/cmp"
)

func TestGetFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	featureCollection := scrapper.GetFeatureCollection(bbox)
	expectedFeatureCollection := getFeatureCollectionTest()
	diff := cmp.Diff(featureCollection, expectedFeatureCollection)
	if diff != "" {
		t.Errorf("GetFeatureCollection() did not produce the expected FeatureCollection:\n%s", diff)
	}
}

func TestGetFeature(t *testing.T) {
	feature := *scrapper.GetFeature(1198)
	expectedFeature := getFeatureCollectionEnrichedTest().Features[1]
	diff := cmp.Diff(feature, expectedFeature)
	if diff != "" {
		t.Errorf("GetFeature() did not produce the expected Feature:\n%s", diff)
	}
}

func TestEnrichFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	featureCollection := scrapper.GetFeatureCollection(bbox)
	scrapper.EnrichFeatureCollection(featureCollection)
	expectedFeatureCollection := getFeatureCollectionEnrichedTest()
	diff := cmp.Diff(featureCollection, expectedFeatureCollection)
	if diff != "" {
		t.Errorf("EnrichFeatureCollection() did not produce the expected FeatureCollection:\n%s", diff)
	}
}
