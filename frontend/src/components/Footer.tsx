import GitHubIcon from "@mui/icons-material/GitHub";
import { Link, Stack } from "@mui/material";
export const Footer = () => {
  return (
    <footer className="text-center text-gray-500 pb-4 px-20">
      <Stack gap={1}>
        <Stack
          direction="row"
          alignItems="center"
          justifyContent="center"
          gap={2}
        >
          © {new Date().getFullYear()} Refuge Navigator
          <Link
            href="https://github.com/anth2o/refugenavigator"
            color="inherit"
          >
            <GitHubIcon />
          </Link>
        </Stack>
        <p className="text-sm">
          The data provided by Refuge Navigator comes from{" "}
          <Link href="https://refuges.info">refuges.info</Link>, is attributed
          to "©Les contributeurs de Refuges.info" and licensed under the{" "}
          <Link href="https://creativecommons.org/licenses/by-sa/2.0/">
            CC BY-SA 2.0
          </Link>
          .
        </p>
      </Stack>
    </footer>
  );
};
