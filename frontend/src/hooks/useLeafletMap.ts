import "./LeafletMap.css";
import L, { latLng, Marker, Polyline, type LeafletEvent } from "leaflet";
import "leaflet-draw";
import "leaflet-draw/dist/leaflet.draw.css";
import "leaflet/dist/leaflet.css";
import { useEffect, useRef, useState } from "react";

window.type = true; // https://github.com/Leaflet/Leaflet.draw/issues/1026#issuecomment-986702652

const defaultIcon = L.icon({
  iconUrl: "https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png",
  iconRetinaUrl:
    "https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png",
  shadowUrl: "https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png",
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  shadowSize: [41, 41],
});
L.Marker.prototype.options.icon = defaultIcon;

const initialCenter = latLng(44.9, 5.5);
const initialZoom = 10;

export const useLeafletMap = (
  rectangleToMarkers: (rectangle: Polyline) => Promise<Marker[] | null>,
) => {
  const [rectangle, setRectangle] = useState<Polyline | null>(null);
  const [markers, setMarkers] = useState<Marker[] | null>(null);
  const [drawHandler, setDrawHandler] = useState<L.Draw.Rectangle | null>(null);
  const mapRef = useRef<L.DrawMap | null>(null);
  const drawnItemsRef = useRef<L.FeatureGroup | null>(null);
  const drawControlRef = useRef<L.Control.Draw | null>(null);

  useEffect(() => {
    if (mapRef.current) return;
    const map = L.map("map").setView(initialCenter, initialZoom);
    L.tileLayer("https://{s}.tile.opentopomap.org/{z}/{x}/{y}.png", {
      attribution:
        '&copy; <a href="https://opentopomap.org/about#roadmap">OpenTopoMap</a> contributors',
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
    map.on("draw:created", async (e: LeafletEvent) => {
      const rectangle: Polyline = e.layer;
      drawnItemsRef.current!.addLayer(rectangle);
      setRectangle(rectangle);
      const markers = await rectangleToMarkers(rectangle);
      if (markers) {
        markers.forEach((marker) => marker.addTo(map));
        setMarkers(markers);
      }

      setDrawHandler(null);
    });
    return () => {
      map.removeControl(drawControl);
    };
  }, []);

  const toggleDrawing = () => {
    drawnItemsRef.current!.clearLayers();
    if (drawHandler) {
      drawHandler.disable();
      setDrawHandler(null);
    } else if (!rectangle) {
      const drawHandler = new L.Draw.Rectangle(mapRef.current!);
      setDrawHandler(drawHandler);
      drawHandler.enable();
    }
    if (markers) {
      markers.forEach((marker) => marker.remove());
      setMarkers(null);
    }
    setRectangle(null);
  };

  const isDrawing = !!drawHandler;

  return {
    rectangle,
    isDrawing,
    toggleDrawing,
  };
};
