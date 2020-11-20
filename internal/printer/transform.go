package printer

import (
	"github.com/axetroy/changelog/internal/commit/parser"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func transform(version string, commits []*object.Commit) *TemplateContext {
	ctx := TemplateContext{
		Version: version,
		Commits: make([]*Commit, 0),
	}
	for _, c := range commits {
		field := parser.Parser(c.Message)

		hash := c.Hash.String()

		commit := Commit{
			Hash:      hash,
			Short:     string(hash[0:7]),
			Message:   c.Message,
			Author:    &c.Author,
			Committer: &c.Committer,
			Field:     field,
		}

		if field.Footer != nil {
			if field.Footer.BreakingChange != nil {
				if ctx.BreakingChanges == nil {
					ctx.BreakingChanges = make([]*Commit, 0)
				}
				ctx.BreakingChanges = append(ctx.BreakingChanges, &commit)
			}
		}

		ctx.Commits = append(ctx.Commits, &commit)

		if field.Header != nil {
			switch field.Header.Type {
			case "build":
				if ctx.Build == nil {
					ctx.Build = make([]*Commit, 0)
				}
				ctx.Build = append(ctx.Build, &commit)
			case "ci":
				if ctx.CI == nil {
					ctx.CI = make([]*Commit, 0)
				}
				ctx.CI = append(ctx.CI, &commit)
			case "docs":
				if ctx.Docs == nil {
					ctx.Docs = make([]*Commit, 0)
				}
				ctx.Docs = append(ctx.Docs, &commit)
			case "feat":
				if ctx.Feat == nil {
					ctx.Feat = make([]*Commit, 0)
				}
				ctx.Feat = append(ctx.Feat, &commit)
			case "fix":
				if ctx.Fix == nil {
					ctx.Fix = make([]*Commit, 0)
				}
				ctx.Fix = append(ctx.Fix, &commit)
			case "perf":
				if ctx.Perf == nil {
					ctx.Perf = make([]*Commit, 0)
				}
				ctx.Perf = append(ctx.Perf, &commit)
			case "refactor":
				if ctx.Refactor == nil {
					ctx.Refactor = make([]*Commit, 0)
				}
				ctx.Refactor = append(ctx.Refactor, &commit)
			case "test":
				if ctx.Test == nil {
					ctx.Test = make([]*Commit, 0)
				}
				ctx.Test = append(ctx.Test, &commit)
			}

		}
	}

	return &ctx
}
