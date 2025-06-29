package tests

import (
	"testing"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

func TestGetPoints(t *testing.T) {
	bbox := getBoundingBoxTest()
	scrapper.GetFeatureCollection(bbox)
}
