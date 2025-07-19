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
	Comments   []Comment  `json:"comments"`
}

type Properties struct {
	Name        string      `json:"nom"`
	Coord       Coord       `json:"coord"`
	Description Description `json:"description"`
	Link        string      `json:"lien,omitempty"`
	Type        Type        `json:"type"`
	Remarque    Remarque    `json:"remarque,omitempty"`
	Acces       Acces       `json:"acces,omitempty"`
}

type Type struct {
	Id     FlexibleInt `json:"id"`
	Valeur string      `json:"valeur"`
	Icone  string      `json:"icone"`
}

type Remarque struct {
	Nom    string `json:"nom"`
	Valeur string `json:"valeur"`
}
type Acces struct {
	Nom    string `json:"nom"`
	Valeur string `json:"valeur"`
}
type Description struct {
	Valeur string `json:"valeur"`
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

type Comment struct {
	ID        string `json:"id"`
	Date      string `json:"date"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	PhotoURL  string `json:"photo_url,omitempty"`
	PhotoDate string `json:"photo_date,omitempty"`
}

type BoundingBox struct {
	SouthWest, NorthEast Point
}

func (b BoundingBox) String() string {
	return fmt.Sprintf("bbox=%v,%v", b.SouthWest, b.NorthEast)
}
