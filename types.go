package refugenavigator

import (
	"fmt"
	"time"

	"github.com/twpayne/go-gpx"
)

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

func (fc FeatureCollection) ToGpx() *gpx.GPX {
	exportedGpx := gpx.GPX{
		Version: "1.1",
		Creator: "Refuges Scrapper",
		Metadata: &gpx.MetadataType{
			Desc: "grp_rep:refuges.info " + time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	for _, feature := range fc.Features {
		exportedGpx.Wpt = append(exportedGpx.Wpt, feature.ToGpx())
	}
	return &exportedGpx
}

type Feature struct {
	Type       string     `json:"type"`
	Id         int        `json:"id"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

func (f Feature) ToGpx() *gpx.WptType {
	return &gpx.WptType{
		Lat:  f.Geometry.Coordinates.latitude(),
		Lon:  f.Geometry.Coordinates.longitude(),
		Name: f.Properties.Name,
		// TODO: summarize the data from the web page relative to this feature into a description
		Desc: "Imported from Refuges.info on " + time.Now().Format("2006-01-02 15:04:05"),
	}
}

type Properties struct {
	Name  string `json:"nom"`
	Coord Coord  `json:"coord"`
}

type Coord struct {
	Altitude float64 `json:"alt"`
}

type Geometry struct {
	Type        string `json:"type"`
	Coordinates Point  `json:"coordinates"`
}

type Point [2]float64

func (p Point) longitude() float64 {
	return p[0]
}
func (p Point) latitude() float64 {
	return p[1]
}
func (p Point) String() string {
	return fmt.Sprintf("%.5f,%.5f", p.longitude(), p.latitude())
}

type BoundingBox struct {
	southWest, northEast Point
}

func (b BoundingBox) String() string {
	return fmt.Sprintf("bbox=%v,%v", b.southWest, b.northEast)
}
