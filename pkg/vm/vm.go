package vm

import "io"

type Vm struct {
	code   []*Instruction
	ip     int
	mem    []int
	dp     int
	input  io.Reader
	output io.Writer
	buf    []byte
}

func NewVm(code []*Instruction, in io.Reader, out io.Writer) *Vm {
	return &Vm{
		code:   code,
		ip:     0,
		mem:    make([]int, 30000),
		dp:     0,
		input:  in,
		output: out,
		buf:    make([]byte, 1),
	}
}

func (v Vm) Memory() []int {
	return v.mem
}

func (v *Vm) SetMemory(idx int, b int) {
	v.mem[idx] = b
}

func (v *Vm) ReadChar() {
	n, err := v.input.Read(v.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("R_ERR")
	}
	v.mem[v.dp] = int(v.buf[0])
}

func (v *Vm) WriteChar() {
	v.buf[0] = byte(v.mem[v.dp])

	n, err := v.output.Write(v.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("W_ERR")
	}
}

func (v *Vm) Execute() {
	for v.ip < len(v.code) {
		ins := v.code[v.ip]

		switch ins.Type {
		case Plus:
			v.mem[v.dp] += ins.Arg
		case Minus:
			v.mem[v.dp] -= ins.Arg
		case Right:
			v.dp += ins.Arg
		case Left:
			v.dp -= ins.Arg
		case ReadChar:
			for i := 0; i < ins.Arg; i++ {
				v.ReadChar()
			}
		case PutChar:
			for i := 0; i < ins.Arg; i++ {
				v.WriteChar()
			}

		case JZ:
			if v.mem[v.dp] == 0 {
				v.ip = ins.Arg
				continue
			}
		case JNZ:
			if v.mem[v.dp] != 0 {
				v.ip = ins.Arg
				continue
			}
		}
		v.ip++
	}
}
