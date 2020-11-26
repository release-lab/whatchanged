package formatter

import (
	"github.com/axetroy/whatchanged/option"
	"github.com/pkg/errors"
	"github.com/shurcooL/markdownfmt/markdown"
)

func Format(src []byte, format option.Format, templateFile string) ([]byte, error) {

	switch format {
	case option.FormatJSON:
		return src, nil
	case option.FormatMarkdown:
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
