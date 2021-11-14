package compiler

import "github.com/TASnomad/bf/pkg/vm"

type Compiler struct {
	code   string
	length int
	pos    int
	instrs []*vm.Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:   code,
		length: len(code),
		pos:    0,
		instrs: []*vm.Instruction{},
	}
}

func (c *Compiler) CompileInstruction(char byte, kind vm.InsType) {
	count := 1

	for c.pos < c.length-1 && c.code[c.pos+1] == char {
		count++
		c.pos++
	}
	c.EmitInstruction(kind, count)
}

func (c *Compiler) EmitInstruction(kind vm.InsType, arg int) int {
	ins := &vm.Instruction{Type: kind, Arg: arg}
	c.instrs = append(c.instrs, ins)
	return len(c.instrs) - 1
}

func (c *Compiler) Compile() []*vm.Instruction {
	loopStack := []int{}

	for c.pos < c.length {
		curr := c.code[c.pos]

		switch curr {
		case '+':
			c.CompileInstruction('+', vm.Plus)
		case '-':
			c.CompileInstruction('-', vm.Minus)
		case '<':
			c.CompileInstruction('<', vm.Left)
		case '>':
			c.CompileInstruction('>', vm.Right)
		case '.':
			c.CompileInstruction('.', vm.PutChar)
		case ',':
			c.CompileInstruction(',', vm.ReadChar)
		case '[':
			insPos := c.EmitInstruction(vm.JZ, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			openIns := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]
			closeIns := c.EmitInstruction(vm.JNZ, openIns)
			c.instrs[openIns].Arg = closeIns
		}
		c.pos++
	}
	return c.instrs
}
