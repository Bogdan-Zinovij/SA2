package lab2

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	cases := map[string]string{
		"1 2 +":                                 "+ 1 2",
		"1 2 3 * +":                             "+ 1 * 2 3",
		"4 2 / 5 4 * 3 + /":                     "/ / 4 2 + * 5 4 3",
		"3 4 + 1 - 6 5 - 2.0 8 7 + * / ^":       "^ - + 3 4 1 / - 6 5 * 2.0 + 8 7",
		"1 2 3 + + 4 - 5 + 8 ^ 3 7 * ^ 2 2 + -": "- ^ ^ + - + 1 + 2 3 4 5 8 * 3 7 + 2 2",
		"": "validation error: expression is not specified",
		"w r o n g s y m b o l s": "validation error: invalid symbol specified",
		"7 23 84 / g o + + -": "validation error: invalid symbol specified",
		"7 23 15": "validation error: invalid expression specified",
		"+ - / * ^": "validation error: invalid expression specified",
		"12 42 + 17 - *": "validation error: invalid expression specified",
	}

	for postfix, expected := range cases {
		res, err := PostfixToPrefix(postfix)
		if  err == nil {
			assert.Equal(t, expected, res)
		} else {
			assert.Equal(t, errors.New(expected), err)
		}
	}
}

func ExamplePostfixToPrefix() {
	res, err := PostfixToPrefix("1 3 + 5 * 2 4 * -")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
