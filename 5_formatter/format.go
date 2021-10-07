package formatter

import (
	"github.com/88250/lute"
	"github.com/pkg/errors"
	"github.com/release-lab/whatchanged/option"
)

func Format(src []byte, format option.Format) ([]byte, error) {
	switch format {
	case option.FormatJSON:
		return src, nil
	case option.FormatMarkdown:
		luteEngine := lute.New()
		return luteEngine.Format("CHANGELOG.md", src), nil
	default:
		return nil, errors.Errorf("invalid format '%s'", format)
	}

}
