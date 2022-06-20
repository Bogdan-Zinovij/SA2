package lab2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const signs = "-+/*^"

func validateInputExpression(inputExp []string) error {

	var signsCounter, numbersCounter int

	if inputExp[0] == "" {
		return errors.New("expression is not specified")
	}

	for _, symbol := range inputExp {
		_, err := strconv.ParseFloat(symbol, 64) // check if symbol is a number
		if err == nil {
			numbersCounter++
			continue
		}

		if strings.Contains(signs, symbol) { // check if symbol is a sign
			signsCounter++
			continue
		}

		return errors.New("invalid symbol specified")
	}

	if numbersCounter != signsCounter+1 {
		return errors.New("invalid expression specified")
	}

	return nil
}

// Postfix to Prefix converts
func PostfixToPrefix(inputExpression string) (string, error) {
	symbols := strings.Split(inputExpression, " ")
	validationErr := validateInputExpression(symbols)
	if validationErr != nil {
		return "", fmt.Errorf("validation error: %v", validationErr)
	}
	for i := 0; i < len(symbols); i++ {
		if strings.Contains(signs, symbols[i]) {
			symbols[i-2] = symbols[i] + " " + symbols[i-2] + " " + symbols[i-1]
			symbols = append(symbols[:i-1], symbols[i+1:]...)
			i = i - 2
		}
	}

	return symbols[0], nil
}
