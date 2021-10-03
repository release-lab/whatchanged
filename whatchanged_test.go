package whatchanged

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whatchanged-community/whatchanged/option"
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
		{
			name:       "whatchanged version range",
			Project:    "./",
			ResultFile: "./__test__/whatchanged-[23448a5482359f28a0089b17280dd2a0a0eaef26~9dff4fc6a9d746ffd9dd10215cf04d2fec2edd2a].CHANGELOG.md",
			options: &option.Options{
				Version: []string{"23448a5482359f28a0089b17280dd2a0a0eaef26~9dff4fc6a9d746ffd9dd10215cf04d2fec2edd2a"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := Generate(context.Background(), tt.Project, w, tt.options); err != tt.wantErr {
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
