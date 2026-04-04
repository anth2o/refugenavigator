import { downloadGpx } from "../api";
import type { BoundingBox } from "../types/coordinates";
import { DownloadSuccessDialog } from "./DownloadSuccessDialog";
import "./Map.css";
import { MapControls } from "./MapControls";
import Stack from "@mui/material/Stack";
import L, { type LeafletEvent } from "leaflet";
import { latLng, Polyline } from "leaflet";
import "leaflet-draw";
import "leaflet-draw/dist/leaflet.draw.css";
import "leaflet/dist/leaflet.css";
import { useEffect, useRef, useState } from "react";

window.type = true; // https://github.com/Leaflet/Leaflet.draw/issues/1026#issuecomment-986702652

const initialCenter = latLng(44.9, 5.5);
const initialZoom = 10;

export const Map = ({ className }: { className?: string }) => {
  const [rectangle, setRectangle] = useState<Polyline | null>(null);
  const [isDrawing, setIsDrawing] = useState(false);
  const [waitingForGpx, setWaitingForGpx] = useState(false);
  const [dialogOpen, setDialogOpen] = useState(false);
  const mapRef = useRef<L.DrawMap | null>(null);
  const drawnItemsRef = useRef<L.FeatureGroup | null>(null);
  const drawControlRef = useRef<L.Control.Draw | null>(null);

  useEffect(() => {
    if (mapRef.current) return;
    const map = L.map("map").setView(initialCenter, initialZoom);
    L.tileLayer("https://{s}.tile.osm.org/{z}/{x}/{y}.png", {
      attribution:
        '&copy; <a href="https://osm.org/copyright">OpenStreetMap</a> contributors',
      referrerPolicy: "strict-origin-when-cross-origin", // https://wiki.openstreetmap.org/wiki/Referer
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
    map.on("draw:created", (e: LeafletEvent) => {
      drawnItemsRef.current!.addLayer(e.layer);
      setRectangle(e.layer);
      setIsDrawing(false);
    });
    return () => {
      map.removeControl(drawControl);
    };
  }, []);

  return (
    <>
      <Stack
        alignItems="center"
        justifyContent="center"
        className={`${className} h-full flex flex-col gap-2 md:gap-6 pt-0 md:pt-12`}
      >
        {/* https://leafletjs.com/examples/quick-start/ */}
        <div
          id="map"
          className="flex-1 w-full min-h-[300px] mx-auto max-w-[1200px] pt-6 pb-6"
        ></div>
        <MapControls
          isDrawing={isDrawing}
          isReadyForDownload={!!rectangle}
          waitingForGpx={waitingForGpx}
          onToggleDrawing={() => {
            drawnItemsRef.current!.clearLayers();
            if (rectangle) {
              setIsDrawing(false);
            } else {
              setIsDrawing(true);
              const drawHandler = new L.Draw.Rectangle(mapRef.current!);
              drawHandler.enable();
            }
            setRectangle(null);
          }}
          onDownloadGpx={async () => {
            if (!rectangle) return;
            setWaitingForGpx(true);
            const bounds = rectangle.getBounds();
            const boundingBox: BoundingBox = {
              northEast: bounds.getNorthEast(),
              southWest: bounds.getSouthWest(),
            };
            await downloadGpx(boundingBox);
            setWaitingForGpx(false);
            setDialogOpen(true);
          }}
        />
      </Stack>
      <DownloadSuccessDialog
        open={dialogOpen}
        onClose={() => setDialogOpen(false)}
      />
    </>
  );
};
