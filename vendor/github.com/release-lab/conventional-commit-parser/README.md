[![Build Status](https://github.com/release-lab/conventional-commit-parser/workflows/ci/badge.svg)](https://github.com/release-lab/conventional-commit-parser/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/release-lab/conventional-commit-parser)](https://goreportcard.com/report/github.com/release-lab/conventional-commit-parser)
![Latest Version](https://img.shields.io/github/v/release/release-lab/conventional-commit-parser.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/release-lab/conventional-commit-parser.svg)

### Conventional Commit Parser

This is a parser for [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

```bash
go get -u github.com/release-lab/conventional-commit-parser
```

```go
package main

import (
  "fmt"
  "github.com/release-lab/conventional-commit-parser"
)

func main() {
  result := conventionalcommitparser.Parse("feat: this is a commit message")

  fmt.Printf("%+v\n", result)
}
```

### License

The [Anti-996 License](LICENSE)
