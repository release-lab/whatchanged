package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/axetroy/whatchanged"
	"github.com/axetroy/whatchanged/option"
	"github.com/pkg/errors"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printHelp() {
	fmt.Println(`whatchanged - a cli to generate changelog from git project

USAGE:
  whatchanged [OPTIONS] [version]

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
  # generate changelog from HEAD to <latest version>. equivalent to 'whatchanged HEAD~tag:0'
  $ whatchanged

  # generate changelog of the specified version
  $ whatchanged v1.2.0

  # generate changelog within the specified range
  $ whatchanged v1.3.0~v1.2.0

  # generate changelog from HEAD to <Nth tag>
  $ whatchanged ~tag:0

  # generate changelog from <0th tag> to <2th tag>
  $ whatchanged tag:0~tag:2

  # generate changelog from HEAD to specified version
  $ whatchanged HEAD~v1.3.0

  # generate all changelog
  $ whatchanged HEAD~

  # generate changelog from two commit hashes
  $ whatchanged 770ed02~585445d

  # Generate changelog for the specified project
  $ whatchanged --dir=/path/to/project v1.0.0

SOURCE CODE:
  https://github.com/axetroy/whatchanged`)
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

	var output = os.Stdout

	if outputFile != "" {
		f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)

		if err != nil {
			return errors.WithStack(err)
		}

		defer f.Close()

		output = f
	}

	if err := whatchanged.Generate(projectDir, output, &option.Options{
		Version:      version,
		Preset:       option.Preset(preset),
		Format:       option.Format(format),
		TemplateFile: templateFile,
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func main() {
	var (
		err error
	)

	defer func() {
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(255)
		}
	}()

	err = run()
}
