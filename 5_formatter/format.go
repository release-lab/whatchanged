package formatter

import (
	"github.com/pkg/errors"
	"github.com/shurcooL/markdownfmt/markdown"
)

func Format(src []byte, format string, templateFile string) ([]byte, error) {

	switch format {
	case "json":
		return src, nil
	case "md":
		var (
			filename   string
			isTerminal = true
		)
		if templateFile == "" {
			filename = "CHANGELOG.md"
			isTerminal = false
		}
		return markdown.Process(filename, src, &markdown.Options{
			Terminal: isTerminal,
		})
	default:
		return nil, errors.Errorf("invalid format '%s'", format)
	}

}
