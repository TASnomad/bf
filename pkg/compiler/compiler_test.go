package compiler_test

import (
	"testing"

	"github.com/TASnomad/bf/pkg/compiler"
	"github.com/TASnomad/bf/pkg/vm"
)

func TestCompiler(t *testing.T) {
	code := `
	+++++
	-----
	+++++
	>>>>>
	<<<<<
	`

	expected := []*vm.Instruction{
		{Type: vm.Plus, Arg: 5},
		{Type: vm.Minus, Arg: 5},
		{Type: vm.Plus, Arg: 5},
		{Type: vm.Right, Arg: 5},
		{Type: vm.Left, Arg: 5},
	}

	c := compiler.NewCompiler(code)
	bc := c.Compile()

	if len(bc) != len(expected) {
		t.Fatalf("Bytecode length is mismatching code length: expected=%+v, got=%+v", len(expected), len(code))
	}

	for i, op := range expected {
		if *bc[i] != *op {
			t.Errorf("Unexpecting instruction: expected=%+v, got=%+v", op, bc[i])
		}
	}
}
