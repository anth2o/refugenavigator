import ClearIcon from "@mui/icons-material/Clear";
import DownloadIcon from "@mui/icons-material/Download";
import DrawIcon from "@mui/icons-material/Draw";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";

interface MapControlsProps {
  isDrawing: boolean;
  isDrawn: boolean;
  waitingForGpx: boolean;
  onToggleDrawing: () => void;
  onDownloadGpx: () => void;
}

export const MapControls = ({
  isDrawing,
  isDrawn,
  waitingForGpx,
  onToggleDrawing,
  onDownloadGpx,
}: MapControlsProps) => {
  return (
    <Stack direction="row" gap={2}>
      <Button
        onClick={onToggleDrawing}
        color={isDrawn || isDrawing ? "error" : "success"}
        size="large"
        endIcon={isDrawn || isDrawing ? <ClearIcon /> : <DrawIcon />}
        variant="contained"
      >
        {isDrawn ? "Clear" : (isDrawing ? "Cancel" : "Draw")}
      </Button>
      <Button
        onClick={onDownloadGpx}
        disabled={!isDrawn}
        loading={waitingForGpx}
        color="primary"
        size="large"
        startIcon={<DownloadIcon />}
        variant="contained"
      >
        Download
      </Button>
    </Stack>
  );
};
