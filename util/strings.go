package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func UpTo(s string, numChars int) string {
	if len(s) > numChars {
		return s[:numChars]
	}
	return s
}

func UpToR(s string, numChars int) string {
	if len(s) > numChars {
		return s[len(s)-numChars:]
	}
	return s
}

func IndentedDump(v interface{}) string {

	firstColLeftMostPrefix := " "
	byts, err := json.MarshalIndent(v, firstColLeftMostPrefix, "\t")
	if err != nil {
		s := fmt.Sprintf("error indent: %v\n", err)
		return s
	}

	byts = bytes.Replace(byts, []byte(`\u003c`), []byte("<"), -1)
	byts = bytes.Replace(byts, []byte(`\u003e`), []byte(">"), -1)
	byts = bytes.Replace(byts, []byte(`\n`), []byte("\n"), -1)

	return string(byts)
}
