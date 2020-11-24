import httpImport from "import-http/rollup";

export default {
  rollupInputOptions: {
    plugins: [httpImport()],
  },
};
