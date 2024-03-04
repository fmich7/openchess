import react from "@vitejs/plugin-react";
import { defineConfig, loadEnv } from "vite";

// https://vitejs.dev/config/
// jsxRuntime: "classic",
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd());

  const PORT = parseInt(`${env.VITE_PORT ?? "5173"}`);

  return {
    plugins: [react()],
    server: {
      port: PORT,
      watch: {
        usePolling: true,
      },
      host: true,
      strictPort: true,
    },
  };
});
