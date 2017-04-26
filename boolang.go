package boolang

import (
	"fmt"
	"strings"
)

var (
	ErrorMismatchedParens = fmt.Errorf("Mismatched parenthesis")
	ErrorSyntax           = fmt.Errorf("Syntax Error")
)

type EvalFunc func(*Leaf, ...interface{}) (bool, error)
type WalkFunc func(*Leaf)

type AST interface {
	Eval(EvalFunc, ...interface{}) (bool, error)
	Walk(WalkFunc)
	String() string
}

type NotAST struct {
	center AST
}

type AndAST struct {
	left  AST
	right AST
}

type OrAST struct {
	left  AST
	right AST
}

type Leaf struct {
	Content  string
	Metadata interface{}
}

func NewLeaf(content string) *Leaf {
	leaf := &Leaf{}
	leaf.Content = content
	return leaf
}

// String
func (t *NotAST) String() string {
	return "Not(" + t.center.String() + ")"
}

func (t *OrAST) String() string {
	return "Or(" + t.left.String() + "," + t.right.String() + ")"
}

func (t *AndAST) String() string {
	return "And(" + t.left.String() + "," + t.right.String() + ")"
}

func (t *Leaf) String() string {
	return "Leaf(" + strings.Trim(t.Content, " ") + ")"
}

// Walk
func (t *NotAST) Walk(fn WalkFunc) {
	t.center.Walk(fn)
}

func (t *AndAST) Walk(fn WalkFunc) {
	t.left.Walk(fn)
	t.right.Walk(fn)
}

func (t *OrAST) Walk(fn WalkFunc) {
	t.left.Walk(fn)
	t.right.Walk(fn)
}

func (t *Leaf) Walk(fn WalkFunc) {
	fn(t)
}

// Eval
func (t *NotAST) Eval(fn EvalFunc, args ...interface{}) (bool, error) {
	b, err := t.center.Eval(fn, args)
	return !b, err
}

func (t *AndAST) Eval(fn EvalFunc, args ...interface{}) (bool, error) {
	left, err := t.left.Eval(fn, args)
	if err != nil {
		return false, err
	}
	if !left {
		return false, nil
	}

	right, err := t.right.Eval(fn, args)
	if err != nil {
		return false, err
	}
	return right, nil
}

func (t *OrAST) Eval(fn EvalFunc, args ...interface{}) (bool, error) {
	left, err := t.left.Eval(fn, args)
	if err != nil {
		return false, err
	}
	if left {
		return true, nil
	}

	right, err := t.right.Eval(fn, args)
	if err != nil {
		return false, err
	}

	return right, nil
}

func (t *Leaf) Eval(fn EvalFunc, args ...interface{}) (bool, error) {
	return fn(t, args)
}
