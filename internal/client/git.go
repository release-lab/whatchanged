package client

import (
	"io"
	"path"
	"sort"

	"github.com/blang/semver/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
)

type Tag struct {
	Name   string         // tag name
	Hash   string         // tag hash
	Commit *object.Commit // commit
}

func NewGitClient(dir string) (*GitClient, error) {
	r, err := git.PlainOpen(path.Join(dir, ".git"))

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &GitClient{repository: r}, nil
}

type GitClient struct {
	repository *git.Repository
	tags       []*Tag
}

func (g *GitClient) TagN(offset int) (*Tag, error) {
	tags, err := g.Tags()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(tags) == 0 {
		return nil, nil
	}

	if len(tags) < offset {
		return nil, nil
	}

	tag := tags[offset]

	return tag, nil
}

func (g *GitClient) NextTag(target *Tag) (*Tag, error) {
	tags, err := g.Tags()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for index, tag := range tags {
		if tag.Hash == target.Hash {
			if len(tags) > index+1 {
				next := tags[index+1]
				return next, nil
			}
		}
	}

	return nil, nil
}

func (g *GitClient) TagName(name string) (*Tag, error) {
	tags, err := g.Tags()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(tags) == 0 {
		return nil, nil
	}

	for _, tag := range tags {
		if tag.Name == name {
			return tag, nil
		}
	}

	return nil, nil
}

func (g *GitClient) Tags() ([]*Tag, error) {
	if g.tags != nil {
		return g.tags, nil
	}
	tags := make([]*Tag, 0)

	tt, _ := g.repository.References()

	for {
		if ref, err := tt.Next(); err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.WithStack(err)
		} else if ref == nil {
			break
		} else {
			name := ref.Name()
			if name.IsTag() {
				commitHash, err := g.repository.ResolveRevision(plumbing.Revision(ref.Hash().String()))

				if err != nil {
					return nil, errors.WithStack(err)
				}

				commit, err := g.repository.CommitObject(*commitHash)

				if err != nil {
					return nil, errors.WithStack(err)
				}

				_, invalid := semver.New(name.Short())

				if invalid == nil {
					tags = append(tags, &Tag{
						Name:   name.Short(),
						Hash:   ref.Hash().String(),
						Commit: commit,
					})
				}
			}
		}
	}

	sort.Slice(tags, func(i int, j int) bool {
		prev := tags[i]
		next := tags[j]

		prevVersion, _ := semver.New(prev.Name)
		nextVersion, _ := semver.New(next.Name)

		// return prev.Commit.Committer.When.UnixNano() > next.Commit.Committer.When.UnixNano()
		return prevVersion.GT(*nextVersion)
	})

	g.tags = tags

	return tags, nil
}

func (g *GitClient) Head() (*plumbing.Reference, error) {
	return g.repository.Head()
}

func (g *GitClient) Logs(from string, to string) ([]*object.Commit, error) {
	options := git.LogOptions{From: plumbing.NewHash(from)}

	if to != "" {
		toCommit, err := g.repository.CommitObject(plumbing.NewHash(to))

		if err != nil {
			return nil, errors.WithStack(err)
		}

		options.Since = &toCommit.Committer.When
	}

	cIter, err := g.repository.Log(&options)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	commits := make([]*object.Commit, 0)

	for {
		if commit, err := cIter.Next(); err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.WithStack(err)
		} else if commit == nil {
			break
		} else if commit.Hash.String() == to {
			break
		} else {
			commits = append(commits, commit)
		}
	}

	return commits, errors.WithStack(err)
}

func (g *GitClient) LogsAuto() ([]*object.Commit, error) {
	var (
		toHash string
	)

	head, err := g.repository.Head()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get latest tag
	tag, err := g.TagN(0)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	// if tag not exist
	if tag != nil {
		toHash = tag.Commit.Hash.String()

		// if head is latest tag's commit
		if head.Hash().String() == toHash {
			// get next tag
			nextTag, err := g.TagN(1)

			if err != nil {
				return nil, errors.WithStack(err)
			}

			if nextTag != nil {
				toHash = nextTag.Commit.Hash.String()
			} else {
				toHash = ""
			}
		}
	}

	cIter, err := g.repository.Log(&git.LogOptions{From: head.Hash()})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	commits := make([]*object.Commit, 0)

	for {
		if commit, er := cIter.Next(); er == io.EOF {
			break
		} else if er != nil {
			return nil, er
		} else if commit == nil {
			break
		} else if commit.Hash.String() == toHash {
			break
		} else {
			commits = append(commits, commit)
		}
	}

	return commits, errors.WithStack(err)
}
