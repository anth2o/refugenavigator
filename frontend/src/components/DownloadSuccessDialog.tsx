import CheckIcon from "@mui/icons-material/Check";
import {
  Alert,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
} from "@mui/material";

interface DownloadSuccessDialogProps {
  open: boolean;
  onClose: () => void;
}

export const DownloadSuccessDialog = ({
  open,
  onClose,
}: DownloadSuccessDialogProps) => {
  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>GPX Download Successful</DialogTitle>
      <DialogContent>
        <Alert
          severity="success"
          icon={<CheckIcon fontSize="inherit" />}
          className="mt-1"
        >
          Your GPX was successfully downloaded. You can now go to your downloads
          and open it with your favorite GPX viewer.
        </Alert>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose} color="primary">
          OK
        </Button>
      </DialogActions>
    </Dialog>
  );
};
