package main

import (
	"github.com/anth2o/refugenavigator/internal/scrapper"
)

func main() {
	bbox := scrapper.BoundingBox{
		NorthEast: scrapper.Point{5.52315, 44.9159},
		SouthWest: scrapper.Point{5.49826, 44.8983},
	}
	featureCollection := scrapper.GetFeatureCollection(bbox, nil)
	scrapper.EnrichFeatureCollection(featureCollection, nil)
	scrapper.ExportFeatureCollection(featureCollection, "output.gpx")
}
