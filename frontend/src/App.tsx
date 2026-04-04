import { Footer } from "./components/Footer";
import { Map } from "./components/Map";
import { Stack } from "@mui/material";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

function App() {
  const queryClient = new QueryClient();
  return (
    <QueryClientProvider client={queryClient}>
      <Stack className="h-dvh flex-1">
        <Map className="flex-1" />
        <Footer />
      </Stack>
    </QueryClientProvider>
  );
}

export default App;
