import { FeatureGroup, MapContainer, TileLayer } from "react-leaflet";
import "./Map.css";
import "leaflet/dist/leaflet.css";
import "leaflet-draw/dist/leaflet.draw.css";
import { latLng } from "leaflet";
import { useState } from "react";
import { EditControl } from "react-leaflet-draw";
import Button from "@mui/material/Button";
import { downloadGpx } from "../api";
import type { BoundingBox } from "../types/coordinates";
import DownloadIcon from "@mui/icons-material/Download";
import Tooltip from "@mui/material/Tooltip";
import Stack from "@mui/material/Stack";

window.type = true; // https://github.com/Leaflet/Leaflet.draw/issues/1026#issuecomment-986702652

const center = latLng(44.9, 5.5);

export const Map = ({ className }: { className?: string }) => {
  const [selectedBoundingBox, setSelectedBoundingBox] =
    useState<BoundingBox | null>(null);
  const [waitingForGpx, setWaitingForGpx] = useState(false);

  const onCreated = (e: any) => {
    const bounds = e.layer.getBounds();
    setSelectedBoundingBox({
      northEast: bounds.getNorthEast(),
      southWest: bounds.getSouthWest(),
    });
  };
  return (
    <Stack
      alignItems="center"
      justifyContent="center"
      gap={2}
      className={className}
    >
      <p className="text-center">
        Select a rectangle on the map with the upper right button
      </p>
      <MapContainer center={center} zoom={10}>
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        <FeatureGroup>
          <EditControl
            position="topright"
            onCreated={onCreated}
            onDeleted={() => setSelectedBoundingBox(null)}
            draw={{
              rectangle: true,
              polygon: false,
              polyline: false,
              circle: false,
              marker: false,
              circlemarker: false,
            }}
          />
        </FeatureGroup>
      </MapContainer>
      <Tooltip title="Download GPX">
        <span>
          <Button
            onClick={async () => {
              setWaitingForGpx(true);
              await downloadGpx(selectedBoundingBox!);
              setWaitingForGpx(false);
            }}
            disabled={!selectedBoundingBox}
            loading={waitingForGpx}
            startIcon={<DownloadIcon />}
            variant="contained"
          >
            Download GPX
          </Button>
        </span>
      </Tooltip>
    </Stack>
  );
};
