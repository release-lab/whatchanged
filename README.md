English | [中文简体](README-zh-CN.md)

[![Build Status](https://github.com/release-lab/whatchanged/workflows/ci/badge.svg)](https://github.com/release-lab/whatchanged/actions)
[![Build Status](https://github.com/release-lab/whatchanged/workflows/playground/badge.svg)](https://github.com/release-lab/whatchanged/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/release-lab/whatchanged)](https://goreportcard.com/report/github.com/release-lab/whatchanged)
![Latest Version](https://img.shields.io/github/v/release/release-lab/whatchanged.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/release-lab/whatchanged.svg)

## whatchanged

An elegant changelog generator. follow the [Conventional Commits Specification](https://www.conventionalcommits.org/en/v1.0.0/) to generate a beautiful and neat change log.

Focus on **Elegant**/**Simple**/**Efficient**/**Scalable**

[Feel the magic online](https://release-lab.github.io)

Feature:

- [x] Cross-platform support
- [x] Generation for local/Remote git repositories
- [x] Preset template for generation
- [x] Custom template file
- [x] [Conventional Commits Specification Parser](https://github.com/release-lab/conventional-commit-parser)
- [x] Generate multiple versions of change logs
- [x] [Github Action](https://github.com/release-lab/setup-whatchanged)
- [x] [Visual Studio Code extension](https://github.com/release-lab/vscode-whatchanged)

### Usage

```bash
$ whatchanged --help
```

### Install

1. Shell (Mac/Linux)

   ```bash
   curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=release-lab/whatchanged
   ```

2. PowerShell (Windows):

   ```powershell
   $r="release-lab/whatchanged";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
   ```

3. [Github release page](https://github.com/release-lab/whatchanged/releases) (All platforms)

   download the executable file and put the executable file to `$PATH`

4. Build and install from source using [Golang](https://golang.org) (All platforms)

   ```bash
   go install github.com/release-lab/whatchanged/cmd/whatchanged
   ```

### Use as library

```bash
# install package
go get -v -u github.com/release-lab/whatchanged
```

```go
package main

import (
   "context"
   "bytes"

   "github.com/release-lab/whatchanged"
)

func main() {
   output := bytes.NewBuffer([]byte{})

   option := whatchanged.Options{
      Version: []string{"HEAD~"}
   }

   err := whatchanged.Generate(context.Background(), "/path/to/git/project", output, &option)

   if err!=nil{
      panic(err)
   }

   println(output)
}
```

### FAQ

1. [How it works?](HOW_IT_WORKS.md)
2. [How to custom generation template?](CUSTOM_TEMPLATE.md)

### License

The [Anti-996 License](LICENSE)
