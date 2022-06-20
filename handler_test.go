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

func TestComputeInvalidSyntaxError(t *testing.T) {
	cases := []string{
		"4 2- 3* 5 +",
		"4 2 & 6 * 9",
		"",
		"err 2 - 3 * 5 +",
	}

	for _, inputExp := range cases {
		input := strings.NewReader(inputExp)
		output := new(bytes.Buffer)
		handler := ComputeHandler{input, output}
		err := handler.Compute()
		assert.NotNil(t, err)
	}
}
