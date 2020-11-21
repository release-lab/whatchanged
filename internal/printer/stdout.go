package printer

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
)

func Stdout(version string, commits []*object.Commit) error {
	ctx := transform(version, commits)

	t := template.New("stdout")

	t.Funcs(template.FuncMap{
		"StringsJoin": strings.Join,
	})

	if t, err := t.Parse(defaultTemplate); err != nil {
		return errors.WithStack(err)
	} else {
		b := bytes.NewBuffer([]byte{})
		if err := t.Execute(b, ctx); err != nil {
			return errors.WithStack(err)
		}

		if _, err := io.Copy(os.Stdout, b); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
