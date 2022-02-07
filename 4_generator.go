package whatchanged

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"regexp"

	sprig "github.com/Masterminds/sprig/v3"

	"github.com/pkg/errors"
	"github.com/release-lab/whatchanged/internal/client"
)

//go:embed template/*.tpl
var TemplateFS embed.FS

// Get all regex functions which use a compiled regex cache.
func getRegexFuncs(regexCache map[string]*regexp.Regexp) template.FuncMap {
	return map[string]interface{}{
		// Regex (but cached)

		"regexMatch": func(regex string, s string) bool {
			r, exists := regexCache[regex]
			if !exists {
				r = regexp.MustCompile(regex)
				regexCache[regex] = r
			}
			return r.MatchString(s)
		},

		"mustRegexMatch": func(regex string, s string) (bool, error) {
			r, exists := regexCache[regex]
			if !exists {
				r, err := regexp.Compile(regex)
				if err != nil {
					return false, errors.WithStack(err)
				}
				regexCache[regex] = r
			}
			return r.MatchString(s), nil
		},

		"regexFindAll": func(regex string, s string, n int) []string {
			r, exists := regexCache[regex]
			if !exists {
				r = regexp.MustCompile(regex)
				regexCache[regex] = r
			}
			return r.FindAllString(s, n)
		},

		"mustRegexFindAll": func(regex string, s string, n int) ([]string, error) {
			r, exists := regexCache[regex]
			if !exists {
				r, err := regexp.Compile(regex)
				if err != nil {
					return []string{}, errors.WithStack(err)
				}
				regexCache[regex] = r
			}
			return r.FindAllString(s, n), nil
		},

		"regexFind": func(regex string, s string) string {
			r, exists := regexCache[regex]
			if !exists {
				r = regexp.MustCompile(regex)
				regexCache[regex] = r
			}
			return r.FindString(s)
		},

		"mustRegexFind": func(regex string, s string) (string, error) {
			r, exists := regexCache[regex]
			if !exists {
				r, err := regexp.Compile(regex)
				if err != nil {
					return "", errors.WithStack(err)
				}
				regexCache[regex] = r
			}
			return r.FindString(s), nil
		},

		"regexReplaceAll": func(regex string, s string, repl string) string {
			r, exists := regexCache[regex]
			if !exists {
				r = regexp.MustCompile(regex)
				regexCache[regex] = r
			}
			return r.ReplaceAllString(s, repl)
		},

		"mustRegexReplaceAll": func(regex string, s string, repl string) (string, error) {
			r, exists := regexCache[regex]
			if !exists {
				r, err := regexp.Compile(regex)
				if err != nil {
					return "", errors.WithStack(err)
				}
				regexCache[regex] = r
			}
			return r.ReplaceAllString(s, repl), nil
		},

		"regexReplaceAllLiteral": func(regex string, s string, repl string) string {
			r, exists := regexCache[regex]
			if !exists {
				r = regexp.MustCompile(regex)
				regexCache[regex] = r
			}
			return r.ReplaceAllLiteralString(s, repl)
		},

		"mustRegexReplaceAllLiteral": func(regex string, s string, repl string) (string, error) {
			r, exists := regexCache[regex]
			if !exists {
				r, err := regexp.Compile(regex)
				if err != nil {
					return "", errors.WithStack(err)
				}
				regexCache[regex] = r
			}
			return r.ReplaceAllLiteralString(s, repl), nil
		},

		"regexSplit": func(regex string, s string, n int) []string {
			r, exists := regexCache[regex]
			if !exists {
				r = regexp.MustCompile(regex)
				regexCache[regex] = r
			}
			return r.Split(s, n)
		},

		"mustRegexSplit": func(regex string, s string, n int) ([]string, error) {
			r, exists := regexCache[regex]
			if !exists {
				r, err := regexp.Compile(regex)
				if err != nil {
					return []string{}, errors.WithStack(err)
				}
				regexCache[regex] = r
			}
			return r.Split(s, n), nil
		},
	}
}

func GenerateFromContext(g *client.GitClient, contexts []*TemplateContext, format EnumFormat, preset EnumPreset, templateFile string, templateStr string) ([]byte, error) {

	switch format {
	case FormatJSON:
		output, err := json.Marshal(contexts)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return output, nil
	case FormatMarkdown:

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
				return nil, errors.WithStack(err)
			} else {
				templateStr = string(b)
			}
		}

		regexCache := make(map[string]*regexp.Regexp)

		for _, ctx := range contexts {
			t := template.New(ctx.Version)

			// commom function for sprig
			t.Funcs(sprig.FuncMap())
			// overwrite regex function with cached one for better performance.
			t.Funcs(getRegexFuncs(regexCache))

			// custom function
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
