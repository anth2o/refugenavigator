import axios from "axios";
import type { AxiosResponse } from "axios";
import { API_CONFIG } from "./config";
import type { BoundingBox } from "./types/coordinates";

function boundingBoxToQueryParams(boundingBox: BoundingBox): string {
  return `SouthWest.Latitude=${boundingBox.southWest.lat}&SouthWest.Longitude=${boundingBox.southWest.lng}&NorthEast.Latitude=${boundingBox.northEast.lat}&NorthEast.Longitude=${boundingBox.northEast.lng}`;
}

export async function downloadGpx(
  boundingBox: BoundingBox,
  filename: string = "route.gpx"
): Promise<void> {
  const response: AxiosResponse<Blob> = await axios.get(
    `${API_CONFIG.baseUrl}/api/gpx?${boundingBoxToQueryParams(boundingBox)}`,
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
  a.download = filename;
  document.body.appendChild(a);
  a.click();
  window.URL.revokeObjectURL(url);
  document.body.removeChild(a);
}
