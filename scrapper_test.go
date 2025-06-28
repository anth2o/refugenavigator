package refugenavigator

import (
	"testing"
)

func TestGetPoints(t *testing.T) {
	bbox := getBoundingBoxTest()
	getFeatureCollection(bbox)
}

func TestMain(t *testing.T) {
	main()
}
