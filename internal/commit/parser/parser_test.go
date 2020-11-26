package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var revertHash = "667ecc1654a317a13331b17617d973392f415f02"

func TestParser(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{
			name: "no format",
			args: args{message: "hello world"},
			want: &Message{
				Title:  "hello world",
				Header: nil,
				Body:   "",
			},
		},
		{
			name: "no format with body",
			args: args{message: `hello world

this is body`},
			want: &Message{
				Title:  "hello world",
				Header: nil,
				Body:   "this is body",
			},
		},
		{
			name: "revert style 1",
			args: args{message: `revert: feat(pencil): add 'graphiteWidth' option

This reverts commit 667ecc1654a317a13331b17617d973392f415f02.`},
			want: &Message{
				Title: "revert: feat(pencil): add 'graphiteWidth' option",
				Header: &Header{
					Type:    "revert",
					Scope:   "",
					Subject: "feat(pencil): add 'graphiteWidth' option",
				},
				Body:   "This reverts commit 667ecc1654a317a13331b17617d973392f415f02.",
				Revert: &revertHash,
			},
		},
		{
			name: "revert style 2",
			args: args{message: `revert test1

This reverts commit 667ecc1654a317a13331b17617d973392f415f02.`},
			want: &Message{
				Title: "revert test1",
				Header: &Header{
					Type:    "revert",
					Scope:   "",
					Subject: "test1",
				},
				Body:   "This reverts commit 667ecc1654a317a13331b17617d973392f415f02.",
				Revert: &revertHash,
			},
		},
		{
			name: "revert style 3",
			args: args{message: `revert "test1"

This reverts commit 667ecc1654a317a13331b17617d973392f415f02.`},
			want: &Message{
				Title: `revert "test1"`,
				Header: &Header{
					Type:    "revert",
					Scope:   "",
					Subject: `test1`,
				},
				Body:   "This reverts commit 667ecc1654a317a13331b17617d973392f415f02.",
				Revert: &revertHash,
			},
		},
		{
			name: "revert style 4",
			args: args{message: `Revert "test1"

This reverts commit 667ecc1654a317a13331b17617d973392f415f02.`},
			want: &Message{
				Title: `Revert "test1"`,
				Header: &Header{
					Type:    "revert",
					Scope:   "",
					Subject: `test1`,
				},
				Body:   "This reverts commit 667ecc1654a317a13331b17617d973392f415f02.",
				Revert: &revertHash,
			},
		},
		{
			name: "basic",
			args: args{message: "feat: add options for function"},
			want: &Message{
				Title: "feat: add options for function",
				Header: &Header{
					Type:    "feat",
					Scope:   "",
					Subject: "add options for function",
				},
				Body: "",
			},
		},
		{
			name: "with scope",
			args: args{message: "feat(http): xxx"},
			want: &Message{
				Title: "feat(http): xxx",
				Header: &Header{
					Type:    "feat",
					Scope:   "http",
					Subject: "xxx",
				},
				Body: "",
			},
		},
		{
			name: "with body",
			args: args{message: `feat(http): xxx

this is commit body
		`},
			want: &Message{
				Title: "feat(http): xxx",
				Header: &Header{
					Type:    "feat",
					Scope:   "http",
					Subject: "xxx",
				},
				Body: `this is commit body
		`,
			},
		},
		{
			name: "multiple body",
			args: args{message: `feat(http): xxx

this is commit body
this is the second commit body
		`},
			want: &Message{
				Title: "feat(http): xxx",
				Header: &Header{
					Type:    "feat",
					Scope:   "http",
					Subject: "xxx",
				},
				Body: `this is commit body
this is the second commit body
		`,
			},
		},
		{
			name: "multiple part body",
			args: args{message: `feat(http): xxx

this is commit body

this is the second commit body

this is the third commit body
		`},
			want: &Message{
				Title: "feat(http): xxx",
				Header: &Header{
					Type:    "feat",
					Scope:   "http",
					Subject: "xxx",
				},
				Body: `this is commit body

this is the second commit body

this is the third commit body
		`,
			},
		},
		{
			name: "breaking change",
			args: args{message: `feat(http): xxx

this is commit body

this is the second commit body

this is the third commit body

BREAKING CHANGE: this is breaking change

before: http.post

after: http.request
		`},
			want: &Message{
				Title: "feat(http): xxx",
				Header: &Header{
					Type:    "feat",
					Scope:   "http",
					Subject: "xxx",
				},
				Body: `this is commit body

this is the second commit body

this is the third commit body`, Footer: &Footer{
					BreakingChange: &BreakingChange{
						Title: "this is breaking change",
						Content: `before: http.post

after: http.request
		`,
					},
				},
			},
		},
		{
			name: "breaking change and close issues",
			args: args{message: `feat(http): xxx

this is commit body

this is the second commit body

this is the third commit body

BREAKING CHANGE: this is breaking change

		before: http.post

		after: http.request

Closes #102, #103,  #104
		`},
			want: &Message{
				Title: "feat(http): xxx",
				Header: &Header{
					Type:    "feat",
					Scope:   "http",
					Subject: "xxx",
				},
				Body: `this is commit body

this is the second commit body

this is the third commit body`, Footer: &Footer{
					BreakingChange: &BreakingChange{
						Title: "this is breaking change",
						Content: `		before: http.post

		after: http.request`,
					},
					Closes: []string{"#102", "#103", "#104"},
				},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := Parser(tt.args.message)
			assert.Equal(t, tt.want.Title, got.Title, "title")
			assert.Equal(t, tt.want.Header, got.Header, "header")
			assert.Equal(t, tt.want.Body, got.Body, "body")
			assert.Equal(t, tt.want.Footer, got.Footer, "footer")
		})
	}
}
