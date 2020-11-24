[![Build Status](https://github.com/axetroy/changelog/workflows/ci/badge.svg)](https://github.com/axetroy/changelog/actions)
[![Build Status](https://github.com/axetroy/changelog/workflows/repl/badge.svg)](https://github.com/axetroy/changelog/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/changelog)](https://goreportcard.com/report/github.com/axetroy/changelog)
![Latest Version](https://img.shields.io/github/v/release/axetroy/changelog.svg)
![License](https://img.shields.io/github/license/axetroy/changelog.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/changelog.svg)

## changelog

Elegant changelog generator.

**Elegant**/**Simple**/**Efficient**/**Scalable**

feature:

- [x] Cross-platform support
- [x] Preset template for generation
- [x] Custom template file
- [x] Conventional Commits Parser
- [x] Generate multiple versions of change logs

Here are some changelog that have been generated

- [github.com/axetroy/vscode-deno](__test__/vscode-deno.CHANGELOG.md)
- [github.com/axetroy/changelog](CHANGELOG.md)

### Usage

```bash
$ changelog --help
changelog - a cli to generate changelog from git project

USAGE:
  changelog [OPTIONS] [version]

ARGUMENTS:
  [version]     Optional version or version range.
                1.null.
                  If you do not specify the version, then it will automatically
                  generate a change log from "HEAD~<latest version>" or
                  "HEAD~<earliest commit>" or "<latest version>-<last version>"
                2.single version. eg. v1.2.0
                  Generate a specific version of the changelog.
                3.version range. eg v1.3.0~v1.2.0
                  Generate changelog within the specified range.
                  For more details, please check the following examples.

OPTIONS:
  --help        Print help information.
  --version     Print version information.
  --dir         Specify the directory to be generated.
                The directory should contain a .git folder. defaults to $PWD.
  --file        Write output to file. default write to stdout.
  --fmt         The changelog format. Available options are "md"/"json".
                Defaults to "md".
  --preset      Cli built-in markdown template. Available options are "default".
                Only available when --fmt=md and --tpl is nil. Defaults to
                "default".
  --tpl         Specify the template file for generating. Only available when --fmt=md.

EXAMPLES:
  # generate changelog from HEAD to <latest version>. equivalent to 'changelog HEAD~tag:0'
  $ changelog

  # generate changelog of the specified version
  $ changelog v1.2.0

  # generate changelog within the specified range
  $ changelog v1.3.0~v1.2.0

  # generate changelog from HEAD to <Nth tag>
  $ changelog ~tag:0

  # generate changelog from <0th tag> to <2th tag>
  $ changelog tag:0~tag:2

  # generate changelog from HEAD to specified version
  $ changelog HEAD~v1.3.0

  # generate all changelog
  $ changelog HEAD~

  # generate changelog from two commit hashes
  $ changelog 770ed02~585445d

  # Generate changelog for the specified project
  $ changelog --dir=/path/to/project v1.0.0

SOURCE CODE:
  https://github.com/axetroy/changelog
```

### Installation

If you are using Linux/macOS. you can install it with the following command:

```shell
# install latest version
curl -fsSL https://raw.githubusercontent.com/axetroy/changelog/master/install.sh | bash
# or install specified version
curl -fsSL https://raw.githubusercontent.com/axetroy/changelog/master/install.sh | bash -s v0.1.0
# or install from gobinaries.com
curl -sf https://gobinaries.com/axetroy/changelog@v0.1.0 | sh
```

Or

Download the executable file for your platform at [release page](https://github.com/axetroy/changelog/releases)

Then set the environment variable.

eg, the executable file is in the `~/bin` directory.

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/bin"
```

then, try it out.

```bash
changelog --help
```

Finally, to use Deno correctly, you also need to set environment variables

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/.deno/bin"
```

### Build from source code

Make sure you have `Golang@v1.15.x` installed.

```shell
$ git clone https://github.com/axetroy/changelog.git $GOPATH/src/github.com/axetroy/changelog
$ cd $GOPATH/src/github.com/axetroy/changelog
$ make build
```

### Test

```bash
$ make test
```

### How it works?

[detail](HOW_IT_WORKS.md)

### License

The [MIT License](LICENSE)
