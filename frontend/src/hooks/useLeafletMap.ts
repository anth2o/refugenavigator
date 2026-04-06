import "./LeafletMap.css";
import L, { type LeafletEvent } from "leaflet";
import { latLng, Polyline } from "leaflet";
import "leaflet-draw";
import "leaflet-draw/dist/leaflet.draw.css";
import "leaflet/dist/leaflet.css";
import { useEffect, useRef, useState } from "react";

window.type = true; // https://github.com/Leaflet/Leaflet.draw/issues/1026#issuecomment-986702652

const initialCenter = latLng(44.9, 5.5);
const initialZoom = 10;

export const useLeafletMap = () => {
  const [rectangle, setRectangle] = useState<Polyline | null>(null);
  const [drawHandler, setDrawHandler] = useState<L.Draw.Rectangle | null>(null)
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
      setDrawHandler(null);
    });
    return () => {
      map.removeControl(drawControl);
    };
  }, []);

  const toggleDrawing = () => {
    drawnItemsRef.current!.clearLayers();
    if (drawHandler) {
      drawHandler.disable()
      setDrawHandler(null);
    } else if (!rectangle) {
      const drawHandler = new L.Draw.Rectangle(mapRef.current!);
      setDrawHandler(drawHandler)
      drawHandler.enable();
    }
    setRectangle(null);
  };
  const isDrawing = !!drawHandler

  return {
    rectangle,
    isDrawing,
    toggleDrawing,
  };
};
