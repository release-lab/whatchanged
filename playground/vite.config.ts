import vue from "@vitejs/plugin-vue";
import httpImport from "import-http/rollup";
import { defineConfig } from "vite";

export default defineConfig({
  base: process.env.NODE_ENV === "production" ? "/whatchanged" : "/",
  root: __dirname,
  plugins: [vue(), httpImport()],
});
