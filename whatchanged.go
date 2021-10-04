package whatchanged

import (
	"context"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"
	parser "github.com/whatchanged-community/whatchanged/1_parser"
	extractor "github.com/whatchanged-community/whatchanged/2_extractor"
	transformer "github.com/whatchanged-community/whatchanged/3_transformer"
	generator "github.com/whatchanged-community/whatchanged/4_generator"
	formatter "github.com/whatchanged-community/whatchanged/5_formatter"
	writer "github.com/whatchanged-community/whatchanged/6_writer"
	"github.com/whatchanged-community/whatchanged/internal/client"
	"github.com/whatchanged-community/whatchanged/option"
)

var (
	gitHTTPURLReg = regexp.MustCompile(`^https?:\/\/.+$`)
	gitSSHURLReg  = regexp.MustCompile(`^git@.+$`)
)

// generate changelog
// project: it could be a file path or a git URL
func Generate(ctx context.Context, project string, w io.Writer, options *option.Options) error {
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

	if options.Branch == "" {
		options.Branch = "master"
	}

	if options.TemplateFile != "" {
		cwd, err := os.Getwd()

		if err != nil {
			return errors.WithStack(err)
		}

		if !filepath.IsAbs(options.TemplateFile) {
			options.TemplateFile = filepath.Join(cwd, options.TemplateFile)
		}
	}

	var g *client.GitClient

	if gitHTTPURLReg.MatchString(project) || gitSSHURLReg.MatchString(project) {
		cloneOptions := git.CloneOptions{
			URL:           project,
			Progress:      os.Stderr,
			SingleBranch:  true,
			ReferenceName: plumbing.NewBranchReferenceName(options.Branch),
			Tags:          git.AllTags,
		}

		if gitHTTPURLReg.MatchString(project) {
			u, err := url.Parse(project)

			if err != nil {
				return errors.WithStack(err)
			}

			// if url has contains the user infomation
			if u.User != nil {
				auth := &http.BasicAuth{
					Username: u.User.Username(),
				}

				if pwd, isSetPwd := u.User.Password(); isSetPwd {
					auth.Password = pwd
				}

				cloneOptions.Auth = auth
			}

		}

		repo, err := git.CloneContext(ctx, memory.NewStorage(), nil, &cloneOptions)

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

	if options.SkipFormat {
		_, err = writer.Write(output, w)

		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	} else {
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
}
