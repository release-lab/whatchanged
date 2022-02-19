# How to custom template for generation?

Use the `--tpl=./template.md` flag to specify the generated template.

`whatchanged` uses the [template](https://golang.org/pkg/text/template/) in the Golang standard library.

Therefore, as long as your template syntax is correct, it can be generated normally.

Before that, let’s explain the render context in the template.

There are 3 parts in Context

1. Variable

   1. **Version**: the version string
   2. **Date**: the date string. the format is `YYYY-MM-dd`
   3. **Feat**: the `commit` list of `new features` which start with `feat(xxx): xxx`
   4. **Fix**: the `commit` list of `bus fixed` which start with `fix(xxx): xxx`
   5. **Refactor**: the `commit` list of `code refactor` which start with `refactor(xxx): xxx`
   6. **Test**: the `commit` list of `test` which start with `test(xxx): xxx`
   7. **Perf**: the `commit` list of `performance` which start with `pref(xxx): xxx`
   8. **Build**: the `commit` list of `build` which start with `build(xxx): xxx`
   9. **Ci**: the `commit` list of `ci` which start with `ci(xxx): xxx`
   10. **Chore**: the `commit` list of `chore` which start with `chore(xxx): xxx`
   11. **Docs**: the `commit` list of `documentation` which start with `docs(xxx): xxx`
   12. **Style**: the `commit` list of `styles` which start with `style(xxx): xxx`
   13. **Revert**: the `commit` list of `revert` which start with `revert: xxx`
   14. **BreakingChanges**: the `commit` list of `revert` which includes `BREAKING CHANGES: xxx` in commit message body
   15. **Commits**: all `commit` of generation

2. Methods

   1. **unescape(text string)**: unescape the string `text`.
   2. All functions from the library [sprig](https://github.com/Masterminds/sprig)

source code:

```go
type Commit struct {
	Hash                string
	HashURL             string
	Short               string
	Message             string
	Author              *object.Signature
	Committer           *object.Signature
	Field               Message
	RevertCommitHash    *string   // Revert hash
	RevertCommitHashURL *string   // Revert hash URL
	Closes              *[]string // Closes issues
}

// https://github.com/angular/angular/blob/master/CONTRIBUTING.md#commit-message-header
type TemplateContext struct {
	Version         string
	Date            string
	Build           []*Commit
	Ci              []*Commit
	Chore           []*Commit
	Docs            []*Commit
	Feat            []*Commit
	Fix             []*Commit
	Perf            []*Commit
	Refactor        []*Commit
	Test            []*Commit
	Style           []*Commit
	Revert          []*Commit
	BreakingChanges []*Commit
	Commits         []*Commit
	Contributors    []string
}

type Header struct {
	Type      string
	Scope     string
	Subject   string
	Important bool
}


type BreakingChange struct {
	Title   string
	Content string
}

type Footer struct {
	BreakingChange *BreakingChange
	Closes         string
}


type Message struct {
	raw    string
	Title  string
	Header *conventionalcommitparser.Header
	Body   string
	Footer *Footer
}


// Signature is used to identify who and when created a commit or tag.
type Signature struct {
	// Name represents a person name. It is an arbitrary string.
	Name string
	// Email is an email, but it cannot be assumed to be well-formed.
	Email string
	// When is the timestamp of the signature.
	When time.Time
}
```

## example

[Default Template](template/default.tpl)

[Full Template](template/full.tpl)

## Source Code

[transformer.go](3_transformer.go)
