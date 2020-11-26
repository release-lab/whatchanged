package whatchanged

import (
	"io"
	"os"
	"path"

	parser "github.com/axetroy/whatchanged/1_parser"
	extractor "github.com/axetroy/whatchanged/2_extractor"
	transformer "github.com/axetroy/whatchanged/3_transformer"
	generator "github.com/axetroy/whatchanged/4_generator"
	formatter "github.com/axetroy/whatchanged/5_formatter"
	writer "github.com/axetroy/whatchanged/6_writer"
	"github.com/axetroy/whatchanged/internal/client"
	"github.com/axetroy/whatchanged/option"
	"github.com/pkg/errors"
)

func Generate(project string, w io.Writer, options *option.Options) error {
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

	client, err := client.NewGitClient(project)

	if err != nil {
		return errors.WithStack(err)
	}

	scope, err := parser.Parse(client, options.Version)

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

	output, err := generator.Generate(client, ctxs, options.Format, options.Preset, options.TemplateFile)

	if err != nil {
		return errors.WithStack(err)
	}

	formattedOutput, err := formatter.Format(output, options.Format, options.TemplateFile)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = writer.Write(formattedOutput, w)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
