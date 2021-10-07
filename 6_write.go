package whatchanged

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
)

func Write(src []byte, w io.Writer) (int64, error) {
	n, err := io.Copy(w, bytes.NewBuffer(src))

	if err != nil {
		return n, errors.WithStack(err)
	}

	return n, nil
}
