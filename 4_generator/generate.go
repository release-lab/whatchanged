package generator

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	"github.com/pkg/errors"
	transformer "github.com/release-lab/whatchanged/3_transformer"
	"github.com/release-lab/whatchanged/internal/client"
	"github.com/release-lab/whatchanged/option"
)

//go:embed template/*.tpl
var templateFS embed.FS

func Generate(g *client.GitClient, contexts []*transformer.TemplateContext, format option.Format, preset option.Preset, templateFile string, templateStr string) ([]byte, error) {
	switch format {
	case option.FormatJSON:
		output, err := json.Marshal(contexts)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return output, nil
	case option.FormatMarkdown:
		var output []byte

		if templateStr != "" {
			// ignore
		} else if templateFile == "" {
			switch preset {
			case option.PresetDefault:
				fallthrough
			case option.PresetSimple:
				fallthrough
			case option.PresetFull:
				b, err := templateFS.ReadFile(fmt.Sprintf("template/%s.tpl", preset))

				if err != nil {
					return nil, errors.WithStack(err)
				}

				templateStr = string(b)
			default:
				return nil, errors.Errorf("invalid preset '%s'", preset)
			}
		} else {
			if b, err := ioutil.ReadFile(templateFile); err != nil {
				return nil, err
			} else {
				templateStr = string(b)
			}
		}

		for _, ctx := range contexts {
			t := template.New(ctx.Version)

			t.Funcs(template.FuncMap{
				"unescape": func(s string) template.HTML {
					return template.HTML(s)
				},
			})

			if t, err := t.Parse(templateStr); err != nil {
				return nil, errors.WithStack(err)
			} else {
				b := bytes.NewBuffer(nil)
				if err := t.Execute(b, ctx); err != nil {
					return nil, errors.WithStack(err)
				}

				output = append(output, b.Bytes()...)
			}
		}

		return output, nil
	default:
		return nil, errors.Errorf("invalid format '%s'", format)
	}

}
