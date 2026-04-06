package tests

import (
	"testing"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

func TestString(t *testing.T) {
	bbox := getBoundingBoxTest()
	southWest := bbox.SouthWest
	southWestString := southWest.String()
	expectedSouthWestString := "5.49826,44.89830"
	if southWestString != expectedSouthWestString {
		t.Errorf("Point.String() = %s, want %s", southWestString, expectedSouthWestString)
	}
	northEast := bbox.NorthEast
	northEastString := northEast.String()
	expectedNorthEastString := "5.52315,44.91590"
	if northEastString != expectedNorthEastString {
		t.Errorf("Point.String() = %s, want %s", northEastString, expectedNorthEastString)
	}
	bboxString := bbox.String()
	expectedBboxString := getBoundingBoxStringTest()
	if bboxString != expectedBboxString {
		t.Errorf("BoundingBox.String() = %s, want %s", bboxString, expectedBboxString)
	}
}

func TestArea(t *testing.T) {
	bbox := getBoundingBoxTest()
	area := bbox.Area()
	expectedArea := getBoundingBoxAreaTest()
	if (area-expectedArea)*(area-expectedArea) >= 0.000000000001 {
		t.Errorf("BoundingBox.Area() = %f, want %f", area, expectedArea)
	}

	reversedBbox := scrapper.BoundingBox{
		NorthEast: scrapper.Point{bbox.SouthWest.Longitude(), bbox.NorthEast.Latitude()},
		SouthWest: scrapper.Point{bbox.NorthEast.Longitude(), bbox.SouthWest.Latitude()},
	}
	area = reversedBbox.Area()
	expectedArea = getBoundingBoxAreaTest()
	if (area-expectedArea)*(area-expectedArea) >= 0.000000000001 {
		t.Errorf("BoundingBox.Area() = %f, want %f", area, expectedArea)
	}

	emptyBox := scrapper.BoundingBox{
		NorthEast: scrapper.Point{bbox.SouthWest.Longitude(), bbox.NorthEast.Latitude()},
		SouthWest: scrapper.Point{bbox.SouthWest.Longitude(), bbox.SouthWest.Latitude()},
	}
	area = emptyBox.Area()
	expectedArea = 0
	if area != expectedArea {
		t.Errorf("BoundingBox.Area() = %f, want %f", area, expectedArea)
	}
}
