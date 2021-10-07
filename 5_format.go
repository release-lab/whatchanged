package whatchanged

import (
	"github.com/88250/lute"
	"github.com/pkg/errors"
)

func Format(src []byte, format EnumFormat) ([]byte, error) {
	switch format {
	case FormatJSON:
		return src, nil
	case FormatMarkdown:
		luteEngine := lute.New()
		return luteEngine.Format("CHANGELOG.md", src), nil
	default:
		return nil, errors.Errorf("invalid format '%s'", format)
	}

}
