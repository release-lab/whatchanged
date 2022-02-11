const ffi = require("ffi-napi");
const path = require("path");

const lib = ffi.Library(path.join(__dirname, "lib", "whatchanged"), {
  Generate: ["string", ["string", "string"]],
});

/**
 * Generate changelog
 * @param {string} repo
 * @param {string} version
 * @returns {Promise<string>}
 */
function generate(repo, version) {
  return new Promise((resolve, reject) => {
    lib.Generate.async(repo, version, function (err, res) {
      if (err) {
        return reject(err);
      }

      const { result: changelog, error } = JSON.parse(res);
      if (error) {
        return reject(new Error(error));
      }

      resolve(changelog)
    });
  });
}

module.exports.generate = generate;
