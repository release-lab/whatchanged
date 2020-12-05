[![Build Status](https://github.com/axetroy/whatchanged/workflows/ci/badge.svg)](https://github.com/axetroy/whatchanged/actions)
[![Build Status](https://github.com/axetroy/whatchanged/workflows/playground/badge.svg)](https://github.com/axetroy/whatchanged/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/whatchanged)](https://goreportcard.com/report/github.com/axetroy/whatchanged)
![Latest Version](https://img.shields.io/github/v/release/axetroy/whatchanged.svg)
![License](https://img.shields.io/github/license/axetroy/whatchanged.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/whatchanged.svg)

## whatchanged

An elegant changelog generator.

Focus on **Elegant**/**Simple**/**Efficient**/**Scalable**

[Feel the magic online](https://axetroy.github.io/whatchanged/)

Feature:

- [x] Cross-platform support
- [x] Preset template for generation
- [x] Custom template file
- [x] Conventional Commits Parser
- [x] Generate multiple versions of change logs
- [x] [Github Action](https://github.com/axetroy/setup-whatchanged)

### Usage

```bash
$ whatchanged --help
whatchanged - a cli to generate changelog from git project

USAGE:
  whatchanged [OPTIONS] [version...]

ARGUMENTS:
  [version...]  Optional version or version range.
                1.null.
                  If you do not specify the version, then it will automatically
                  generate a change log from "HEAD~<latest version>" or
                  "HEAD~<earliest commit>" or "<latest version>-<last version>"
                2.single version. eg. "v1.2.0"
                  Generate a specific version of the changelog.
                3.multiple versions. eg. "v2.0.0 v1.0.0"
                4.version range. eg v1.3.0~v1.2.0
                  Generate changelog within the specified range.
                  For more details, please check the following examples.

OPTIONS:
  --help        Print help information.
  --version     Print version information.

  --project     Specify the project to be generated. It can be a relative path.
                or an absolute path or even a remote Git URL. eg.
                --project=/path/to/project/which/contains/.git/folder
                --project=https://github.com/axetroy/whatchanged.git
                Defaults to "--project=$PWD".
  --output      Write output to file. default write to stdout.

  --fmt         Specify the changelog format. Available options:
                --fmt=md
                --fmt=json
                Defaults to "--fmt=md".
  --preset      Cli built-in markdown template. Available options:
                --preset=default
                --preset=full
                Only available when --fmt=md and --tpl is nil.
                Defaults to "--preset=default".
  --tpl         Specify the template file for generating. Only available when
                --fmt=md.

EXAMPLES:
  # generate changelog from HEAD to <latest version>.
  # if HEAD is not the latest tag. then this should be a unreleased version
  # otherwise it should be the latest version
  $ whatchanged

  # generate changelog of the specified version
  $ whatchanged v1.2.0

  # Generate the specified two versions
  # Separate by a comma, and only generate these two versions
  # the middle version will not be generated
  $ whatchanged v2.0.0 v1.0.0

  # generate HEAD to latest tag and <Nth tag>
  $ whatchanged HEAD~@0 @1 @2

  # generate changelog within the specified range
  $ whatchanged v1.3.0~v1.2.0

  # generate changelog from HEAD to <Nth tag>
  $ whatchanged ~@0

  # generate changelog from <0th tag> to <2th tag>
  $ whatchanged @0~@2

  # generate changelog from HEAD to specified version
  $ whatchanged HEAD~v1.3.0

  # generate all changelog
  $ whatchanged HEAD~

  # generate changelog from two commit hashes
  $ whatchanged 770ed02~585445d

  # Generate changelog for the specified project
  $ whatchanged --project=/path/to/project v1.0.0

  # Generate changelog for the remote project
  $ whatchanged --project=https://github.com/axetroy/whatchanged.git v0.1.0

SOURCE CODE:
  https://github.com/axetroy/whatchanged
```

### Installation

#### Install via npm

```bash
# install in global
npm install -g whatchanged
# run the command once
npx whatchanged
```

#### Install via shell

If you are using Linux/macOS. you can install it with the following command:

```shell
# install latest version
curl -fsSL https://raw.githubusercontent.com/axetroy/whatchanged/master/install.sh | bash
# or install specified version
curl -fsSL https://raw.githubusercontent.com/axetroy/whatchanged/master/install.sh | bash -s v0.2.1
# or install from gobinaries.com
curl -sf https://gobinaries.com/axetroy/whatchanged@v0.2.2 | sh
```

#### Install from Github release page

Download the executable file for your platform at [release page](https://github.com/axetroy/whatchanged/releases)

Then set the environment variable.

eg, the executable file is in the `~/bin` directory.

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/bin"
```

then, try it out.

```bash
$ whatchanged --help
```

Finally, to use Deno correctly, you also need to set environment variables

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/.deno/bin"
```

### Build from source code

Make sure you have `Golang@v1.15.x` installed.

```shell
$ git clone https://github.com/axetroy/whatchanged.git $GOPATH/src/github.com/axetroy/whatchanged
$ cd $GOPATH/src/github.com/axetroy/whatchanged
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
