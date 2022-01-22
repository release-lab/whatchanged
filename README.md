[English](README-en-US.md) | 中文简体

[![Build Status](https://github.com/release-lab/whatchanged/workflows/ci/badge.svg)](https://github.com/release-lab/whatchanged/actions)
[![Build Status](https://github.com/release-lab/whatchanged/workflows/playground/badge.svg)](https://github.com/release-lab/whatchanged/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/release-lab/whatchanged)](https://goreportcard.com/report/github.com/release-lab/whatchanged)
![Latest Version](https://img.shields.io/github/v/release/release-lab/whatchanged.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/release-lab/whatchanged.svg)

## whatchanged

一个优雅的变更日志生成器，只需遵循 [Conventional Commits 规范](https://www.conventionalcommits.org/en/v1.0.0/)，即可生成一个漂亮的，工整的变更日志。

致力于**优雅**/**简单**/**高效**/**可扩展**

[在线感受魔法](https://release-lab.github.io)

特性:

- [x] 跨平台支持
- [x] 生成 本地/远程 的 git 仓库
- [x] 内置预设的生成模板
- [x] 自定义模版文件
- [x] [Conventional Commits 规范解析器](https://github.com/release-lab/conventional-commit-parser)
- [x] 支持多个版本的变更日志生成
- [x] [Github Action](https://github.com/release-lab/setup-whatchanged)
- [x] [Visual Studio Code extension](https://github.com/release-lab/vscode-whatchanged)

### 使用

```bash
$ whatchanged --help
```

### 安装

1. Shell (Mac/Linux)

   ```bash
   curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=release-lab/whatchanged
   ```

2. PowerShell (Windows):

   ```powershell
   $r="release-lab/whatchanged";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
   ```

3. 从 [Github Release Page](https://github.com/release-lab/whatchanged/releases) 下载 (全平台支持)

   下载可执行文件，并且把它加入到`$PATH` 环境变量中

4. 使用 [Golang](https://golang.org) 从源码中构建并安装 (全平台支持)

   ```bash
   go install github.com/release-lab/whatchanged/cmd/whatchanged
   ```

### 常见问题

1. [它是如何工作的?](HOW_IT_WORKS.md)
2. [如何自定义生成的模版?](CUSTOM_TEMPLATE.md)

### 开源许可

The [Anti-996 License](LICENSE_zh-CN)
