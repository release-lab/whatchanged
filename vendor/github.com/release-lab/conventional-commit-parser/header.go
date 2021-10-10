package conventionalcommitparser

import (
	"regexp"
	"strings"
)

type Header struct {
	Type      string
	Scope     string
	Subject   string
	Important bool
}

var (
	HEADER_PATTERN        = regexp.MustCompile(`^(?i)([\s\w-]*)(\((.*)\))?(!?):\s+(.*)$`)
	REVERT_HEADER_PATTERN = regexp.MustCompile(`^(?i)revert\s(.*)$`)
)

func ParseHeader(txt string) Header {
	headerMatchers := HEADER_PATTERN.FindStringSubmatch(txt)
	revertHeaderMatchers := REVERT_HEADER_PATTERN.FindStringSubmatch(txt)
	header := Header{}

	if len(headerMatchers) != 0 { // conventional commit
		header.Type = strings.TrimSpace(strings.ToLower(headerMatchers[1]))
		header.Scope = strings.TrimSpace(headerMatchers[3])
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
		header.Subject = txt
	}

	return header
}
