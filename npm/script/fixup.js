const path = require("path");
const fs = require("fs");

function write(filepath, content) {
  fs.writeFileSync(filepath, JSON.stringify(content, null, 2));
}

write(path.join(__dirname, "..", "output", "cjs", "package.json"), {
  type: "commonjs",
});

write(path.join(__dirname, "..", "output", "mjs", "package.json"), {
  type: "module",
});
