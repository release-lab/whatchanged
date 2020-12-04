package client

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
)

var (
	versionRegexp         = regexp.MustCompile(`^v`)
	winAbsolutePathRegegp = regexp.MustCompile(`^(?i)[a-z]:\.+`)
)

type Tag struct {
	Name   string         // tag name
	Hash   string         // tag hash
	Commit *object.Commit // commit
}

func NewGitClient(dir string) (*GitClient, error) {
	var (
		isProjectDir bool = false
	)

	cwd, err := os.Getwd()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !path.IsAbs(dir) && !winAbsolutePathRegegp.MatchString(dir) {
		dir = path.Join(cwd, dir)
	}

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, file := range files {
		if file.Name() == ".git" {
			if file.IsDir() {
				isProjectDir = true
			} else {
				// git submodule
				content, err := ioutil.ReadFile(path.Join(dir, file.Name()))

				if err != nil {
					return nil, errors.WithStack(err)
				}

				matcher := regexp.MustCompile(`gitdir: (.*)`).FindStringSubmatch(string(content))

				subModuleGitDir := path.Join(dir, matcher[1])

				dir = subModuleGitDir
			}
			break
		}
	}

	if isProjectDir {
		dir = path.Join(dir, ".git")
	}

	r, err := git.PlainOpen(dir)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &GitClient{Repository: r}, nil
}

type GitClient struct {
	Repository *git.Repository
	tags       []*Tag
}

// Get the Nth tag
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

// Get next tag from the specified tag
func (g *GitClient) PrevTag(target *Tag) (*Tag, error) {
	tags, err := g.Tags()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for index, tag := range tags {
		if tag.Hash == target.Hash {
			i := index - 1
			if i < 0 {
				return nil, nil
			}
			return tags[i], nil
		}
	}

	return nil, nil
}

// Get next tag from the specified tag
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

// Get tag information from tag name
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

// Get all tags
func (g *GitClient) Tags() ([]*Tag, error) {
	if g.tags != nil {
		return g.tags, nil
	}
	tags := make([]*Tag, 0)

	tt, _ := g.Repository.References()

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
				commitHash, err := g.Repository.ResolveRevision(plumbing.Revision(ref.Hash().String()))

				if err != nil {
					return nil, errors.WithStack(err)
				}

				commit, err := g.Repository.CommitObject(*commitHash)

				if err != nil {
					return nil, errors.WithStack(err)
				}

				version := versionRegexp.ReplaceAllString(name.Short(), "")

				_, invalid := semver.New(version)

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

		prevVersion, _ := semver.New(versionRegexp.ReplaceAllString(prev.Name, ""))
		nextVersion, _ := semver.New(versionRegexp.ReplaceAllString(next.Name, ""))

		// return prev.Commit.Committer.When.UnixNano() > next.Commit.Committer.When.UnixNano()
		return prevVersion.GT(*nextVersion)
	})

	g.tags = tags

	return tags, nil
}

// Get HEAD of tree
func (g *GitClient) Head() (*plumbing.Reference, error) {
	return g.Repository.Head()
}

// Get commit of HASH
func (g *GitClient) Commit(longHash string) (*object.Commit, error) {
	return g.Repository.CommitObject(plumbing.NewHash(longHash))
}

// Get commit of SHORT HASH
func (g *GitClient) CommitByShort(shortHash string) (*object.Commit, error) {
	cIter, err := g.Repository.Log(&git.LogOptions{})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for {
		if commit, err := cIter.Next(); err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.WithStack(err)
		} else if commit != nil {
			if strings.HasPrefix(commit.Hash.String(), shortHash) {
				return commit, nil
			}
		}
	}

	return nil, nil
}

// Get first commit
func (g *GitClient) GetFirstCommit() (*object.Commit, error) {
	options := git.LogOptions{}

	cIter, err := g.Repository.Log(&options)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var c *object.Commit

	for {
		if commit, err := cIter.Next(); err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.WithStack(err)
		} else if commit != nil {
			if c == nil {
				c = commit
			} else {
				if c.Committer.When.UnixNano() > commit.Committer.When.UnixNano() {
					c = commit
				}
			}
		}
	}

	return c, nil
}

// Get commit list from range
// to: is optional
func (g *GitClient) Logs(from string, to string) ([]*object.Commit, error) {
	options := git.LogOptions{From: plumbing.NewHash(from)}

	if to != "" {
		toCommit, err := g.Repository.CommitObject(plumbing.NewHash(to))

		if err != nil {
			return nil, errors.WithStack(err)
		}

		options.Since = &toCommit.Committer.When
	}

	cIter, err := g.Repository.Log(&options)

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

// Get next commit behind a commit
func (g *GitClient) PrevCommit(hash string) (*object.Commit, error) {
	options := git.LogOptions{From: plumbing.NewHash(hash)}

	cIter, err := g.Repository.Log(&options)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if commit, err := cIter.Next(); err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	} else {
		return commit, nil
	}
}

// Get next commit behind a commit
func (g *GitClient) NextCommit(hash string) (*object.Commit, error) {
	options := git.LogOptions{From: plumbing.NewHash(hash)}

	cIter, err := g.Repository.Log(&options)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var next *object.Commit

	if commit, err := cIter.Next(); err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	} else {
		if next.Hash.String() == hash {
			return next, nil
		}
		next = commit
	}

	return next, nil
}

// Get tags with range
// eg. Get all tags between v2.0.0 and v1.0.0 and including thenself
// start: v2.0.0
// end: v1.0.0
func (g *GitClient) GetTagRangesByTagName(start string, end string) ([]*Tag, error) {
	tags, err := g.Tags()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	startIndex := 0
	endIndex := 0

	for index, tag := range tags {
		if tag.Name == start {
			startIndex = index
		} else if tag.Name == end {
			endIndex = index
		}
	}

	result := tags[startIndex : endIndex+1]

	return result, nil
}

// Get tags with commit range
func (g *GitClient) GetTagRangesByCommit(start *object.Commit, end *object.Commit) ([]*Tag, error) {
	tags, err := g.Tags()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make([]*Tag, 0)

	for _, tag := range tags {
		tagTime := tag.Commit.Committer.When
		if (tagTime.Before(start.Author.When) || tagTime.Equal(start.Author.When)) && (tagTime.After(end.Author.When) || tagTime.Equal(end.Author.When)) {
			result = append(result, tag)
		}
	}

	return result, nil
}

func (g *GitClient) GetTagByCommit(commit *object.Commit) (*Tag, error) {
	tags, err := g.Tags()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, tag := range tags {
		if tag.Commit.Hash.String() == commit.Hash.String() {
			return tag, nil
		}
	}

	return nil, nil
}

func (g *GitClient) GetLastCommitOfTag(tag *Tag) (*object.Commit, error) {
	nextTag, err := g.NextTag(tag)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	options := git.LogOptions{From: tag.Commit.Hash}

	cIter, err := g.Repository.Log(&options)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if nextTag == nil {
		options.Since = &tag.Commit.Committer.When
		var lastCommit *object.Commit
		for {
			if commit, err := cIter.Next(); err == io.EOF {
				break
			} else if err != nil {
				return nil, errors.WithStack(err)
			} else if commit == nil {
				break
			} else {
				lastCommit = commit
			}
		}

		return lastCommit, nil
	} else {

		var lastCommit *object.Commit

		for {
			if commit, err := cIter.Next(); err == io.EOF {
				break
			} else if err != nil {
				return nil, errors.WithStack(err)
			} else if commit == nil {
				break
			} else {
				if commit.Hash.String() == nextTag.Commit.Hash.String() {
					return lastCommit, nil
				}
				lastCommit = commit
			}
		}

		return lastCommit, nil
	}
}

func (g *GitClient) GetPrevCommitOfTag(tag *Tag) (*object.Commit, error) {
	prevTag, err := g.PrevTag(tag)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	options := git.LogOptions{}

	if prevTag != nil {
		options.From = prevTag.Commit.Hash
	}

	cIter, err := g.Repository.Log(&options)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var prevCommit *object.Commit

	for {
		if commit, err := cIter.Next(); err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.WithStack(err)
		} else if commit == nil {
			break
		} else {
			if commit.Hash.String() == tag.Commit.Hash.String() {
				return prevCommit, nil
			}
			prevCommit = commit
		}
	}

	return prevCommit, nil
}

func (g *GitClient) GetRemote() (*config.RemoteConfig, error) {
	repoConfig, err := g.Repository.Config()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if remote, ok := repoConfig.Remotes["origin"]; ok {
		return remote, nil
	} else {
		for _, remote := range repoConfig.Remotes {
			for _, url := range remote.URLs {
				// prefer github/gitlab remote
				if strings.HasPrefix(url, "https://github.com") {
					return remote, nil
				} else if strings.HasPrefix(url, "https://gitlab.com") {
					return remote, nil
				}
			}
		}

		for _, remote := range repoConfig.Remotes {
			return remote, nil
		}
	}
	return nil, nil
}
