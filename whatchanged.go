package whatchanged

import (
	"context"
	"io"
	"os"
	"path"
	"regexp"

	parser "github.com/axetroy/whatchanged/1_parser"
	extractor "github.com/axetroy/whatchanged/2_extractor"
	transformer "github.com/axetroy/whatchanged/3_transformer"
	generator "github.com/axetroy/whatchanged/4_generator"
	formatter "github.com/axetroy/whatchanged/5_formatter"
	writer "github.com/axetroy/whatchanged/6_writer"
	"github.com/axetroy/whatchanged/internal/client"
	"github.com/axetroy/whatchanged/option"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"
)

var (
	gitHTTPURLReg = regexp.MustCompile(`^http?s:\/\/.+$`)
	gitSSHURLReg  = regexp.MustCompile(`^git@.+$`)
)

// generate changelog
// project: it could be a file path or a git URL
func Generate(project string, w io.Writer, options *option.Options) error {
	var (
		err error
	)

	if options == nil {
		options = &option.Options{
			Format: option.FormatMarkdown,
			Preset: option.PresetDefault,
		}
	} else {
		if options.Format == "" {
			options.Format = option.FormatMarkdown
		}
		if options.Preset == "" {
			options.Preset = option.PresetDefault
		}
	}

	if options.TemplateFile != "" {
		cwd, err := os.Getwd()

		if err != nil {
			return errors.WithStack(err)
		}

		if !path.IsAbs(options.TemplateFile) {
			options.TemplateFile = path.Join(cwd, options.TemplateFile)
		}
	}

	var g *client.GitClient

	if gitHTTPURLReg.MatchString(project) || gitSSHURLReg.MatchString(project) {
		repo, err := git.CloneContext(context.Background(), memory.NewStorage(), nil, &git.CloneOptions{
			URL:          project,
			Progress:     os.Stderr,
			SingleBranch: true,
			Tags:         git.AllTags,
		})

		if err != nil {
			return errors.WithStack(err)
		}

		g = &client.GitClient{
			Repository: repo,
		}
	} else {
		g, err = client.NewGitClient(project)

		if err != nil {
			return errors.WithStack(err)
		}
	}

	scope, err := parser.Parse(g, options.Version)

	if err != nil {
		return errors.WithStack(err)
	}

	splices, err := extractor.Extract(g, scope)

	if err != nil {
		return errors.WithStack(err)
	}

	ctxs, err := transformer.Transform(g, splices)

	if err != nil {
		return errors.WithStack(err)
	}
	output, err := generator.Generate(g, ctxs, options.Format, options.Preset, options.TemplateFile, options.Template)

	if err != nil {
		return errors.WithStack(err)
	}

	formattedOutput, err := formatter.Format(output, options.Format)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = writer.Write(formattedOutput, w)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
