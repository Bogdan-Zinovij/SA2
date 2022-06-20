package lab2

import (
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
	}

	for postfix, expected := range cases {
		res, err := PostfixToPrefix(postfix)
		if assert.Nil(t, err) {
			assert.Equal(t, expected, res)
		} else {
			assert.Equal(t, expected, err)
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
