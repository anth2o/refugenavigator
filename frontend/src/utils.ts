import type { BoundingBox } from "./types/coordinates";
import { Polyline } from "leaflet";

export function rectangleToBoundingBox(rectangle: Polyline): BoundingBox {
  const bounds = rectangle.getBounds();
  return {
    northEast: bounds.getNorthEast(),
    southWest: bounds.getSouthWest(),
  };
}
export function boundingBoxToQueryParams(boundingBox: BoundingBox): string {
  return `SouthWest.Latitude=${boundingBox.southWest.lat}&SouthWest.Longitude=${boundingBox.southWest.lng}&NorthEast.Latitude=${boundingBox.northEast.lat}&NorthEast.Longitude=${boundingBox.northEast.lng}`;
}
