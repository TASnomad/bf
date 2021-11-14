package vm_test

import (
	"bytes"
	"testing"

	"github.com/TASnomad/bf/pkg/compiler"
	"github.com/TASnomad/bf/pkg/vm"
)

func TestReadChar(t *testing.T) {
	stdin := bytes.NewBufferString("ABCDEF")
	stdout := new(bytes.Buffer)

	c := compiler.NewCompiler(",>,>,>,>,>,>")
	instrs := c.Compile()

	m := vm.NewVm(instrs, stdin, stdout)
	m.Execute()

	expectedMem := []int{
		int('A'),
		int('B'),
		int('C'),
		int('D'),
		int('E'),
		int('F'),
	}

	mem := m.Memory()
	for i, expected := range expectedMem {
		if mem[i] != expected {
			t.Errorf("Wrong value in Memory[%d]: expecting=%d, got=%d\n", i, expected, mem[i])
		}
	}
}

func TestPutChar(t *testing.T) {
	stdin := bytes.NewBufferString("")
	stdout := new(bytes.Buffer)

	c := compiler.NewCompiler(".>.>.>.>.>.>")
	instrs := c.Compile()

	m := vm.NewVm(instrs, stdin, stdout)

	expectedMem := []int{
		int('A'),
		int('B'),
		int('C'),
		int('D'),
		int('E'),
		int('F'),
	}

	for i, b := range expectedMem {
		m.SetMemory(i, b)
	}

	m.Execute()
	out := stdout.String()
	if out != "ABCDEF" {
		t.Errorf("Wroung output: got=%q", out)
	}
}

func TestHelloWorld(t *testing.T) {
	code := `++++++++[>++++[>++>+++>+++>+<<<<-]>+> +>->>+[<]<-]>>.>---.+++++++ ..+ ++.>>.<-.<.+++.------.--------.>>+.>++.`

	stdin := bytes.NewBufferString("")
	stdout := new(bytes.Buffer)

	c := compiler.NewCompiler(code)
	instrs := c.Compile()

	m := vm.NewVm(instrs, stdin, stdout)
	m.Execute()
	out := stdout.String()
	if out != "Hello World!\n" {
		t.Errorf("Wroung output: got=%q", out)
	}
}
