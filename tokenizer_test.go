package boolang_test

import (
	"testing"

	"github.com/apourchet/boolang"
	"github.com/stretchr/testify/require"
)

func TestTokenizer(t *testing.T) {
	table := []struct {
		str    string
		output []string
	}{
		{"||", []string{"||"}},
		{"&&", []string{"&&"}},
		{"!", []string{"!"}},
		{"(", []string{"("}},
		{")", []string{")"}},
		{"asd || qwe", []string{"asd ", "||", " qwe"}},
		{"asd && qwe", []string{"asd ", "&&", " qwe"}},
		{"!asd", []string{"!", "asd"}},
		{"!asd || qwe", []string{"!", "asd ", "||", " qwe"}},
		{"!(asd||qwe)", []string{"!", "(", "asd", "||", "qwe", ")"}},
		{"(asd|qwe||qwe)", []string{"(", "asd|qwe", "||", "qwe", ")"}},
		{"!A == 1 || B == 2 && !C == 3", []string{"!", "A == 1 ", "||", " B == 2 ", "&&", "!", "C == 3"}},
	}

	for _, test := range table {
		tokenizer := boolang.NewTokenizer(test.str)
		tokens := tokenizer.Tokenize()
		require.Equal(t, test.output, tokens)
	}
}
