package extractor

import (
	"io"

	parser "github.com/axetroy/whatchanged/1_parser"
	"github.com/axetroy/whatchanged/internal/client"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
)

type ExtractSplice struct {
	Name   string           // The name of splice
	Commit []*object.Commit // The commits including in this splice
	Tag    *client.Tag      // If this splice is a tag splice
}

func getTagOfCommit(tags []*client.Tag, commit *object.Commit) (result []*client.Tag) {
	for _, tag := range tags {
		if tag.Commit.Hash.String() == commit.Hash.String() {
			result = append(result, tag)
		}
	}

	return
}

func Extract(g *client.GitClient, scopes []*parser.Scope) ([]*ExtractSplice, error) {
	splices := make([]*ExtractSplice, 0)

	for _, scope := range scopes {
		options := git.LogOptions{From: plumbing.NewHash(scope.Start.Commit.Hash.String())}

		cIter, err := g.Repository.Log(&options)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		commits := make([]*object.Commit, 0)

		for {
			if c, err := cIter.Next(); err == io.EOF {
				break
			} else if err != nil {
				return nil, errors.WithStack(err)
			} else if c == nil {
				break
			} else if c.Hash.String() == scope.End.Commit.Hash.String() {
				commits = append(commits, c)
				break
			} else {
				commits = append(commits, c)
			}
		}

		if err != nil {
			return nil, errors.WithStack(err)
		}

		// if there is no tags in this scope
		// if scope.Tags == nil || len(scope.Tags) == 0 {
		// 	splices = append(splices, &ExtractSplice{
		// 		Name:   "Unreleased",
		// 		Commit: commits,
		// 		Tag:    nil,
		// 	})

		// 	return splices, nil
		// }

		index := 0

		tagIndex := 0

		for {
			if index == len(commits) {
				break
			}

			currentCommit := commits[index]

			item := &ExtractSplice{
				Name:   "Unreleased",
				Commit: make([]*object.Commit, 0),
			}

			if tags := getTagOfCommit(scope.Tags, currentCommit); len(tags) != 0 {
				if tagIndex+1 == len(tags) {
					item.Tag = tags[tagIndex]
					item.Name = tags[tagIndex].Name
					item.Commit = append(item.Commit, currentCommit)
					tagIndex = 0 // reset tag index to zero
				} else {
					item.Tag = tags[tagIndex]
					item.Name = tags[tagIndex].Name
					splices = append(splices, item)
					tagIndex++
					continue
				}
			} else {
				item.Commit = append(item.Commit, currentCommit)
			}

			index++

		loop:
			for {
				if index == len(commits) {
					break loop
				}

				nextCommit := commits[index]

				if t := getTagOfCommit(scope.Tags, nextCommit); t != nil {
					break loop
				}

				item.Commit = append(item.Commit, nextCommit)

				index++
			}

			splices = append(splices, item)

		}
	}

	return splices, nil
}
