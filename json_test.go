package boolang

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseJson(t *testing.T) {
	var data struct {
		AST JsonAST `json:"ast"`
	}
	var content = []byte(`{"ast":"A==1 && B == 2"}`)
	err := json.Unmarshal(content, &data)
	require.Nil(t, err)

	counter := 0
	data.AST.Walk(func(l *Leaf) { counter++ })
	require.Equal(t, 2, counter)
}

func TestJsonMarshal(t *testing.T) {
	var data struct {
		AST JsonAST `json:"ast"`
	}
	var content = []byte(`{"ast":"A==1 && B == 2"}`)
	err := json.Unmarshal(content, &data)
	require.Nil(t, err)

	newcontent, err := json.Marshal(data)
	require.NoError(t, err)

	err = json.Unmarshal(newcontent, &data)
	require.NoError(t, err)

	counter := 0
	data.AST.Walk(func(l *Leaf) { counter++ })
	require.Equal(t, 2, counter)
}

func TestParseJsonPtr(t *testing.T) {
	var data struct {
		AST *JsonAST `json:"ast"`
	}
	var content = []byte(`{"ast":"A==1 && B == 2"}`)
	err := json.Unmarshal(content, &data)
	require.Nil(t, err)

	counter := 0
	data.AST.Walk(func(l *Leaf) { counter++ })
	require.Equal(t, 2, counter)
}

func TestParseJsonError(t *testing.T) {
	var data struct {
		AST JsonAST `json:"ast"`
	}
	var content = []byte(`{"ast":"A==1 &&) B == 2"}`)
	err := json.Unmarshal(content, &data)
	require.NotNil(t, err)
}
