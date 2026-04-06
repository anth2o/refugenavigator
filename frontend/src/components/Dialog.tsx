import CheckIcon from "@mui/icons-material/Check";
import ErrorIcon from "@mui/icons-material/Error";
import {
  Alert,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
} from "@mui/material";

interface DialogProps {
  open: boolean;
  onClose: () => void;
  message?: string | null;
}

export const DownloadSuccessDialog = ({ open, onClose }: DialogProps) => {
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

export const ErrorDialog = ({ open, onClose, message }: DialogProps) => {
  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>An error occured</DialogTitle>
      <DialogContent>
        <Alert
          severity="error"
          icon={<ErrorIcon fontSize="inherit" />}
          className="mt-1"
        >
          {message}
        </Alert>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose} color="primary">
          Close
        </Button>
      </DialogActions>
    </Dialog>
  );
};
