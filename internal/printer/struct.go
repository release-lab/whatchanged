package printer

import (
	"github.com/axetroy/changelog/internal/commit/parser"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Model struct {
	Commit *Commit
}

type Commit struct {
	Hash      string
	Short     string
	Message   string
	Author    *object.Signature
	Committer *object.Signature
	Field     *parser.Message
}

// https://github.com/angular/angular/blob/master/CONTRIBUTING.md#commit-message-header
type TemplateContext struct {
	Version         string
	Build           []*Commit
	CI              []*Commit
	Docs            []*Commit
	Feat            []*Commit
	Fix             []*Commit
	Perf            []*Commit
	Refactor        []*Commit
	Test            []*Commit
	Revert          []*Commit
	BreakingChanges []*Commit
	Commits         []*Commit
}
