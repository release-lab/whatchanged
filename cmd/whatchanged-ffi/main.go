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

	"github.com/release-lab/whatchanged"
)

type FFIResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

//export Generate
func Generate(repo *C.char, optionStr *C.char) *C.char {
	var (
		result FFIResult
		option = whatchanged.Options{}
	)

	if err := json.Unmarshal([]byte(C.GoString(optionStr)), &option); err != nil {
		result.Error = fmt.Sprintf("%+v\n", err)
		b, _ := json.Marshal(result)

		return C.CString(string(b))
	}

	output := bytes.NewBuffer([]byte{})

	err := whatchanged.Generate(context.Background(), C.GoString(repo), output, &option)

	if err != nil {
		result.Error = fmt.Sprintf("%+v\n", err)
	} else {
		result.Result = output.String()
	}

	b, _ := json.Marshal(result)

	return C.CString(string(b))
}

func main() {}
