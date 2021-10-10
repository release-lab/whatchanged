package conventionalcommitparser

// Conventional Commits specification
// https://www.conventionalcommits.org/en/v1.0.0/

import (
	"regexp"
	"strings"
)

type Message struct {
	Header string
	Body   string
	Footer []string
}

var (
	EMPTY_LINE_PATTERN  = regexp.MustCompile(`^\s*$`)
	REVERT_BODY_PATTERN = regexp.MustCompile(`(?i)This\sreverts\scommit\s(\w+)\.?`)
)

func splitToLines(text string) []string {
	return strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
}

func (m Message) ParseHeader() Header {
	return ParseHeader(m.Header)
}

func (m Message) ParseFooter() []Footer {
	footers := make([]Footer, 0)

	for _, m := range m.Footer {
		footers = append(footers, ParseFooter(m))
	}

	header := m.ParseHeader()

	if header.Type == "revert" {
		matcher := REVERT_BODY_PATTERN.FindStringSubmatch(m.Body)
		content := ""

		if len(matcher) > 0 {
			content = matcher[1]
		}

		footer := Footer{
			Tag:     "revert",
			Title:   header.Subject,
			Content: content,
		}

		footers = append(footers, footer)
	}

	return footers
}

func (m Message) GetCloses() []string {
	footer := m.GetFooterByField("Close", "close", "Closes", "closes", "Fix", "fix", "Fixes", "fixes")

	closes := make([]string, 0)

	if footer == nil {
		return nil
	}

	// #1, #2, #3,#4, #110,#123
	arr := strings.Split(footer.Title, ",")

	for _, r := range arr {
		hash := strings.TrimSpace(r)

		if hash != "" {
			closes = append(closes, strings.TrimSpace(r))
		}
	}

	return closes
}

func (m Message) GetFooterByField(tags ...string) *Footer {
	footers := m.ParseFooter()

	for _, tag := range tags {
		for _, f := range footers {
			if strings.EqualFold(f.Tag, tag) {
				return &f
			}
		}
	}

	return nil
}

/**
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]

fix: prevent racing of requests

Introduce a request id and a reference to latest request. Dismiss
incoming responses other than from latest request.

Remove timeouts which were used to mitigate the racing issue but are
obsolete now.

BREAKING CHANGE: use `.use()` instea of `.load()`

before:
```javascript
app.load({})
```

after:
```javascript
app.use({})
```

Reviewed-by: Z
Refs: #123
*/
func Parse(message string) Message {
	var (
		msg    Message
		header string   = ""
		body   []string = make([]string, 0)
		footer []string = make([]string, 0)
	)

	lines := splitToLines(message)
	index := 0

	for {
		// last break
		if index >= len(lines) {
			break
		}

		line := lines[index]

		// The first line should be header
		if index == 0 {
			header = line
			index++
			continue
		}

		// The second line should be blank
		if index == 1 {
			body = append(body, line)
			index++
			continue
		}

		previousLine := lines[index-1]

		// if is a footer start
		if isFooterParagraph(line) && (EMPTY_LINE_PATTERN.MatchString(previousLine) || isFooterParagraph(previousLine)) {
			footerContent := []string{line}

			index++

			// if this line is last line
			if index >= len(lines) {
				footer = append(footer, strings.TrimSpace(strings.Join(footerContent, "\n")))
				continue
			}

		innerLoop:
			for {
				if index >= len(lines) {
					footer = append(footer, strings.TrimSpace(strings.Join(footerContent, "\n")))
					break innerLoop
				}

				line := lines[index]

				// if match the next footer tag
				if isFooterParagraph(line) {
					footer = append(footer, strings.TrimSpace(strings.Join(footerContent, "\n")))
					break innerLoop
				} else {
					footerContent = append(footerContent, line)
					index++
				}
			}
		} else {
			body = append(body, line)
			index++
			continue
		}
	}

	msg.Header = header
	msg.Body = strings.TrimSpace(strings.Join(body, "\n"))
	msg.Footer = footer

	return msg
}
