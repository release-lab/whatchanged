const { generate } = require("./index");

generate(path.join(__dirname, ".."), "HEAD~")
  .then((res) => {
    console.log(res);
  })
  .catch((err) => {
    console.error(err);
  });
