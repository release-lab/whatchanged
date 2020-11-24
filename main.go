package main

import (
	"flag"
	"fmt"
	"os"

	parser "github.com/axetroy/changelog/1_parser"
	extractor "github.com/axetroy/changelog/2_extractor"
	transformer "github.com/axetroy/changelog/3_transformer"
	generator "github.com/axetroy/changelog/4_generator"
	formatter "github.com/axetroy/changelog/5_formatter"
	writer "github.com/axetroy/changelog/6_writer"
	"github.com/axetroy/changelog/internal/client"
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
  --file        Write output to file. default write to stdout.
  --fmt         The changelog format. Available options are "md"/"json".
                Defaults to "md".
  --preset      Cli built-in markdown template. Available options are "default"
                /"full".
                Only available when --fmt=md and --tpl is nil. Defaults to
                "default".
  --tpl         Specify the template file for generating. Only available when --fmt=md.

EXAMPLES:
  # generate changelog from HEAD to <latest version>. equivalent to 'changelog HEAD~tag:0'
  $ changelog

  # generate changelog of the specified version
  $ changelog v1.2.0

  # generate changelog within the specified range
  $ changelog v1.3.0~v1.2.0

  # generate changelog from HEAD to <Nth tag>
  $ changelog ~tag:0

  # generate changelog from <0th tag> to <2th tag>
  $ changelog tag:0~tag:2

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
		showHelp     bool
		showVersion  bool
		projectDir   string
		preset       string
		templateFile string
		format       string
		outputFile   string
	)

	cwd, err := os.Getwd()

	if err != nil {
		return errors.WithStack(err)
	}

	flag.StringVar(&projectDir, "dir", cwd, "Project dir")
	flag.StringVar(&format, "fmt", "md", "The changelog format")
	flag.StringVar(&preset, "preset", "default", "Cli built-in markdown template")
	flag.StringVar(&templateFile, "tpl", "", "Specify the template when generating")
	flag.StringVar(&outputFile, "file", "", "Specify output result to file.")
	flag.BoolVar(&showHelp, "help", false, "Print help information")
	flag.BoolVar(&showVersion, "version", false, "Print version information")

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

	scope, err := parser.Parse(client, version)

	if err != nil {
		return errors.WithStack(err)
	}

	splices, err := extractor.Extract(client, scope)

	if err != nil {
		return errors.WithStack(err)
	}

	ctxs, err := transformer.Transform(client, splices)

	if err != nil {
		return errors.WithStack(err)
	}

	output, err := generator.Generate(client, ctxs, format, preset, templateFile)

	if err != nil {
		return errors.WithStack(err)
	}

	formattedOutput, err := formatter.Format(output, format, templateFile)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = writer.Write(formattedOutput, outputFile)

	if err != nil {
		return errors.WithStack(err)
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
