package tests

import (
	"os"
	"reflect"
	"testing"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

type TestQuerier struct {
	t *testing.T
}

func (q TestQuerier) QueryUrl(url string) []byte {
	data, err := os.ReadFile("../data/example.json")
	if err != nil {
		q.t.Errorf("Error reading example.json: %v", err)
		return nil
	}
	return data
}

func TestGetFeatureCollection(t *testing.T) {
	bbox := getBoundingBoxTest()
	querier := TestQuerier{t: t}
	featureCollection := scrapper.GetFeatureCollection(bbox, querier)
	var features []scrapper.Feature = []scrapper.Feature{}
	features = append(features, scrapper.Feature{
		Type:       "Feature",
		Id:         28,
		Properties: scrapper.Properties{Name: "Refuge de la Jasse du Play", Coord: scrapper.Coord{Altitude: 1629}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.5021, 44.91067}},
	})
	features = append(features, scrapper.Feature{
		Type:       "Feature",
		Id:         1198,
		Properties: scrapper.Properties{Name: "Fontaine du Play", Coord: scrapper.Coord{Altitude: 1670}},
		Geometry:   scrapper.Geometry{Type: "Point", Coordinates: scrapper.Point{5.51051, 44.90526}},
	})
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
	expectedFeatureCollection := scrapper.FeatureCollection{Features: features}

	if len(featureCollection.Features) != len(expectedFeatureCollection.Features) {
		t.Errorf("GetFeatureCollection() has length %v, want %v", len(featureCollection.Features), len(expectedFeatureCollection.Features))
	}
	for i := range featureCollection.Features {
		if !reflect.DeepEqual(featureCollection.Features[i], expectedFeatureCollection.Features[i]) {
			t.Errorf("GetFeatureCollection() = %v, want %v", featureCollection.Features[i], expectedFeatureCollection.Features[i])
		}
	}
}
