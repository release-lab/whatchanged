package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/axetroy/whatchanged/internal/client"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
)

type Range struct {
	*object.Commit
	*client.Tag
}

type Scope struct {
	Start Range         // The range start
	End   Range         // The range end
	Tags  []*client.Tag // Tags included in the range
}

var (
	regShortHash = regexp.MustCompile(`^[a-z\d]{7,8}$`)
	regLongHash  = regexp.MustCompile(`^[a-z\d]{40}$`)
	regTag       = regexp.MustCompile(`^@(\d+)$`)
)

func getCommitFromHEAD(git *client.GitClient) (*object.Commit, error) {
	ref, err := git.Head()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	commit, err := git.Commit(ref.Hash().String())

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return commit, err
}

func getCommit(git *client.GitClient, version string, isStart bool) (*object.Commit, error) {
	if version == "" {
		var (
			commit *object.Commit
			err    error
		)
		if isStart {
			// if version is a empty string, then it should be HEAD
			commit, err = getCommitFromHEAD(git)
		} else {
			// if version is a empty string, then it should be the first commit
			commit, err = git.GetFirstCommit()
		}

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return commit, nil
	} else if version == "HEAD" {
		commit, err := getCommitFromHEAD(git)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return commit, nil
	} else if regTag.MatchString(version) {
		matcher := regTag.FindStringSubmatch(version)

		tagIndex, err := strconv.ParseInt(matcher[1], 10, 10)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		tag, err := git.TagN(int(tagIndex))

		if err != nil {
			return nil, errors.WithStack(err)
		}

		if tag == nil {
			return nil, nil
		}

		return tag.Commit, nil
	} else if regShortHash.MatchString(version) {
		return git.CommitByShort(version)
	} else if regLongHash.MatchString(version) {
		return git.Commit(version)
	} else {
		// else if it's a tag
		tag, err := git.TagName(version)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		if tag == nil {
			return nil, errors.New(fmt.Sprintf("tag '%s' not found", version))
		}

		return tag.Commit, nil
	}
}

func resolveVersion(git *client.GitClient, versionRanges string) ([]string, error) {
	versions := strings.Split(versionRanges, "~")

	length := len(versions)

	switch length {
	case 0:
		versions[0] = "HEAD"
		latestTag, err := git.TagN(0)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		if latestTag != nil {
			commit, err := git.GetPrevCommitOfTag(latestTag)

			if err != nil {
				return nil, errors.WithStack(err)
			}

			// if there is no prev tag
			// tag is the HEAD
			if commit == nil {
				nextTag, err := git.NextTag(latestTag)

				if err != nil {
					return nil, errors.WithStack(err)
				}

				if nextTag != nil {
					if preCommitOfNextTag, err := git.GetPrevCommitOfTag(nextTag); err != nil {
						return nil, errors.WithStack(err)
					} else {
						// if the pre-commit of the tag is nil
						if preCommitOfNextTag == nil {
							versions = append(versions, nextTag.Commit.Hash.String())
						} else {
							versions = append(versions, preCommitOfNextTag.Hash.String())
						}
					}
				} else {
					versions = append(versions, "")
				}
			} else {
				versions = append(versions, commit.Hash.String())
			}

		} else {
			versions = append(versions, "")
		}
	case 1:
		var (
			tag *client.Tag
			err error
		)

		if regTag.MatchString(versionRanges) {
			matcher := regTag.FindStringSubmatch(versionRanges)

			tagIndex, err := strconv.ParseInt(matcher[1], 10, 10)

			if err != nil {
				return nil, errors.WithStack(err)
			}

			if tag, err = git.TagN(int(tagIndex)); err != nil {
				return nil, errors.WithStack(err)
			}
		} else {
			if tag, err = git.TagName(versionRanges); err != nil {
				return nil, errors.WithStack(err)
			}
		}

		if tag != nil {
			commit, err := git.GetLastCommitOfTag(tag)

			if err != nil {
				return nil, errors.WithStack(err)
			} else if commit != nil {
				versions = append(versions, commit.Hash.String())
			}

		} else {
			versions = append(versions, "")
		}
	}

	return versions, nil
}

func Parse(git *client.GitClient, rangesSlice []string) ([]*Scope, error) {
	var (
		scopeRange = make([]*Scope, 0)
	)

	if len(rangesSlice) == 0 {
		rangesSlice = []string{""}
	}

	for _, ranges := range rangesSlice {
		var (
			startCommit *object.Commit
			endCommit   *object.Commit
			scope       = &Scope{}
		)

		ranges = strings.TrimSpace(ranges)

		versions, err := resolveVersion(git, ranges)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		if startCommit, err = getCommit(git, versions[0], true); err != nil {
			return nil, errors.WithStack(err)
		} else if startCommit == nil {
			return nil, errors.New(fmt.Sprintf("invalid version '%s'", versions[0]))
		} else {
			scope.Start = Range{Commit: startCommit}
			if tag, err := git.GetTagByCommit(startCommit); err != nil {
				return nil, errors.WithStack(err)
			} else if tag != nil {
				scope.Start.Tag = tag
			}
		}

		if endCommit, err = getCommit(git, versions[1], false); err != nil {
			return nil, errors.WithStack(err)
		} else if endCommit == nil {
			return nil, errors.New(fmt.Sprintf("invalid version '%s'", versions[1]))
		} else {
			scope.End = Range{Commit: endCommit}
			if tag, err := git.GetTagByCommit(endCommit); err != nil {
				return nil, errors.WithStack(err)
			} else if tag != nil {
				scope.End.Tag = tag

				lastCommitOfTag, err := git.GetLastCommitOfTag(tag)

				if err != nil {
					return nil, errors.WithStack(err)
				} else if lastCommitOfTag != nil {
					scope.End.Commit = lastCommitOfTag
				}

			}
		}

		tags, err := git.GetTagRangesByCommit(startCommit, endCommit)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		if len(tags) != 0 {
			scope.Tags = tags
		}

		scopeRange = append(scopeRange, scope)
	}

	return scopeRange, nil
}
