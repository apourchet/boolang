package boolang_test

import (
	"encoding/json"
	"testing"

	"github.com/apourchet/boolang"
	"github.com/stretchr/testify/require"
)

func TestParseJson(t *testing.T) {
	var data struct {
		AST boolang.JsonAST `json:"ast"`
	}
	var content = []byte(`{"ast":"A==1 && B == 2"}`)
	err := json.Unmarshal(content, &data)
	require.Nil(t, err)

	counter := 0
	data.AST.Walk(func(l *boolang.Leaf) { counter += 1 })
	require.Equal(t, 2, counter)
}

func TestParseJsonPtr(t *testing.T) {
	var data struct {
		AST *boolang.JsonAST `json:"ast"`
	}
	var content = []byte(`{"ast":"A==1 && B == 2"}`)
	err := json.Unmarshal(content, &data)
	require.Nil(t, err)

	counter := 0
	data.AST.Walk(func(l *boolang.Leaf) { counter += 1 })
	require.Equal(t, 2, counter)
}

func TestParseJsonError(t *testing.T) {
	var data struct {
		AST boolang.JsonAST `json:"ast"`
	}
	var content = []byte(`{"ast":"A==1 &&) B == 2"}`)
	err := json.Unmarshal(content, &data)
	require.NotNil(t, err)
}
