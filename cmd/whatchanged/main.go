package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/cloudfoundry/jibber_jabber"
	"github.com/pkg/errors"
	"github.com/release-lab/whatchanged"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

//go:embed help_cn.txt
var helpCN string

//go:embed help_en.txt
var helpEN string

func printHelp() {
	userLocale, _ := jibber_jabber.DetectIETF()
	if userLocale == "zh-CN" || userLocale == "zh-TW" || userLocale == "zh-HK" {
		println(helpCN)
	} else {
		println(helpEN)
	}
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
		branch       string
		skipFormat   bool
	)

	cwd, err := os.Getwd()

	if err != nil {
		return errors.WithStack(err)
	}

	flag.StringVar(&project, "project", cwd, "Project filepath or remote URL")
	flag.StringVar(&outputFile, "output", "", "Specify the file to be written")
	flag.StringVar(&branch, "branch", "master", "Specify branch for remote git project")
	flag.StringVar(&format, "fmt", string(whatchanged.FormatMarkdown), "The changelog format")
	flag.StringVar(&preset, "preset", string(whatchanged.PresetDefault), "Cli built-in markdown template")
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

	if err := whatchanged.Generate(context.Background(), project, output, &whatchanged.Options{
		Version:      ranges,
		Branch:       branch,
		Preset:       whatchanged.EnumPreset(preset),
		Format:       whatchanged.EnumFormat(format),
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
