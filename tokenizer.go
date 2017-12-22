package boolang

import "io"

type Tokenizer struct {
	program string
	index   int
}

func NewTokenizer(program string) *Tokenizer {
	return &Tokenizer{
		program: program,
		index:   0,
	}
}

func (t *Tokenizer) Tokenize() []string {
	tokens := []string{}
	for token, err := t.next(); err == nil; token, err = t.next() {
		if token == " " {
			continue
		}
		tokens = append(tokens, token)
	}
	return tokens
}

func (t *Tokenizer) next() (string, error) {
	token := ""
	lastchar := ""
	for t.peek() != "" {
		char := t.get()
		if char == "(" || char == ")" {
			if token == "" {
				return char, nil
			}
			t.back(1)
			return token, nil
		} else if char == "&" {
			if lastchar == "&" && token != "&" {
				t.back(2)
				return token[:len(token)-1], nil
			} else if lastchar == "&" {
				return "&&", nil
			}
		} else if char == "|" {
			if lastchar == "|" && token != "|" {
				t.back(2)
				return token[:len(token)-1], nil
			} else if lastchar == "|" {
				return "||", nil
			}
		} else if char == "!" {
			if token == "" {
				return "!", nil
			}
			t.back(1)
			return token, nil
		}
		token += char
		lastchar = char
	}

	if token == "" {
		return "", io.EOF
	}
	return token, nil
}

func (t *Tokenizer) peek() string {
	if t.index >= len(t.program) {
		return ""
	}
	return string(t.program[t.index])
}

func (t *Tokenizer) get() string {
	if t.index >= len(t.program) {
		return ""
	}
	ch := t.program[t.index]
	t.index++
	return string(ch)
}

func (t *Tokenizer) back(n int) {
	t.index -= n
}
