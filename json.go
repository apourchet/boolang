package boolang

import "encoding/json"

var _ AST = &JsonAST{}

type JsonAST struct{ AST }

func (ast *JsonAST) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tree, err := Parse(s)
	if err != nil {
		return err
	}
	*ast = JsonAST{tree}
	return nil
}

func (ast JsonAST) MarshalJSON() ([]byte, error) {
	return []byte(ast.String()), nil
}
