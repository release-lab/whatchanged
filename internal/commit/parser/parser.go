package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type Header struct {
	Type    string
	Scope   string
	Subject string
}

type BreakingChange struct {
	Title   string
	Content string
}

type Footer struct {
	BreakingChange *BreakingChange
	Closes         []string
}

/*
feat(http): add form-date params for post method

this PR add 'form-date' params for post method.
Because in certain scenarios

BREAKING CHANGE: `data` is not required field any more.

before:

```javascript
post({ url: '/api/xxx', data: {  } })
```

after:

```javascript
post({ url: '/api/xxx', 'form-data': {  } })
```

Close #101,#109
*/
type Message struct {
	raw    string
	Title  string
	Header *Header
	Body   string
	Footer *Footer
}

var (
	headerrPattern = regexp.MustCompile(`^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-]*)\))?\: (.*)$`)
	footerPattern  = regexp.MustCompile(`^BREAKING\sCHANGE:\s*(.*)$`)
	closePattern   = regexp.MustCompile(`^Closes\s(.+)$`)
)

func (m *Message) String() string {
	return m.raw
}

func Parser(message string) *Message {
	m := Message{
		raw: message,
	}

	var (
		body                  = ""
		isBodyBlock           = false
		breakingChange        = ""
		isBreakingChangeBlock = false
	)

	lines := strings.Split(message, "\n")

	for index, line := range lines {
		if index == 0 {
			m.Title = line
			headerMatchers := headerrPattern.FindStringSubmatch(line)

			if len(headerMatchers) != 0 {
				headerType := headerMatchers[1]
				scope := headerMatchers[3]
				subject := headerMatchers[4]

				m.Header = &Header{}
				m.Header.Type = headerType
				m.Header.Scope = scope
				m.Header.Subject = subject
			}
		}

		if index == 1 {
			isBodyBlock = true
			if len(line) == 0 {
				continue
			}
		}

		breakingChangeFooterMatchers := footerPattern.FindStringSubmatch(line)
		closeFooterMatchers := closePattern.FindStringSubmatch(line)

		if len(closeFooterMatchers) != 0 {
			isBreakingChangeBlock = false
			isBodyBlock = false
			issues := strings.Split(closeFooterMatchers[1], ",")
			if m.Footer == nil {
				m.Footer = &Footer{}
			}
			for i, issue := range issues {
				issues[i] = strings.TrimRight(strings.TrimSpace(issue), ".")
			}
			m.Footer.Closes = issues
			continue
		} else if len(breakingChangeFooterMatchers) != 0 {
			isBreakingChangeBlock = true
			isBodyBlock = false
			if m.Footer == nil {
				m.Footer = &Footer{}
			}
			if m.Footer.BreakingChange == nil {
				m.Footer.BreakingChange = &BreakingChange{}
			}
			m.Footer.BreakingChange.Title = breakingChangeFooterMatchers[1]
			continue
		}

		if isBreakingChangeBlock {
			breakingChange += fmt.Sprintf("%s\n", line)
		} else if isBodyBlock {
			body += fmt.Sprintf("%s\n", line)
		}
	}

	m.Body = strings.Trim(body, "\n")
	if len(breakingChange) > 0 {
		if m.Footer.BreakingChange == nil {
			m.Footer.BreakingChange = &BreakingChange{}
		}
		m.Footer.BreakingChange.Content = strings.Trim(breakingChange, "\n")
	}

	return &m
}
