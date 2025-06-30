package tests

import (
	"fmt"
	"strings"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

func getBoundingBoxTest() scrapper.BoundingBox {
	return scrapper.BoundingBox{
		NorthEast: scrapper.Point{5.52315, 44.9159},
		SouthWest: scrapper.Point{5.49826, 44.8983},
	}
}

func getBoundingBoxStringTest() string {
	return "bbox=5.49826,44.89830,5.52315,44.91590"
}

func getFeatureTest() scrapper.Feature {
	return scrapper.Feature{
		Type:       "Feature",
		Id:         1198,
		Properties: scrapper.Properties{Name: "Fontaine du Play", Coord: scrapper.Coord{Altitude: 1670}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.51051, 44.90526}},
	}
}

func getFeatureCollectionTest() scrapper.FeatureCollection {
	var features []scrapper.Feature = []scrapper.Feature{}
	features = append(features, scrapper.Feature{
		Type:       "Feature",
		Id:         28,
		Properties: scrapper.Properties{Name: "Refuge de la Jasse du Play", Coord: scrapper.Coord{Altitude: 1629}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.5021, 44.91067}},
	})
	features = append(features, getFeatureTest())
	features = append(features, scrapper.Feature{
		Type:       "Feature",
		Id:         1199,
		Properties: scrapper.Properties{Name: "Deuxième fontaine du Play", Coord: scrapper.Coord{Altitude: 1670}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.5093, 44.9035}},
	})
	features = append(features, scrapper.Feature{
		Type:       "Feature",
		Id:         1987,
		Properties: scrapper.Properties{Name: "Rocher de Séguret", Coord: scrapper.Coord{Altitude: 2051}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.52081, 44.90792}},
	})
	features = append(features, scrapper.Feature{
		Type:       "Feature",
		Id:         1986,
		Properties: scrapper.Properties{Name: "Pas de Bèrrièves", Coord: scrapper.Coord{Altitude: 1887}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.5173, 44.90996}},
	})
	return scrapper.FeatureCollection{Features: features}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func diffLines(expected, actual string) (string, error) {
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(actual, "\n")
	diff := ""
	for i := 0; i < len(expectedLines) || i < len(actualLines); i++ {
		expectedLine := ""
		if i < len(expectedLines) {
			expectedLine = expectedLines[i]
		}
		actualLine := ""
		if i < len(actualLines) {
			actualLine = actualLines[i]
		}
		if expectedLine != actualLine {
			diff += fmt.Sprintf("< %s\n", expectedLine)
			diff += fmt.Sprintf("> %s\n", actualLine)
		}
	}
	return diff, nil
}
