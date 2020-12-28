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
	println(`whatchanged - a cli to generate changelog from git project

USAGE:
  whatchanged [OPTIONS] [version...]

ARGUMENTS:
  [version...]  Optional version or version range.
                1.null.
                  If you do not specify the version, then it will automatically
                  generate a change log from "HEAD~<latest version>" or
                  "HEAD~<earliest commit>" or "<latest version>-<last version>"
                2.single version. eg. "v1.2.0"
                  Generate a specific version of the changelog.
                3.multiple versions. eg. "v2.0.0 v1.0.0"
                4.version range. eg v1.3.0~v1.2.0
                  Generate changelog within the specified range.
                  For more details, please check the following examples.

OPTIONS:
  --help        Print help information.
  --version     Print version information.

  --project     Specify the project to be generated. It can be a relative path.
                or an absolute path or even a remote Git URL. eg.
                --project=/path/to/project/which/contains/.git/folder
                --project=https://github.com/axetroy/whatchanged.git
                Defaults to "--project=$PWD".
  --output      Write output to file. default write to stdout.

  --fmt         Specify the changelog format. Available options:
                --fmt=md
                --fmt=json
                Defaults to "--fmt=md".
  --preset      Cli built-in markdown template. Available options:
                --preset=default
                --preset=full
                Only available when --fmt=md and --tpl is nil.
                Defaults to "--preset=default".
  --tpl         Specify the template file for generating. Only available when
                --fmt=md.
  --skip-format Skip the formatting process, which is very useful for keeping the
                original format.

EXAMPLES:
  # generate changelog from HEAD to <latest version>.
  # if HEAD is not the latest tag. then this should be a unreleased version
  # otherwise it should be the latest version
  $ whatchanged

  # generate changelog of the specified version
  $ whatchanged v1.2.0

  # Generate the specified two versions
  # Separate by a comma, and only generate these two versions
  # the middle version will not be generated
  $ whatchanged v2.0.0 v1.0.0

  # generate HEAD to latest tag and <Nth tag>
  $ whatchanged HEAD~@0 @1 @2

  # generate changelog within the specified range
  $ whatchanged v1.3.0~v1.2.0

  # generate changelog from HEAD to <Nth tag>
  $ whatchanged ~@0

  # generate changelog from <0th tag> to <2th tag>
  $ whatchanged @0~@2

  # generate changelog from HEAD to specified version
  $ whatchanged HEAD~v1.3.0

  # generate all changelog
  $ whatchanged HEAD~

  # generate changelog from two commit hashes
  $ whatchanged 770ed02~585445d

  # Generate changelog for the specified project
  $ whatchanged --project=/path/to/project v1.0.0

  # Generate changelog for the remote project
  $ whatchanged --project=https://github.com/axetroy/whatchanged.git v0.1.0

SOURCE CODE:
  https://github.com/axetroy/whatchanged`)
}

func run() error {
	var (
		showHelp     bool
		showVersion  bool
		project      string
		preset       string
		templateFile string
		format       string
		outputFile   string
		skipFormat   bool
	)

	cwd, err := os.Getwd()

	if err != nil {
		return errors.WithStack(err)
	}

	flag.StringVar(&project, "project", cwd, "Project filepath or remote URL")
	flag.StringVar(&outputFile, "output", "", "Specify the file to be written")
	flag.StringVar(&format, "fmt", string(option.FormatMarkdown), "The changelog format")
	flag.StringVar(&preset, "preset", string(option.PresetDefault), "Cli built-in markdown template")
	flag.StringVar(&templateFile, "tpl", "", "Specify the template when generating")
	flag.BoolVar(&skipFormat, "skip-format", false, "Skip the formatting process, which is very useful for keeping the original format.")
	flag.BoolVar(&showHelp, "help", false, "Print help information")
	flag.BoolVar(&showVersion, "version", false, "Print version information")

	flag.Usage = printHelp

	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	if showVersion {
		println(fmt.Sprintf("%s %s %s", version, commit, date))
		os.Exit(0)
	}

	ranges := flag.Args()

	var output = os.Stdout

	if outputFile != "" {
		f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)

		if err != nil {
			return errors.WithStack(err)
		}

		defer f.Close()

		output = f
	}

	if err := whatchanged.Generate(project, output, &option.Options{
		Version:      ranges,
		Preset:       option.Preset(preset),
		Format:       option.Format(format),
		SkipFormat:   skipFormat,
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
