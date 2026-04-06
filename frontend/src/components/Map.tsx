import { downloadGpx, getPoints } from "../api";
import { useLeafletMap } from "../hooks/useLeafletMap";
import { rectangleToBoundingBox } from "../utils";
import { DownloadSuccessDialog } from "./DownloadSuccessDialog";
import { MapControls } from "./MapControls";
import Stack from "@mui/material/Stack";
import { LatLng, Marker, Polyline } from "leaflet";
import { useState } from "react";

export const Map = ({ className }: { className?: string }) => {
  const [waitingForGpx, setWaitingForGpx] = useState(false);
  const [dialogOpen, setDialogOpen] = useState(false);
  const { rectangle, isDrawing, toggleDrawing } = useLeafletMap(
    async (rectangle: Polyline) => {
      const points = await getPoints(rectangleToBoundingBox(rectangle));
      if (!points || !points.features) return null;
      return points.features.map((feature) => {
        const latLng = new LatLng(
          feature.geometry.coordinates[1],
          feature.geometry.coordinates[0],
        );
        return new Marker(latLng, { title: feature.properties.nom });
      });
    },
  );

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
          isDrawn={!!rectangle}
          waitingForGpx={waitingForGpx}
          onToggleDrawing={toggleDrawing}
          onDownloadGpx={async () => {
            if (!rectangle) return;
            setWaitingForGpx(true);
            await downloadGpx(rectangleToBoundingBox(rectangle));
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
