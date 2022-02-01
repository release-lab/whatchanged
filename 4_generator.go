package whatchanged

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"regexp"

	"github.com/pkg/errors"
	"github.com/release-lab/whatchanged/internal/client"
)

//go:embed template/*.tpl
var TemplateFS embed.FS

func GenerateFromContext(g *client.GitClient, contexts []*TemplateContext, format EnumFormat, preset EnumPreset, templateFile string, templateStr string) ([]byte, error) {

	switch format {
	case FormatJSON:
		output, err := json.Marshal(contexts)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return output, nil
	case FormatMarkdown:
		regexCache := make(map[string]*regexp.Regexp)
		var output []byte

		if templateStr != "" {
			// ignore
		} else if templateFile == "" {
			switch preset {
			case PresetDefault:
				fallthrough
			case PresetFull:
				b, err := TemplateFS.ReadFile(fmt.Sprintf("template/%s.tpl", preset))

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
				"gsub": func(pattern string, repl string, s string) string {
					v, exists := regexCache[pattern]
					if !exists {
						v = regexp.MustCompile(pattern)
						regexCache[pattern] = v
					}
					return v.ReplaceAllString(s, repl)
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
