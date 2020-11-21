package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/axetroy/changelog/internal/client"
	"github.com/axetroy/changelog/internal/printer"
	"github.com/pkg/errors"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

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
		flag.Usage()
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
		commits, err := client.LogsAuto()

		if err != nil {
			return errors.WithStack(err)
		}

		if err := printer.Stdout("unreleased", commits); err != nil {
			return errors.WithStack(err)
		}
	} else {
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
