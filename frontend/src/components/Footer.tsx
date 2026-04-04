import { getGitTag } from "../api";
import GitHubIcon from "@mui/icons-material/GitHub";
import { Link, Stack } from "@mui/material";
import { useQuery } from "@tanstack/react-query";

export const Footer = () => {
  const { data: gitTag } = useQuery<string | undefined>({
    queryKey: ["git-tag"],
    queryFn: getGitTag,
  });
  return (
    <footer className="text-center text-gray-500 pt-3 pb-1 px-1 text-xs flex-wrap">
      <Stack gap={0.5}>
        <p>
          {" "}
          © {new Date().getFullYear()} Refuge Navigator
          {gitTag ? ` (${gitTag})` : ""}
        </p>

        <Link href="https://github.com/anth2o/refugenavigator" color="inherit">
          <GitHubIcon />
        </Link>
        <p className="text-[10px] md:text-xs wrap-break-word">
          The data provided by Refuge Navigator comes from{" "}
          <Link href="https://refuges.info">refuges.info</Link>, is attributed
          to "Les contributeurs de Refuges.info" and licensed under the{" "}
          <Link href="https://creativecommons.org/licenses/by-sa/2.0/">
            CC BY-SA 2.0
          </Link>
          .
        </p>
      </Stack>
    </footer>
  );
};
