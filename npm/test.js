const path = require("path");
const { WhatchangedFFI } = require("./");

const lib = new WhatchangedFFI(path.join(__dirname, "lib", "whatchanged"));

lib
  .generate(path.join(__dirname, ".."), {
    version: ["HEAD~"],
  })
  .then((res) => {
    console.log(res);
  })
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });
