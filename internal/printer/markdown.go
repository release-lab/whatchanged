package printer

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func Markdown(commits []*object.Commit) error {
	for _, c := range commits {
		scanner := bufio.NewScanner(strings.NewReader(c.Message))
		scanner.Scan()
		title := strings.TrimSpace(scanner.Text())
		shotHash := string(c.Hash.String()[0:7])
		line := fmt.Sprintf(`%s %s`, shotHash, title)
		fmt.Println(line)
	}

	return nil
}
