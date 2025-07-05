import { Map } from "./components/Map";
import { Stack } from "@mui/material";
import { Footer } from "./components/Footer";

function App() {
  return (
    <Stack className="h-screen">
      <Map className="h-full" />
      <Footer />
    </Stack>
  );
}

export default App;
