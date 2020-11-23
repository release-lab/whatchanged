package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"

	transformer "github.com/axetroy/changelog/3_transformer"
	"github.com/axetroy/changelog/internal/client"
	"github.com/pkg/errors"
	giturls "github.com/whilp/git-urls"
)

func Generate(g *client.GitClient, contexts []*transformer.TemplateContext, format string, preset string, templateFile string) ([]byte, error) {
	remote, err := g.GetRemote()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var remoteURL *url.URL

	if remote != nil || len(remote.URLs) > 0 {
		for _, urlStr := range remote.URLs {
			if remoteURL, err = giturls.Parse(urlStr); err != nil {
				return nil, errors.WithStack(err)
			} else {
				if remoteURL.Host == "github.com" || remoteURL.Host == "gitlab.com" {
					break
				}
			}
		}

		if remoteURL != nil {
			urlPath := strings.TrimSuffix(remoteURL.Path, ".git")
			switch remoteURL.Scheme {
			case "http":
				fallthrough
			case "https":
				if remoteURL, err = url.Parse(fmt.Sprintf("%s://%s/%s", remoteURL.Scheme, remoteURL.Host, urlPath)); err != nil {
					return nil, errors.WithStack(err)
				}
			case "ssh":
				if remoteURL, err = url.Parse(fmt.Sprintf("https://%s/%s", remoteURL.Host, urlPath)); err != nil {
					return nil, errors.WithStack(err)
				}
			}
		}
	}

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
			case "full":
				templateStr = FULL_TEMPLATE
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
				"hashURL": func(longHash string) string {
					short := string(longHash[0:7])
					if remoteURL == nil {
						return short
					}

					u, _ := url.Parse(remoteURL.String())

					u.Path = u.Path + "/commit/" + longHash

					return fmt.Sprintf(`[%s](%s)`, short, u.String())
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
