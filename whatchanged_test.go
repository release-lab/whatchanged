package whatchanged

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/axetroy/whatchanged/option"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	type TestCase struct {
		name       string
		options    *option.Options
		Project    string
		ResultFile string
		wantErr    error
	}

	var tests = []TestCase{
		{
			name:       "vscode-deno",
			Project:    "./__test__/vscode-deno",
			ResultFile: "./__test__/vscode-deno.CHANGELOG.md",
			options: &option.Options{
				Version: []string{"HEAD~"},
			},
		},
		{
			name:       "v",
			Project:    "./__test__/v",
			ResultFile: "./__test__/v.CHANGELOG.md",
			options: &option.Options{
				Version: []string{"HEAD~"},
			},
		},
		{
			name:       "whatchanged single version",
			Project:    "./",
			ResultFile: "./__test__/whatchanged-[v0.2.0].CHANGELOG.md",
			options: &option.Options{
				Version: []string{"v0.2.0"},
			},
		},
		{
			name:       "whatchanged version range",
			Project:    "./",
			ResultFile: "./__test__/whatchanged-[v0.2.0~v0.1.0].CHANGELOG.md",
			options: &option.Options{
				Version: []string{"v0.2.0~v0.1.0"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := Generate(tt.Project, w, tt.options); err != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if b, err := ioutil.ReadFile(tt.ResultFile); err != nil {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if gotW := w.String(); gotW != string(b) {
					assert.Equal(t, w.String(), string(b), tt.name)
				}
			}
		})
	}
}
