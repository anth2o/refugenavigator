import axios from "axios";
import type { AxiosResponse } from "axios";
import type { BoundingBox } from "./types/coordinates";

function boundingBoxToQueryParams(boundingBox: BoundingBox): string {
  return `SouthWest.Latitude=${boundingBox.southWest.lat}&SouthWest.Longitude=${boundingBox.southWest.lng}&NorthEast.Latitude=${boundingBox.northEast.lat}&NorthEast.Longitude=${boundingBox.northEast.lng}`;
}

function BoundingBoxToFileName(boundingBox: BoundingBox): string {
  return `refugenavigator_export_${boundingBox.southWest.lat.toFixed(3)}_${boundingBox.southWest.lng.toFixed(3)}_${boundingBox.northEast.lat.toFixed(3)}_${boundingBox.northEast.lng.toFixed(3)}.gpx`;
}

export async function downloadGpx(boundingBox: BoundingBox): Promise<void> {
  let baseUrl = "";
  if (import.meta.env.MODE === "development") {
    baseUrl = "http://127.0.0.1:8080";
  }
  const response: AxiosResponse<Blob> = await axios.get(
    `${baseUrl}/api/gpx?${boundingBoxToQueryParams(boundingBox)}`,
    {
      responseType: "blob",
      headers: {
        Accept: "application/gpx+xml",
      },
    }
  );

  // Create a blob URL and trigger download
  const blob = new Blob([response.data], { type: "application/gpx+xml" });
  const url = window.URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = BoundingBoxToFileName(boundingBox);
  document.body.appendChild(a);
  a.click();
  window.URL.revokeObjectURL(url);
  document.body.removeChild(a);
}
