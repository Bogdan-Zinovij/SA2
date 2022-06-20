package lab2

import (
	// "fmt"
	"io"
	"strings"
)

func processInputStr(inputStr string) string {
	processedStr := strings.Trim(inputStr, "\x00")
	processedStr = strings.TrimSpace(processedStr)
	processedStr = strings.Trim(processedStr, "\n")
	return processedStr
}

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var slice = make([]byte, 64)
	_, readErr := ch.Input.Read(slice)

	inputStr := string(slice)
	inputStr = processInputStr(inputStr)

	if readErr != nil {
		return readErr
	}

	res, convertErr := PostfixToPrefix(inputStr)
	if convertErr != nil {
		return convertErr
	}

	_, writeErr := ch.Output.Write([]byte(res))
	if writeErr != nil {
		return writeErr
	}

	return nil
}
