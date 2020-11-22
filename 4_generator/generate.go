package generator

import (
	"bytes"
	"html/template"
	"strings"

	transform "github.com/axetroy/changelog/3_transform"
	"github.com/axetroy/changelog/internal/client"
	"github.com/pkg/errors"
)

func Generate(g *client.GitClient, contexts []*transform.TemplateContext) ([]byte, error) {
	var output []byte

	for _, ctx := range contexts {
		t := template.New(ctx.Version)

		t.Funcs(template.FuncMap{"StringsJoin": strings.Join})

		if t, err := t.Parse(defaultTemplate); err != nil {
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
}
