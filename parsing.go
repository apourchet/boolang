package boolang

import "strings"

func Parse(program string) (AST, error) {
	program = strings.Replace(program, "&&", "&", -1)
	program = strings.Replace(program, "||", "|", -1)
	valid := checkParens(program)
	if !valid {
		return nil, ErrorMismatchedParens
	}
	tokens := tokenize(program)
	tree, err := buildTree(tokens)
	return tree, err
}

func tokenize(program string) []string {
	tokens := []string{}
	current_token := ""
	for index := 0; index < len(program); index++ {
		char := program[index]
		if char == '(' || char == '!' || char == '&' || char == '|' {
			tokens = append(tokens, current_token)
			current_token = ""
		}
		if char == '(' {
			end := matchingParens(program[index:]) + index
			inners := tokenize(program[index+1 : end])
			tokens = append(tokens, "(")
			tokens = append(tokens, inners...)
			tokens = append(tokens, ")")
			index = end
		} else if char == '&' {
			tokens = append(tokens, "&")
		} else if char == '|' {
			tokens = append(tokens, "|")
		} else if char == '!' {
			tokens = append(tokens, "!")
		} else {
			current_token += string(char)
		}
	}

	// Remove empty tokens
	tokens = append(tokens, current_token)
	ret := []string{}
	for _, token := range tokens {
		if strings.Trim(token, " ") != "" {
			ret = append(ret, token)
		}
	}

	return ret
}

func buildTree(tokens []string) (AST, error) {
	if len(tokens) == 0 {
		return nil, ErrorSyntax
	} else if len(tokens) == 1 {
		return NewLeaf(tokens[0]), nil
	}

	orIndex := findToken(tokens, "|")
	if orIndex == 0 || orIndex == len(tokens)-1 {
		return nil, ErrorSyntax
	} else if orIndex > 0 {
		left, err := buildTree(tokens[0:orIndex])
		if err != nil {
			return nil, err
		}
		right, err := buildTree(tokens[orIndex+1:])
		if err != nil {
			return nil, err
		}
		return &OrAST{left, right}, nil
	}

	andIndex := findToken(tokens, "&")
	if andIndex == 0 || andIndex == len(tokens)-1 {
		return nil, ErrorSyntax
	} else if andIndex > 0 {
		left, err := buildTree(tokens[0:andIndex])
		if err != nil {
			return nil, err
		}
		right, err := buildTree(tokens[andIndex+1:])
		if err != nil {
			return nil, err
		}
		return &AndAST{left, right}, nil
	}

	notIndex := findToken(tokens, "!")
	if notIndex > 0 {
		return nil, ErrorSyntax
	} else if notIndex == 0 {
		center, err := buildTree(tokens[1:])
		if err != nil {
			return nil, err
		}
		return &NotAST{center}, nil
	}

	if tokens[0] != "(" || tokens[len(tokens)-1] != ")" {
		return nil, ErrorSyntax
	} else {
		return buildTree(tokens[1 : len(tokens)-1])
	}

	return nil, ErrorSyntax
}

func findToken(tokens []string, token string) int {
	level := 0
	for index, current := range tokens {
		if current == "(" {
			level += 1
		} else if current == ")" {
			level -= 1
		} else if level == 0 {
			if current == token {
				return index
			}
		}
	}
	return -1
}

func matchingParens(program string) int {
	level := 0
	for i, char := range program {
		if char == '(' {
			level += 1
		} else if char == ')' {
			level -= 1
		}
		if level == 0 {
			return i
		}
	}
	return -1
}

func checkParens(program string) bool {
	level := 0
	for _, char := range program {
		if char == '(' {
			level += 1
		} else if char == ')' {
			level -= 1
		}
		if level < 0 {
			return false
		}
	}
	return level == 0
}
