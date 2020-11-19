[![Build Status](https://github.com/axetroy/go-cli-boilerplate/workflows/ci/badge.svg)](https://github.com/axetroy/go-cli-boilerplate/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/go-cli-boilerplate)](https://goreportcard.com/report/github.com/axetroy/go-cli-boilerplate)
![Latest Version](https://img.shields.io/github/v/release/axetroy/go-cli-boilerplate.svg)
![License](https://img.shields.io/github/license/axetroy/go-cli-boilerplate.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/go-cli-boilerplate.svg)

## go-cli-boilerplate

> go-cli-boilerplate

### Usage

```bash

```

### Installation

If you are using Linux/macOS. you can install it with the following command:

```shell
# install latest version
curl -fsSL https://raw.githubusercontent.com/axetroy/go-cli-boilerplate/master/install.sh | bash
# or install specified version
curl -fsSL https://raw.githubusercontent.com/axetroy/go-cli-boilerplate/master/install.sh | bash -s v1.3.0
# or install from gobinaries.com
curl -sf https://gobinaries.com/axetroy/go-cli-boilerplate@v1.3.0 | sh
```

Or

Download the executable file for your platform at [release page](https://github.com/axetroy/go-cli-boilerplate/releases)

Then set the environment variable.

eg, the executable file is in the `~/bin` directory.

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/bin"
```

then, try it out.

```bash
go-cli-boilerplate --help
```

Finally, to use Deno correctly, you also need to set environment variables

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/.deno/bin"
```

### Build from source code

Make sure you have `Golang@v1.15.x` installed.

```shell
$ git clone https://github.com/axetroy/go-cli-boilerplate.git $GOPATH/src/github.com/axetroy/go-cli-boilerplate
$ cd $GOPATH/src/github.com/axetroy/go-cli-boilerplate
$ make build
```

### Test

```bash
$ make test
```

### License

The [MIT License](LICENSE)
