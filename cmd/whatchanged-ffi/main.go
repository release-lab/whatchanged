// filename: whatchanged.so
// filename: whatchanged.a
// filename: whatchanged.dll
// filename: whatchanged.dylib
package main

import "C"

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/release-lab/whatchanged"
)

//export Generate
func Generate(repo *C.char, version *C.char) *C.char {
	type Result struct {
		Error  string `json:"error"`
		Result string `json:"result"`
	}

	var result Result

	output := bytes.NewBuffer([]byte{})

	err := whatchanged.Generate(context.Background(), C.GoString(repo), output, &whatchanged.Options{
		Version: regexp.MustCompile(`\s+`).Split(C.GoString(version), -1),
		Branch:  "master",
		Preset:  whatchanged.EnumPreset("full"),
	})

	if err != nil {
		result.Error = fmt.Sprintf("%+v\n", err)
	} else {
		result.Result = output.String()
	}

	b, _ := json.Marshal(result)

	return C.CString(string(b))
}

func main() {}
