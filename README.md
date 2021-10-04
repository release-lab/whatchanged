[English](README.md) | 中文简体

[![Build Status](https://github.com/whatchanged-community/whatchanged/workflows/ci/badge.svg)](https://github.com/whatchanged-community/whatchanged/actions)
[![Build Status](https://github.com/whatchanged-community/whatchanged/workflows/playground/badge.svg)](https://github.com/whatchanged-community/whatchanged/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/whatchanged-community/whatchanged)](https://goreportcard.com/report/github.com/whatchanged-community/whatchanged)
![Latest Version](https://img.shields.io/github/v/release/whatchanged-community/whatchanged.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/whatchanged-community/whatchanged.svg)

## whatchanged

一个优雅的变更日志生成器

致力于**优雅**/**简单**/**高效**/**可扩展**

[在线感受魔法](https://whatchanged-community.github.io)

特性:

- [x] 跨平台支持
- [x] 生成 本地/远程 的 git 仓库
- [x] 内置预设的生成模板
- [x] 自定义模版文件
- [x] 常规的 commit 解析器
- [x] 支持多个版本的变更日志生成
- [x] [Github Action](https://github.com/whatchanged-community/setup-whatchanged)

### 使用

```bash
$ whatchanged --help
whatchanged - 一个为 Git 项目生成变更日志的命令行工具

USAGE:
  whatchanged [OPTIONS] [ARGUMENTS...]

ARGUMENTS:
  [version...]  可选的版本或者版本范围，它有一下几种值
                1.null.
                  如果你没有指定版本，则会自动从 "HEAD~<latest tag>" /
                  "HEAD~<earliest commit>" / "<latest tag>-<prev tag>"
                  中选择一个进行生成
                2.单版本，例如 "v1.2.0"
                  生成单个指定版本的变更日志
                3.多版本，例如 "v2.0.0 v1.0.0"
                4.版本范围，例如 v1.3.0~v1.2.0
                  生成这两个版本以及中间版本的变更日志。
                  格式: <开始>~<结束>

OPTIONS:
  --help        打印帮助信息。
  --version     打印版本信息。
  --project     指定要生成的 Git 项目，它可以是一个相对/绝对路径或者远程地址。例如：
                --project=/path/to/project/which/contains/.git/folder
                --project=https://github.com/whatchanged-community/whatchanged.git
                默认值: "--project=$PWD".
  --branch      指定要克隆的分支，仅在生成远程仓库时有效。
                默认值: "--branch=master"
  --output      指定输入文件，默认输出到 stdout（标准输出流）。
  --fmt         指定生成的格式，可选项：
                --fmt=md
                --fmt=json
                默认值: "--fmt=md".
  --preset      指定内置的 markdown 预设模版. 可选项:
                --preset=default
                --preset=full
                --preset=simple
                仅在设置 --fmt=md 并且 --tpl 为空时有效。
                默认值: "--preset=default".
  --tpl         指定渲染的 markdown 模版。仅在 --fmt=md 时有效。
  --skip-format 跳过格式化过程，一般情况下可以忽略

EXAMPLES:
  # 生成从 HEAD 到 <latest tag> 的变更日志。
  # 如果 HEAD 不是最新的 tag，则变更日志应该是一个未发布版本。
  # 否则生成最新版
  $ whatchanged

  # 生成指定版本的变更日志
  $ whatchanged v1.2.0

  # 生成指定两个版本的变更日志
  # 只会生成这两个版本的日志，中间版本会被忽略
  $ whatchanged v2.0.0 v1.0.0

  # 生成从 HEAD 到最新一个 tag 和第 N 个 tag 的变更日志
  $ whatchanged HEAD~@0 @1 @2

  # 生成两个版本范围内的变更日志，中间版本也会生成，例如 v1.3.0, v1.2.1, v1.2.0
  $ whatchanged v1.3.0~v1.2.0

  # 生成从 HEAD 到第 N 个 tag 的变更日志
  $ whatchanged ~@0

  # 生成第 0 个 tag 到第 2 个tag 的变更日志
  $ whatchanged @0~@2

  # 生成 HEAD 到指定版本的变更日志
  $ whatchanged HEAD~v1.3.0

  # 生成所有的变更日志
  $ whatchanged HEAD~

  # 生成两次提交之间的变更日志
  $ whatchanged 770ed02~585445d

  # 生成本地项目的变更日志
  $ whatchanged --project=/path/to/project v1.0.0

  # 生成远程项目的变更日志
  $ whatchanged --project=https://github.com/whatchanged-community/whatchanged.git v0.1.0

SOURCE CODE:
  https://github.com/whatchanged-community/whatchanged
```

### 安装

1. Shell (Mac/Linux)

```bash
# 安装最新版本
curl -fsSL https://raw.githubusercontent.com/whatchanged-community/whatchanged/master/install.sh | bash
# 或者安装指定版本
curl -fsSL https://raw.githubusercontent.com/whatchanged-community/whatchanged/master/install.sh | bash -s v0.4.1
# 或者通过 gobinaries.com 安装
curl -sf https://gobinaries.com/whatchanged-community/whatchanged@v0.4.1 | sh
```

2. PowerShell (Windows):

```bash
# 安装最新版本
iwr https://github.com/whatchanged-community/whatchanged/raw/master/install.ps1 -useb | iex
# 或者安装指定版本
$v="v0.4.1"; iwr https://github.com/whatchanged-community/whatchanged/raw/master/install.ps1 -useb | iex
```

3. [Github release page](https://github.com/whatchanged-community/whatchanged/releases) (全平台支持))

下载可执行文件，并且把它加入到`$PATH` 环境变量中，然后尝试以下命令：

```bash
$ whatchanged --help
```

4. 使用 [Golang](https://golang.org) 从源码中构建并安装 (全平台支持)

```bash
go install github.com/whatchanged-community/whatchanged/cmd/whatchanged@v0.4.1
```

### 常见问题

1. [它是如何工作的?](HOW_IT_WORKS.md)
2. [如何自定义生成的模版?](CUSTOM_TEMPLATE.md)

### 开源许可

The [Anti-996 License](LICENSE_zh-CN)
