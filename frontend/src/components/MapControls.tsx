import ClearIcon from "@mui/icons-material/Clear";
import DownloadIcon from "@mui/icons-material/Download";
import DrawIcon from "@mui/icons-material/Draw";
import IconButton from "@mui/material/IconButton";
import Stack from "@mui/material/Stack";

interface MapControlsProps {
  isDrawing: boolean;
  isReadyForDownload: boolean;
  waitingForGpx: boolean;
  onToggleDrawing: () => void;
  onDownloadGpx: () => void;
}

export const MapControls = ({
  isDrawing,
  isReadyForDownload,
  waitingForGpx,
  onToggleDrawing,
  onDownloadGpx,
}: MapControlsProps) => {
  return (
    <Stack direction="row" gap={2}>
      <IconButton
        onClick={onToggleDrawing}
        disabled={waitingForGpx || isDrawing}
        color={isReadyForDownload ? "error" : "success"}
        size="large"
      >
        {isReadyForDownload ? <ClearIcon /> : <DrawIcon />}
      </IconButton>
      <IconButton
        onClick={onDownloadGpx}
        disabled={!isReadyForDownload}
        loading={waitingForGpx}
        color="primary"
        size="large"
      >
        <DownloadIcon />
      </IconButton>
    </Stack>
  );
};
