import vue from "@vitejs/plugin-vue";
import httpImport from "import-http/rollup";
import { defineConfig } from "vite";

export default defineConfig({
  root: __dirname,
  build: {
    base: process.env.NODE_ENV === "production" ? "/whatchanged" : "/",
  },
  plugins: [vue(), httpImport()],
});
