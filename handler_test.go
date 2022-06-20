package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputWriting(t *testing.T) {
	cases := map[string]string{
		"4 2 - 3 * 5 +":             "+ * - 4 2 3 5",
		"4 2 - 3 * 5 + 5 5 5 - - +": "+ + * - 4 2 3 5 - 5 - 5 5",
	}

	for inputData, outputData := range cases {
		input := strings.NewReader(inputData)
		output := new(bytes.Buffer)
		handler := ComputeHandler{input, output}
		handler.Compute()
		assert.Equal(t, output.String(), outputData)
	}
}
