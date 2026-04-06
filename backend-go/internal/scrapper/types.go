package scrapper

import (
	"fmt"
)

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Id         int        `json:"id"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type Properties struct {
	Name        string      `json:"nom"`
	Coord       Coord       `json:"coord"`
	Description Description `json:"description"`
	Link        string      `json:"lien,omitempty"`
}

type Description struct {
	Value string `json:"valeur"`
}

type Coord struct {
	Altitude FlexibleInt `json:"alt"` // when fetching FeatureCollection, altitude is a int, but when fetching Feature, it is a string
}

type Geometry struct {
	Type        string `json:"type"`
	Coordinates Point  `json:"coordinates"`
}

type Point [2]float64

func (p Point) Longitude() float64 {
	return p[0]
}
func (p Point) Latitude() float64 {
	return p[1]
}
func (p Point) String() string {
	return fmt.Sprintf("%.5f,%.5f", p.Longitude(), p.Latitude())
}

type BoundingBox struct {
	SouthWest, NorthEast Point
}

func (b BoundingBox) String() string {
	return fmt.Sprintf("bbox=%v,%v", b.SouthWest, b.NorthEast)
}

func (b BoundingBox) Area() float64 {
	area := (b.NorthEast.Latitude() - b.SouthWest.Latitude()) * (b.NorthEast.Longitude() - b.SouthWest.Longitude())
	return max(area, -area)
}
