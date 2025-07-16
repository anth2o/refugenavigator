import { Map } from "./components/Map";
import { Stack } from "@mui/material";
import { Footer } from "./components/Footer";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

function App() {
  const queryClient = new QueryClient();
  return (
    <QueryClientProvider client={queryClient}>
      <Stack className="h-screen">
        <Map className="h-full" />
        <Footer />
      </Stack>
    </QueryClientProvider>
  );
}

export default App;
