// hand written code that should ideally be generated from backend-go/internal/server/api.go
import type { BoundingBox } from "./types/coordinates";
import type { FeatureCollection } from "./types/points";
import { boundingBoxToQueryParams } from "./utils";
import type { AxiosResponse } from "axios";
import axios from "axios";

function BoundingBoxToFileName(boundingBox: BoundingBox): string {
  return `refugenavigator_export_${boundingBox.southWest.lat.toFixed(3)}_${boundingBox.southWest.lng.toFixed(3)}_${boundingBox.northEast.lat.toFixed(3)}_${boundingBox.northEast.lng.toFixed(3)}.gpx`;
}

function getApiUrl(): string {
  const baseUrl =
    import.meta.env.MODE === "development" ? "http://127.0.0.1:8080" : "";
  return baseUrl + "/api";
}

export async function getPoints(
  boundingBox: BoundingBox,
): Promise<FeatureCollection> {
  const response = await axios.get<FeatureCollection>(
    `${getApiUrl()}/points?${boundingBoxToQueryParams(boundingBox)}`,
    {
      headers: {
        Accept: "application/json",
      },
    },
  );
  return response.data;
}

export async function downloadGpx(boundingBox: BoundingBox): Promise<void> {
  const response: AxiosResponse<Blob> = await axios.get(
    `${getApiUrl()}/gpx?${boundingBoxToQueryParams(boundingBox)}`,
    {
      responseType: "blob",
      headers: {
        Accept: "application/gpx+xml",
      },
    },
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

export async function getGitTag(): Promise<string> {
  const response = await axios.get<{ tag: string }>(`${getApiUrl()}/git-tag`);
  return response.data.tag;
}
