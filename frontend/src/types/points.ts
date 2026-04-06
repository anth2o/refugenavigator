// hand written code that should ideally be generated from backend-go/internal/scrapper/types.go

export interface FeatureCollection {
  features: Feature[];
}

export interface Feature {
  id: number;
  properties: Properties;
  geometry: Geometry;
}

export interface Properties {
  nom: string;
  coord: Coord;
  description: Description;
  lien?: string;
}

export interface Description {
  valeur: string;
}

export interface Coord {
  alt: number | string;
}

export interface Geometry {
  type: string;
  coordinates: Point;
}

export type Point = [number, number];
