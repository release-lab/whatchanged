package transform

import (
	extractor "github.com/axetroy/changelog/2_extractor"
	"github.com/axetroy/changelog/internal/client"
	"github.com/axetroy/changelog/internal/commit/parser"
	"github.com/go-git/go-git/v5/plumbing/object"
)

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

func Transform(g *client.GitClient, splices []*extractor.ExtractSplice) ([]*TemplateContext, error) {
	context := make([]*TemplateContext, 0)

	for _, splice := range splices {
		ctx := &TemplateContext{
			Version: splice.Name,
		}
		if len(splice.Commit) != 0 {
			ctx.Commits = make([]*Commit, 0)
		}
		for _, commit := range splice.Commit {
			hash := commit.Hash.String()
			c := &Commit{
				Hash:      hash,
				Short:     string(hash[0:7]),
				Message:   commit.Message,
				Author:    &commit.Author,
				Committer: &commit.Committer,
			}
			field := parser.Parser(commit.Message)

			c.Field = field

			if field.Footer != nil {
				if field.Footer.BreakingChange != nil {
					if ctx.BreakingChanges == nil {
						ctx.BreakingChanges = make([]*Commit, 0)
					}
					ctx.BreakingChanges = append(ctx.BreakingChanges, c)
				}
			}

			ctx.Commits = append(ctx.Commits, c)

			if field.Header != nil {
				switch field.Header.Type {
				case "build":
					if ctx.Build == nil {
						ctx.Build = make([]*Commit, 0)
					}
					ctx.Build = append(ctx.Build, c)
				case "ci":
					if ctx.CI == nil {
						ctx.CI = make([]*Commit, 0)
					}
					ctx.CI = append(ctx.CI, c)
				case "docs":
					if ctx.Docs == nil {
						ctx.Docs = make([]*Commit, 0)
					}
					ctx.Docs = append(ctx.Docs, c)
				case "feat":
					if ctx.Feat == nil {
						ctx.Feat = make([]*Commit, 0)
					}
					ctx.Feat = append(ctx.Feat, c)
				case "fix":
					if ctx.Fix == nil {
						ctx.Fix = make([]*Commit, 0)
					}
					ctx.Fix = append(ctx.Fix, c)
				case "perf":
					if ctx.Perf == nil {
						ctx.Perf = make([]*Commit, 0)
					}
					ctx.Perf = append(ctx.Perf, c)
				case "refactor":
					if ctx.Refactor == nil {
						ctx.Refactor = make([]*Commit, 0)
					}
					ctx.Refactor = append(ctx.Refactor, c)
				case "test":
					if ctx.Test == nil {
						ctx.Test = make([]*Commit, 0)
					}
					ctx.Test = append(ctx.Test, c)
				}

			}
		}

		context = append(context, ctx)
	}

	return context, nil
}
