import L, { type LeafletEvent } from "leaflet";
import "./Map.css";
import "leaflet/dist/leaflet.css";
import "leaflet-draw/dist/leaflet.draw.css";
import "leaflet-draw";
import { latLng, Polyline } from "leaflet";
import { useEffect, useRef, useState } from "react";
import Button from "@mui/material/Button";
import { downloadGpx } from "../api";
import type { BoundingBox } from "../types/coordinates";
import DownloadIcon from "@mui/icons-material/Download";
import Stack from "@mui/material/Stack";
import ClearIcon from "@mui/icons-material/Clear";
import DrawIcon from "@mui/icons-material/Draw";

window.type = true; // https://github.com/Leaflet/Leaflet.draw/issues/1026#issuecomment-986702652

const initialCenter = latLng(44.9, 5.5);
const initialZoom = 10;

export const Map = ({ className }: { className?: string }) => {
  const [rectangle, setRectangle] = useState<Polyline | null>(null);
  const [isDrawing, setIsDrawing] = useState(false);
  const [waitingForGpx, setWaitingForGpx] = useState(false);
  const mapRef = useRef<L.DrawMap | null>(null);
  const drawnItemsRef = useRef<L.FeatureGroup | null>(null);
  const drawControlRef = useRef<L.Control.Draw | null>(null);

  useEffect(() => {
    if (mapRef.current) return;
    const map = L.map("map").setView(initialCenter, initialZoom);
    L.tileLayer("http://{s}.tile.osm.org/{z}/{x}/{y}.png", {
      attribution:
        '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
    }).addTo(map);
    mapRef.current = map;

    const drawnItems = new L.FeatureGroup();
    map.addLayer(drawnItems);
    drawnItemsRef.current = drawnItems;
    const drawControl = new L.Control.Draw({
      // don't remove following options: for an unknown reason the associated buttons
      // don't appear on dev mode, but they do appear in the JS bunde if the false aren't
      // specified
      draw: {
        polyline: false,
        polygon: false,
        rectangle: false,
        circle: false,
        circlemarker: false,
        marker: false,
      },
      edit: {
        featureGroup: drawnItems,
        edit: false,
        remove: false,
      },
    });
    map.addControl(drawControl);
    drawControlRef.current = drawControl;
    map.on("draw:created", onCreateDrawing);
    return () => {
      map.removeControl(drawControl);
    };
  }, []);

  function clearDrawing() {
    setRectangle(null);
    drawnItemsRef.current!.clearLayers();
    setIsDrawing(false);
  }

  function onStartDrawing() {
    clearDrawing();
    setIsDrawing(true);
    const drawHandler = new L.Draw.Rectangle(mapRef.current!);
    drawHandler.enable();
  }

  function onCreateDrawing(e: LeafletEvent) {
    drawnItemsRef.current!.addLayer(e.layer);
    setRectangle(e.layer);
    setIsDrawing(false);
  }

  async function onDownloadGpx() {
    if (!rectangle) return;
    setWaitingForGpx(true);
    const bounds = rectangle.getBounds();
    const boundingBox: BoundingBox = {
      northEast: bounds.getNorthEast(),
      southWest: bounds.getSouthWest(),
    };
    await downloadGpx(boundingBox);
    setWaitingForGpx(false);
  }

  return (
    <Stack
      alignItems="center"
      justifyContent="center"
      gap={2}
      className={className}
    >
      {/* https://leafletjs.com/examples/quick-start/ */}
      <div id="map"></div>
      <Stack direction="row" gap={2}>
        <Button
          onClick={clearDrawing}
          variant="contained"
          disabled={!rectangle}
          color="error"
          startIcon={<ClearIcon />}
        >
          Clear
        </Button>
        <Button
          onClick={onStartDrawing}
          variant="contained"
          disabled={isDrawing}
          color="success"
          startIcon={<DrawIcon />}
        >
          Draw
        </Button>
        <Button
          onClick={onDownloadGpx}
          disabled={!rectangle}
          loading={waitingForGpx}
          startIcon={<DownloadIcon />}
          variant="contained"
        >
          Download
        </Button>
      </Stack>
    </Stack>
  );
};
