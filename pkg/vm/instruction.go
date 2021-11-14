package vm

type InsType byte

const (
	Plus     InsType = '+'
	Minus    InsType = '-'
	Right    InsType = '>'
	Left     InsType = '<'
	PutChar  InsType = '.'
	ReadChar InsType = ','
	JZ       InsType = '['
	JNZ      InsType = ']'
)

type Instruction struct {
	Type InsType
	Arg  int
}
