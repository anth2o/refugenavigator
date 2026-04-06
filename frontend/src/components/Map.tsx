import { downloadGpx } from "../api";
import { useLeafletMap } from "../hooks/useLeafletMap";
import type { BoundingBox } from "../types/coordinates";
import { DownloadSuccessDialog } from "./DownloadSuccessDialog";
import { MapControls } from "./MapControls";
import Stack from "@mui/material/Stack";
import { useState } from "react";

export const Map = ({ className }: { className?: string }) => {
  const [waitingForGpx, setWaitingForGpx] = useState(false);
  const [dialogOpen, setDialogOpen] = useState(false);
  const { rectangle, isDrawing, toggleDrawing } = useLeafletMap();

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
