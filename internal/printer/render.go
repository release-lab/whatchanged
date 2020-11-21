package printer

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
)

func Bytes(version string, commits []*object.Commit) ([]byte, error) {
	ctx := transform(version, commits)

	t := template.New("bytes")

	t.Funcs(template.FuncMap{"StringsJoin": strings.Join})

	if t, err := t.Parse(defaultTemplate); err != nil {
		return nil, errors.WithStack(err)
	} else {
		b := bytes.NewBuffer([]byte{})
		if err := t.Execute(b, ctx); err != nil {
			return nil, errors.WithStack(err)
		}

		return b.Bytes(), nil
	}
}
