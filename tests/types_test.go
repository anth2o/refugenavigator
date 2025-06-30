package tests

import (
	"testing"
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
