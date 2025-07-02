import { FeatureGroup, MapContainer, TileLayer } from "react-leaflet";
import "./Map.css";
import "leaflet/dist/leaflet.css";
import "leaflet-draw/dist/leaflet.draw.css";
import { latLng } from "leaflet";
import { useState } from "react";
import { EditControl } from "react-leaflet-draw";

window.type = true; // https://github.com/Leaflet/Leaflet.draw/issues/1026#issuecomment-986702652

const center = latLng(44.9, 5.5);

type BoundingBox = {
  northEast: L.LatLng;
  southWest: L.LatLng;
};
export const Map = () => {
  const [selectedArea, setSelectedArea] = useState<BoundingBox | null>(null);

  const onCreated = (e: any) => {
    const bounds = e.layer.getBounds();
    setSelectedArea({
      northEast: bounds.getNorthEast(),
      southWest: bounds.getSouthWest(),
    });
  };
  return (
    <>
      <MapContainer center={center} zoom={10}>
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        <FeatureGroup>
          <EditControl
            position="topright"
            onCreated={onCreated}
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
      {selectedArea && (
        <>
          <div>
            {selectedArea.northEast.lat} {selectedArea.northEast.lng}
          </div>
          <div>
            {selectedArea.southWest.lat} {selectedArea.southWest.lng}
          </div>
        </>
      )}
    </>
  );
};
