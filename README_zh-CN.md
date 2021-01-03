[English](README.md) | 中文简体

[![Build Status](https://github.com/axetroy/whatchanged/workflows/ci/badge.svg)](https://github.com/axetroy/whatchanged/actions)
[![Build Status](https://github.com/axetroy/whatchanged/workflows/playground/badge.svg)](https://github.com/axetroy/whatchanged/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/whatchanged)](https://goreportcard.com/report/github.com/axetroy/whatchanged)
![Latest Version](https://img.shields.io/github/v/release/axetroy/whatchanged.svg)
![License](https://img.shields.io/github/license/axetroy/whatchanged.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/whatchanged.svg)

## whatchanged

一个优雅的变更日志生成器

致力于**优雅**/**简单**/**高效**/**可扩展**

[在线感受魔法](https://axetroy.github.io/whatchanged/)

Feature:

- [x] 跨平台支持
- [x] 生成 本地/远程 的git仓库
- [x] 内置预设的生成模板
- [x] 自定义模版文件
- [x] 常规的 commit 解析器
- [x] 支持多个版本的变更日志生成
- [x] [Github Action](https://github.com/axetroy/setup-whatchanged)

### 使用

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
  $ whatchanged --project=https://github.com/axetroy/whatchanged.git v0.1.0

SOURCE CODE:
  https://github.com/axetroy/whatchanged
```

### 安装

#### 通过 npm 安装

```bash
# install in global
npm install -g @axetroy/whatchanged
# run the command once
npx @axetroy/whatchanged
```

#### 通过 shell 安装

如果你使用的是 Linux/macOS，你可以通过以下命令进行安装

```shell
# 安装最新版本
curl -fsSL https://raw.githubusercontent.com/axetroy/whatchanged/master/install.sh | bash
# 或者安装指定版本
curl -fsSL https://raw.githubusercontent.com/axetroy/whatchanged/master/install.sh | bash -s v0.3.3
# 或者通过 gobinaries.com 安装
curl -sf https://gobinaries.com/axetroy/whatchanged@v0.3.3 | sh
```

#### 通过 Github release 页面下载安装

从[release page](https://github.com/axetroy/whatchanged/releases)页面下载对应平台的可执行文件，并且把它加入到 `$PATH` 环境变量中，并尝试以下命令

```bash
$ whatchanged --help
```

### 从源码中构建

确保你已安装 `Golang@v1.16.x` 和 [goreleaser](https://github.com/goreleaser/goreleaser)

```shell
$ git clone https://github.com/axetroy/whatchanged.git $GOPATH/src/github.com/axetroy/whatchanged
$ cd $GOPATH/src/github.com/axetroy/whatchanged
$ make build
```

### 测试

```bash
$ make test
```

### 常见问题

1. [它是如何工作的?](HOW_IT_WORKS.md)
2. [如何自定义生成的模版?](CUSTOM_TEMPLATE.md)

### 开源许可

The [MIT License](LICENSE)
