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

type Header struct {
	Type      string
	Scope     string
	Subject   string
	Important bool
}

type Footer struct {
	Tag     string
	Title   string
	Content string
}

var (
	EMPTY_LINE_PATTERN             = regexp.MustCompile(`^\s*$`)
	HEADER_PATTERN                 = regexp.MustCompile(`^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-]*)\))?(!?):\s(.*)$`)
	FOOTER_TAG_PATTERN             = regexp.MustCompile(`(?i)^([a-z]+(-[a-z]+)*):\s?(.*)$`)
	FOOTER_HASH_PATTERN            = regexp.MustCompile(`^(?i)^([a-z][a-z]+)\s(((,\s*)?#\d+(,\s)?)+)$`)
	FOOTER_BREAKING_CHANGE_PATTERN = regexp.MustCompile(`^(BREAKING\sCHANGE):\s?(.*)$`)
	REVERT_HEADER_PATTERN          = regexp.MustCompile(`^(?i)revert\s(.*)$`)
	REVERT_BODY_PATTERN            = regexp.MustCompile(`(?i)This\sreverts\scommit\s(\w+)\.?`)
)

func paseFooterParagraph(txt string) Footer {
	footer := Footer{}

	tagMatcher := FOOTER_TAG_PATTERN.FindStringSubmatch(txt)
	breakingChangeMatcher := FOOTER_BREAKING_CHANGE_PATTERN.FindStringSubmatch(txt)
	hashTagMatcher := FOOTER_HASH_PATTERN.FindStringSubmatch(txt)

	if len(breakingChangeMatcher) != 0 {
		footer.Tag = strings.TrimSpace(breakingChangeMatcher[1])
		footer.Title = strings.TrimSpace(breakingChangeMatcher[2])
	} else if len(tagMatcher) != 0 {
		footer.Tag = strings.TrimSpace(tagMatcher[1])
		footer.Title = strings.TrimSpace(tagMatcher[3])
	} else if len(hashTagMatcher) != 0 {
		footer.Tag = strings.TrimSpace(hashTagMatcher[1])
		footer.Title = strings.TrimSpace(hashTagMatcher[2])
	} else {
		footer.Tag = ""
		footer.Title = txt
	}

	return footer
}

func isFooterParagraph(txt string) bool {
	if FOOTER_BREAKING_CHANGE_PATTERN.MatchString(txt) {
		return true
	}

	if FOOTER_TAG_PATTERN.MatchString(txt) {
		return true
	}

	if FOOTER_HASH_PATTERN.MatchString(txt) {
		return true
	}

	return false
}

func splitToLines(text string) []string {
	return strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
}

func (m Message) ParseHeader() Header {
	headerMatchers := HEADER_PATTERN.FindStringSubmatch(m.Header)
	revertHeaderMatchers := REVERT_HEADER_PATTERN.FindStringSubmatch(m.Header)
	header := Header{}

	if len(headerMatchers) != 0 { // conventional commit
		header.Type = strings.ToLower(headerMatchers[1])
		header.Scope = headerMatchers[3]
		header.Important = headerMatchers[4] == "!"
		header.Subject = headerMatchers[5]
	} else if len(revertHeaderMatchers) != 0 { // revert commit
		subject := strings.Trim(revertHeaderMatchers[1], "\"")
		subject = strings.Trim(subject, "'")
		header.Type = "revert"
		header.Subject = subject
	} else { // commom commit
		header.Type = ""
		header.Scope = ""
		header.Subject = m.Header
	}

	return header
}

func (m Message) ParseFooter() []Footer {
	footers := make([]Footer, 0)

	for _, m := range m.Footer {
		lines := splitToLines(m)

		footer := Footer{}

		contents := make([]string, 0)

	lineLoop:
		for index, line := range lines {
			if index == 0 {
				f := paseFooterParagraph(line)
				footer.Tag = f.Tag
				footer.Title = f.Title
				continue lineLoop
			} else {
				contents = append(contents, line)
			}
		}

		footer.Content = strings.TrimSpace(strings.Join(contents, "\n"))

		footers = append(footers, footer)
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
