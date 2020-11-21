package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/axetroy/changelog/internal/client"
	"github.com/axetroy/changelog/internal/printer"
	"github.com/pkg/errors"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printHelp() {
	fmt.Println(`changelog - a cli to generate changelog from git project

USAGE:
  changelog [OPTIONS] [version]

ARGUMENTS:
  [version]     Optional version or version range.
                1.null.
                  If you do not specify the version, then it will automatically
                  generate a change log from "HEAD~<latest version>" or
                  "HEAD~<earliest commit>" or "<latest version>-<last version>"
                2.single version. eg. v1.2.0
                  Generate a specific version of the changelog.
                3.version range. eg v1.3.0~v1.2.0
                  Generate changelog within the specified range.
                  For more details, please check the following examples.

OPTIONS:
  --help        Print help information.
  --version     Print version information.
  --dir         Specify the directory to be generated.
                The directory should contain a .git folder. defaults to $PWD.
  --tpl         Specify the directory to be generated.

EXAMPLES:
  # generate changelog from HEAD to <latest version>
  $ changelog

  # generate changelog of the specified version
  $ changelog v1.2.0

  # generate changelog within the specified range
  $ changelog v1.3.0~v1.2.0

  # generate changelog from HEAD to specified version
  $ changelog HEAD~v1.3.0

  # generate all changelog
  $ changelog HEAD~

  # generate changelog from two commit hashes
  $ changelog 770ed02~585445d

  # Generate changelog for the specified project
  $ changelog --dir=/path/to/project v1.0.0

SOURCE CODE:
  https://github.com/axetroy/changelog`)
}

func run() error {
	var (
		showHelp    bool
		showVersion bool
		projectDir  string
	)

	cwd, err := os.Getwd()

	if err != nil {
		return errors.WithStack(err)
	}
	flag.StringVar(&projectDir, "dir", cwd, "project dir")
	flag.StringVar(&projectDir, "tpl", cwd, "TODO: generate changelog with template")
	flag.BoolVar(&showHelp, "help", false, "print help information")
	flag.BoolVar(&showVersion, "version", false, "print version information")

	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	if showVersion {
		fmt.Printf("%s %s %s\n", version, commit, date)
		os.Exit(0)
	}

	version := flag.Arg(0)

	client, err := client.NewGitClient(projectDir)

	if err != nil {
		return errors.WithStack(err)
	}

	// If no tag is specified, it will be generate automatically
	if version == "" {
		head, err := client.Head()

		if err != nil {
			return errors.WithStack(err)
		}

		latestTag, err := client.TagN(0)

		if err != nil {
			return errors.WithStack(err)
		}

		if latestTag == nil {
			// if there is not tag
			// then it's unreleased
			commits, err := client.Logs(head.Hash().String(), "")

			if err != nil {
				return errors.WithStack(err)
			}

			if err := printer.Stdout("unreleased", commits); err != nil {
				return errors.WithStack(err)
			}
			return nil
		} else {
			if latestTag.Commit.Hash.String() == head.Hash().String() {
				// if the current head is the latest tag
				nextTag, err := client.NextTag(latestTag)

				if err != nil {
					return errors.WithStack(err)
				}

				toHash := ""

				if nextTag != nil {
					toHash = nextTag.Commit.Hash.String()
				}

				commits, err := client.Logs(latestTag.Commit.Hash.String(), toHash)

				if err != nil {
					return errors.WithStack(err)
				}

				if err := printer.Stdout(latestTag.Name, commits); err != nil {
					return errors.WithStack(err)
				}

				return nil
			} else {
				toHash := latestTag.Commit.Hash.String()

				commits, err := client.Logs(latestTag.Commit.Hash.String(), toHash)

				if err != nil {
					return errors.WithStack(err)
				}

				if err := printer.Stdout(latestTag.Name, commits); err != nil {
					return errors.WithStack(err)
				}
				return nil
			}
		}
	} else {
		ranges := strings.Split(version, "~")
		length := len(ranges)

		// handle version range
		// v2.0.0~v1.0.0
		// HEAD~
		if length == 2 {
			tags, err := client.GetTagRangesByTagName(ranges[0], ranges[1])

			if err != nil {
				return errors.WithStack(err)
			}

			var output []byte

			for index, tag := range tags {
				fromHash := tag.Commit.Hash.String()
				toHash := ""
				// if not last element
				if index != len(tags)-1 {
					toHash = tags[index+1].Commit.Hash.String()
				} else {
					nextTag, err := client.NextTag(tag)

					if err != nil {
						return errors.WithStack(err)
					}

					if nextTag != nil {
						toHash = nextTag.Commit.Hash.String()
					}
				}

				commits, err := client.Logs(fromHash, toHash)

				if err != nil {
					return errors.WithStack(err)
				}

				if b, err := printer.Bytes(tag.Name, commits); err != nil {
					return errors.WithStack(err)
				} else {
					output = append(output, b...)
				}
			}

			_, err = io.Copy(os.Stdout, bytes.NewBuffer(output))

			if err != nil {
				return errors.WithStack(err)
			}
		} else {
			// only output one version of the changelog
			from, err := client.TagName(version)

			if err != nil {
				return errors.WithStack(err)
			}

			if from == nil {
				return errors.New(fmt.Sprintf("can not found tag %s", version))
			}

			to, err := client.NextTag(from)

			if err != nil {
				return errors.WithStack(err)
			}

			toHash := ""

			// if don't have next tag
			if to != nil {
				toHash = to.Commit.Hash.String()
			}

			commits, err := client.Logs(from.Commit.Hash.String(), toHash)

			if err != nil {
				return errors.WithStack(err)
			}

			if err := printer.Stdout(version, commits); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}

func main() {
	var (
		err error
	)

	defer func() {
		// if r := recover(); r != nil {
		// 	fmt.Printf("%+v\n", r)
		// 	os.Exit(255)
		// }

		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(255)
		}
	}()

	err = run()
}
