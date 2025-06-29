package tests

import "github.com/anth2o/refugenavigator/internal/scrapper"

func getBoundingBoxTest() scrapper.BoundingBox {
	return scrapper.BoundingBox{
		NorthEast: scrapper.Point{5.52315, 44.9159},
		SouthWest: scrapper.Point{5.49826, 44.8983},
	}
}
