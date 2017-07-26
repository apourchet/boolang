package boolang

var _ AST = &JsonAST{}

type JsonAST struct{ AST }

func (ast *JsonAST) UnmarshalJSON(b []byte) error {
	tree, err := Parse(string(b))
	if err != nil {
		return err
	}
	*ast = JsonAST{tree}
	return nil
}

func (ast JsonAST) MarshalJSON() ([]byte, error) {
	return []byte(ast.String()), nil
}
