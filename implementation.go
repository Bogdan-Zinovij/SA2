package lab2

import (
	"fmt"
	"strings"
)

// TODO: document this function.
// Postfix to Prefix converts
func PostfixToPrefix(inputExpression string) (string, error) {
	symbols := strings.Split(inputExpression, " ")
	signs := "-+/*^"
	for i := 0; i < len(symbols); i++ {
		if strings.Contains(signs, symbols[i]) {
			symbols[i-2] = symbols[i] + " " + symbols[i-2] + " " + symbols[i-1]
			symbols = append(symbols[:i-1], symbols[i+1:]...)
			i = i - 2
		}
	}

	return symbols[0], fmt.Errorf("TODO")
}
