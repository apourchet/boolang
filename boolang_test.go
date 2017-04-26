package boolang_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/apourchet/boolang"
	"github.com/stretchr/testify/assert"
)

func TestLeafOnly(t *testing.T) {
	tree, err := boolang.Parse("A == 1")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Leaf(A == 1)", tree.String())
}

func TestOrOnly(t *testing.T) {
	tree, err := boolang.Parse("A == 1 || B == 2")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Or(Leaf(A == 1),Leaf(B == 2))", tree.String())
}

func TestAndOnly(t *testing.T) {
	tree, err := boolang.Parse("A == 1 && B == 2")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "And(Leaf(A == 1),Leaf(B == 2))", tree.String())
}

func TestAndOr(t *testing.T) {
	tree, err := boolang.Parse("A == 1 && B == 2 || C == 3")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Or(And(Leaf(A == 1),Leaf(B == 2)),Leaf(C == 3))", tree.String())
}

func TestNotOnly(t *testing.T) {
	tree, err := boolang.Parse("!A == 1")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Not(Leaf(A == 1))", tree.String())
}

func TestAndOrNot(t *testing.T) {
	tree, err := boolang.Parse("!A == 1 || B == 2 && !C == 3")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Or(Not(Leaf(A == 1)),And(Leaf(B == 2),Not(Leaf(C == 3))))", tree.String())
}

func TestSimpleParens(t *testing.T) {
	tree, err := boolang.Parse("(A == 1)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Leaf(A == 1)", tree.String())
}

func TestParensInsideOr(t *testing.T) {
	tree, err := boolang.Parse("(A == 1 || B == 2)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Or(Leaf(A == 1),Leaf(B == 2))", tree.String())
}

func TestParensInsideAnd(t *testing.T) {
	tree, err := boolang.Parse("(A == 1 && B == 2)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "And(Leaf(A == 1),Leaf(B == 2))", tree.String())
}

func TestParensInsideAndOr(t *testing.T) {
	tree, err := boolang.Parse("(A == 1 && B == 2 || C == 3)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Or(And(Leaf(A == 1),Leaf(B == 2)),Leaf(C == 3))", tree.String())
}

func TestParensNot(t *testing.T) {
	tree, err := boolang.Parse("!(A == 1 || B == 2)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Not(Or(Leaf(A == 1),Leaf(B == 2)))", tree.String())
}

func TestAll(t *testing.T) {
	tree, err := boolang.Parse("(A == 1 || B == 2) && (C == 3 || D == 4 && !E == 5) || !(F == 6)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "Or(And(Or(Leaf(A == 1),Leaf(B == 2)),Or(Leaf(C == 3),And(Leaf(D == 4),Not(Leaf(E == 5))))),Not(Leaf(F == 6)))", tree.String())
}

func TestWalk(t *testing.T) {
	tree, err := boolang.Parse("(A == 1 || B == 2) && (C == 3 || D == 4 && !E == 5) || !(F == 6)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)

	counter := 0
	fn := func(l *boolang.Leaf) { counter += 1 }
	tree.Walk(fn)

	assert.Equal(t, 6, counter)
}

func TestEvalCount(t *testing.T) {
	tree, err := boolang.Parse("(A == 1 && B == 2 || C == 3)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)

	counter := 0
	fn := func(l *boolang.Leaf, _ ...interface{}) (bool, error) {
		counter += 1
		return true, nil
	}
	tree.Eval(fn)

	assert.Equal(t, 2, counter)
}

func TestEval(t *testing.T) {
	tree, err := boolang.Parse("(false || true) && (true || error && error) || !(error)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)

	fn := func(l *boolang.Leaf, _ ...interface{}) (bool, error) {
		content := strings.Trim(l.Content, " ")
		if content == "false" {
			return false, nil
		} else if content == "true" {
			return true, nil
		}
		return false, fmt.Errorf("Should not have evaluated: %s", content)
	}
	val, err := tree.Eval(fn)
	assert.Nil(t, err)
	assert.True(t, val)
}

func TestEvalError(t *testing.T) {
	tree, err := boolang.Parse("(false || true) && (error || error && error) || !(error)")
	assert.Nil(t, err)
	assert.NotNil(t, tree)

	fn := func(l *boolang.Leaf, _ ...interface{}) (bool, error) {
		content := strings.Trim(l.Content, " ")
		if content == "false" {
			return false, nil
		} else if content == "true" {
			return true, nil
		}
		return false, fmt.Errorf("Should not have evaluated: %s", content)
	}

	val, err := tree.Eval(fn)
	assert.NotNil(t, err)
	assert.False(t, val)
}
