package writer

import (
	"bytes"
	"io"
	"os"
	"path"

	"github.com/pkg/errors"
)

func Write(src []byte, templateFile string) (int64, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return 0, errors.WithStack(err)
	}

	var writer io.Writer

	if templateFile != "" {
		if !path.IsAbs(templateFile) {
			templateFile = path.Join(cwd, templateFile)
		}
		file, err := os.Create(templateFile)

		if err != nil {
			return 0, errors.WithStack(err)
		}

		writer = file
	} else {
		writer = os.Stdout
	}

	n, err := io.Copy(writer, bytes.NewBuffer(src))

	if err != nil {
		return n, errors.WithStack(err)
	}

	return n, nil
}
