English | [中文简体](README_zh-CN.md)

[![Build Status](https://github.com/whatchanged-community/whatchanged/workflows/ci/badge.svg)](https://github.com/whatchanged-community/whatchanged/actions)
[![Build Status](https://github.com/whatchanged-community/whatchanged/workflows/playground/badge.svg)](https://github.com/whatchanged-community/whatchanged/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/whatchanged-community/whatchanged)](https://goreportcard.com/report/github.com/whatchanged-community/whatchanged)
![Latest Version](https://img.shields.io/github/v/release/whatchanged-community/whatchanged.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/whatchanged-community/whatchanged.svg)

## whatchanged

An elegant changelog generator.

Focus on **Elegant**/**Simple**/**Efficient**/**Scalable**

[Feel the magic online](https://whatchanged-community.github.io/whatchanged/)

Feature:

- [x] Cross-platform support
- [x] Generation for local/Remote git repositories
- [x] Preset template for generation
- [x] Custom template file
- [x] Conventional Commits Parser
- [x] Generate multiple versions of change logs
- [x] [Github Action](https://github.com/whatchanged-community/setup-whatchanged)

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
                --project=https://github.com/whatchanged-community/whatchanged.git
                Defaults to "--project=$PWD".
  --branch      Specify clone branch. only works for remote git project.
                Defaults to "--branch=master"
  --output      Write output to file. default write to stdout.
  --fmt         Specify the changelog format. Available options:
                --fmt=md
                --fmt=json
                Defaults to "--fmt=md".
  --preset      Cli built-in markdown template. Available options:
                --preset=default
                --preset=full
                --preset=simple
                Only available when --fmt=md and --tpl is nil.
                Defaults to "--preset=default".
  --tpl         Specify the template file for generating. Only available when
                --fmt=md.
  --skip-format Skip the formatting process, which is very useful for keeping the
                original format.

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
  $ whatchanged --project=https://github.com/whatchanged-community/whatchanged.git v0.1.0

SOURCE CODE:
  https://github.com/whatchanged-community/whatchanged
```

### Install

1. Shell (Mac/Linux)

```bash
# install latest version
curl -fsSL https://raw.githubusercontent.com/whatchanged-community/whatchanged/master/install.sh | bash
# or install specified version
curl -fsSL https://raw.githubusercontent.com/whatchanged-community/whatchanged/master/install.sh | bash -s v0.4.0
# or install from gobinaries.com
curl -sf https://gobinaries.com/whatchanged-community/whatchanged@v0.4.0 | sh
```

2. PowerShell (Windows):

```bash
# install latest version
iwr https://github.com/whatchanged-community/whatchanged/raw/master/install.ps1 -useb | iex
# or install specified version
$v="v0.4.0"; iwr https://github.com/whatchanged-community/whatchanged/raw/master/install.ps1 -useb | iex
```

3. [Github release page](https://github.com/whatchanged-community/whatchanged/releases) (All platforms)

download the executable file and put the executable file to `$PATH` then try the following command:

```bash
$ whatchanged --help
```

4. Build and install from source using [Golang](https://golang.org) (All platforms)

```bash
go install github.com/whatchanged-community/whatchanged@v0.4.0/cmd/whatchanged
```

### FAQ

1. [How it works?](HOW_IT_WORKS.md)
2. [How to custom generation template?](CUSTOM_TEMPLATE.md)

### License

The [Anti-996 License](LICENSE)
