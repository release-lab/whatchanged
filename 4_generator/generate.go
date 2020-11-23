package generator

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"

	transformer "github.com/axetroy/changelog/3_transformer"
	"github.com/axetroy/changelog/internal/client"
	"github.com/pkg/errors"
)

func Generate(g *client.GitClient, contexts []*transformer.TemplateContext, format string, preset string, templateFile string) ([]byte, error) {

	switch format {
	case "json":
		output, err := json.Marshal(contexts)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return output, nil
	case "md":
		var output []byte

		var templateStr string

		cwd, err := os.Getwd()

		if err != nil {
			return nil, errors.WithStack(err)
		}

		if templateFile != "" {
			if !path.IsAbs(templateFile) {
				templateFile = path.Join(cwd, templateFile)
			}
			fileBytes, err := ioutil.ReadFile(templateFile)

			if err != nil {
				return nil, errors.WithStack(err)
			}

			templateStr = string(fileBytes)
		} else {
			switch preset {
			case "default":
				templateStr = DEFAULT_TEMPLATE
			default:
				return nil, errors.Errorf("invalid preset '%s'", preset)
			}
		}

		for _, ctx := range contexts {
			t := template.New(ctx.Version)

			t.Funcs(template.FuncMap{
				"stringsJoin": strings.Join,
				"unescape": func(s string) template.HTML {
					return template.HTML(s)
				},
			})

			if t, err := t.Parse(templateStr); err != nil {
				return nil, errors.WithStack(err)
			} else {
				b := bytes.NewBuffer([]byte{})
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
