### [whatchanged] binding

### Usage

```bash
npm install @axetroy/whatchanged --save-exact
```

#### CommonJS

```js
const { WhatchangedFFI } = require("@axetroy/whatchanged");

const lib = new WhatchangedFFI("/path/to/whatchanged");

lib
  .generate("/path/to/git/project/or/url", {
    version: ["HEAD~"],
  })
  .then((changelog) => {
    console.log(changelog);
  })
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });
```

#### ES module

```js
import path from "path";
import { WhatchangedFFI } from "@axetroy/whatchanged";

const lib = new WhatchangedFFI("/path/to/whatchanged");

lib
  .generate("/path/to/git/project/or/url", {
    version: ["HEAD~"],
  })
  .then((changelog) => {
    console.log(changelog);
  })
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });
```
