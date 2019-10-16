package ast // import "github.com/andrewarchi/nebula/ast"

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/andrewarchi/nebula/bigint"
	"github.com/andrewarchi/nebula/token"
)

func TestJoinSimpleCalls(t *testing.T) {
	tokens := []token.Token{
		{Type: token.Push, Arg: big.NewInt(1)},  // 0
		{Type: token.Add},                       // 1
		{Type: token.Mul},                       // 2
		{Type: token.Label, Arg: big.NewInt(1)}, // 3
		{Type: token.Copy, Arg: big.NewInt(5)},  // 4
		{Type: token.Mod},                       // 5
		{Type: token.Slide, Arg: big.NewInt(2)}, // 6
	}

	ast, err := Parse(tokens, nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	v1 := Val(&ConstVal{big.NewInt(1)})
	s0 := Val(&StackVal{0})
	s1 := Val(&StackVal{1})
	s2 := Val(&StackVal{2})
	sn1 := Val(&StackVal{-1})
	sn2 := Val(&StackVal{-2})
	sn7 := Val(&StackVal{-7})

	var stack Stack
	stack.Push(&v1) // 0
	stack.Pop()     // 1
	stack.Pop()     // 1
	stack.Push(&s0) // 1
	stack.Pop()     // 2
	stack.Pop()     // 2
	stack.Push(&s1) // 2
	stack.Copy(5)   // 4
	stack.Pop()     // 5
	stack.Pop()     // 5
	stack.Push(&s2) // 5
	stack.Slide(2)  // 6

	constVals := bigint.NewMap(nil)
	constVals.Put(big.NewInt(1), &v1)

	blockJoined := &BasicBlock{
		Stack: stack,
		Nodes: []Node{
			&AssignStmt{Assign: &s0, Expr: &ArithExpr{Op: token.Add, LHS: &sn1, RHS: &v1}},
			&AssignStmt{Assign: &s1, Expr: &ArithExpr{Op: token.Mul, LHS: &sn2, RHS: &s0}},
			&AssignStmt{Assign: &s2, Expr: &ArithExpr{Op: token.Mod, LHS: &s1, RHS: &sn7}},
		},
		Terminator: &EndStmt{},
		Entries:    []*BasicBlock{entryBlock},
		Callers:    []*BasicBlock{entryBlock},
	}
	astJoined := &AST{
		Blocks:      []*BasicBlock{blockJoined},
		Entry:       blockJoined,
		ConstVals:   *constVals,
		NextBlockID: 2,
		NextStackID: 3,
	}

	ast.JoinSimpleCalls()
	if !reflect.DeepEqual(ast, astJoined) {
		t.Errorf("join not equal\ngot:\n%v\nwant:\n%v", ast, astJoined)
	}
}
