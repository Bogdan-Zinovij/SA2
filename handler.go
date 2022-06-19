package lab2

import (
	// "fmt"
	"io"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	input  io.Reader
	output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var slice = make([]byte, 64)
	_, readErr := ch.input.Read(slice)

	str := strings.Trim(string(slice), "\x00")

	if readErr != nil {
		return readErr
	}

	res, convertErr := PostfixToPrefix(str)
	if convertErr != nil {
		return convertErr
	}

	_, writeErr := ch.output.Write([]byte(res))
	if writeErr != nil {
		return writeErr
	}

	return nil
}
