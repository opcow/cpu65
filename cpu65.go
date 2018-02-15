// Package cpu65 a 6502 disassembler/assembler
package cpu65

// MaxMem is the maximum memory size for or CPU
const (
	MaxMem = 0x10000
	Modes  = Imp + 1
)

// CPU modes
const (
	unk = iota
	Imm // immediate mode
	Zpa // zero page
	Zpx // zero page x
	Zpy // zero page y
	Abs // absolute
	Abx // absolute x
	Aby // absolute y
	Ind // indirect
	Inx // indirect x
	Iny // indirect y
	Acc // accumulator
	Rel // relative
	Imp // implied
)

type Opcode struct {
	code     int
	length   int
	mode     int
	mnemonic string
}

var opcodes = [256]Opcode{
	0x69: Opcode{0x69, 2, Imm, "ADC"},
	0x65: Opcode{0x65, 2, Zpa, "ADC"},
	0x75: Opcode{0x75, 2, Zpx, "ADC"},
	0x6d: Opcode{0x6d, 3, Abs, "ADC"},
	0x7d: Opcode{0x7d, 3, Abx, "ADC"},
	0x79: Opcode{0x79, 3, Aby, "ADC"},
	0x61: Opcode{0x61, 2, Inx, "ADC"},
	0x71: Opcode{0x71, 2, Iny, "ADC"},
	0x29: Opcode{0x29, 2, Imm, "AND"},
	0x25: Opcode{0x25, 2, Zpa, "AND"},
	0x35: Opcode{0x35, 2, Zpx, "AND"},
	0x2d: Opcode{0x2d, 3, Abs, "AND"},
	0x3d: Opcode{0x3d, 3, Abx, "AND"},
	0x39: Opcode{0x39, 3, Aby, "AND"},
	0x21: Opcode{0x21, 2, Inx, "AND"},
	0x31: Opcode{0x31, 2, Iny, "AND"},
	0x0a: Opcode{0x0a, 1, Acc, "ASL"},
	0x06: Opcode{0x06, 2, Zpa, "ASL"},
	0x16: Opcode{0x16, 2, Zpx, "ASL"},
	0x0e: Opcode{0x0e, 3, Abs, "ASL"},
	0x1e: Opcode{0x1e, 3, Abx, "ASL"},
	0x24: Opcode{0x24, 2, Zpa, "BIT"},
	0x2c: Opcode{0x2c, 3, Abs, "BIT"},
	0x10: Opcode{0x10, 2, Rel, "BPL"},
	0x30: Opcode{0x30, 2, Rel, "BMI"},
	0x50: Opcode{0x50, 2, Rel, "BVC"},
	0x70: Opcode{0x70, 2, Rel, "BVS"},
	0x90: Opcode{0x90, 2, Rel, "BCC"},
	0xb0: Opcode{0xb0, 2, Rel, "BCS"},
	0xd0: Opcode{0xd0, 2, Rel, "BNE"},
	0xf0: Opcode{0xf0, 2, Rel, "BEQ"},
	0x00: Opcode{0x00, 1, Imp, "BRK"},
	0xc9: Opcode{0xc9, 2, Imm, "CMP"},
	0xc5: Opcode{0xc5, 2, Zpa, "CMP"},
	0xd5: Opcode{0xd5, 2, Zpx, "CMP"},
	0xcd: Opcode{0xcd, 3, Abs, "CMP"},
	0xdd: Opcode{0xdd, 3, Abx, "CMP"},
	0xd9: Opcode{0xd9, 3, Aby, "CMP"},
	0xc1: Opcode{0xc1, 2, Inx, "CMP"},
	0xd1: Opcode{0xd1, 2, Iny, "CMP"},
	0xe0: Opcode{0xe0, 2, Imm, "CPX"},
	0xe4: Opcode{0xe4, 2, Zpa, "CPX"},
	0xec: Opcode{0xec, 3, Abs, "CPX"},
	0xc0: Opcode{0xc0, 2, Imm, "CPY"},
	0xc4: Opcode{0xc4, 2, Zpa, "CPY"},
	0xcc: Opcode{0xcc, 3, Abs, "CPY"},
	0xc6: Opcode{0xc6, 2, Zpa, "DEC"},
	0xd6: Opcode{0xd6, 2, Zpx, "DEC"},
	0xce: Opcode{0xce, 3, Abs, "DEC"},
	0xde: Opcode{0xde, 3, Abx, "DEC"},
	0x49: Opcode{0x49, 2, Imm, "EOR"},
	0x45: Opcode{0x45, 2, Zpa, "EOR"},
	0x55: Opcode{0x55, 2, Zpx, "EOR"},
	0x4d: Opcode{0x4d, 3, Abs, "EOR"},
	0x5d: Opcode{0x5d, 3, Abx, "EOR"},
	0x59: Opcode{0x59, 3, Aby, "EOR"},
	0x41: Opcode{0x41, 2, Inx, "EOR"},
	0x51: Opcode{0x51, 2, Iny, "EOR"},
	0x18: Opcode{0x18, 1, Imp, "CLC"},
	0x38: Opcode{0x38, 1, Imp, "SEC"},
	0x58: Opcode{0x58, 1, Imp, "CLI"},
	0x78: Opcode{0x78, 1, Imp, "SEI"},
	0xb8: Opcode{0xb8, 1, Imp, "CLV"},
	0xd8: Opcode{0xd8, 1, Imp, "CLD"},
	0xf8: Opcode{0xf8, 1, Imp, "SED"},
	0xe6: Opcode{0xe6, 2, Zpa, "INC"},
	0xf6: Opcode{0xf6, 2, Zpx, "INC"},
	0xee: Opcode{0xee, 3, Abs, "INC"},
	0xfe: Opcode{0xfe, 3, Abx, "INC"},
	0x4c: Opcode{0x4c, 3, Abs, "JMP"},
	0x6c: Opcode{0x6c, 3, Ind, "JMP"},
	0x20: Opcode{0x20, 3, Abs, "JSR"},
	0xa9: Opcode{0xa9, 2, Imm, "LDA"},
	0xa5: Opcode{0xa5, 2, Zpa, "LDA"},
	0xb5: Opcode{0xb5, 2, Zpx, "LDA"},
	0xad: Opcode{0xad, 3, Abs, "LDA"},
	0xbd: Opcode{0xbd, 3, Abx, "LDA"},
	0xb9: Opcode{0xb9, 3, Aby, "LDA"},
	0xa1: Opcode{0xa1, 2, Inx, "LDA"},
	0xb1: Opcode{0xb1, 2, Iny, "LDA"},
	0xa2: Opcode{0xa2, 2, Imm, "LDX"},
	0xa6: Opcode{0xa6, 2, Zpa, "LDX"},
	0xb6: Opcode{0xb6, 2, Zpy, "LDX"},
	0xae: Opcode{0xae, 3, Abs, "LDX"},
	0xbe: Opcode{0xbe, 3, Aby, "LDX"},
	0xa0: Opcode{0xa0, 2, Imm, "LDY"},
	0xa4: Opcode{0xa4, 2, Zpa, "LDY"},
	0xb4: Opcode{0xb4, 2, Zpx, "LDY"},
	0xac: Opcode{0xac, 3, Abs, "LDY"},
	0xbc: Opcode{0xbc, 3, Abx, "LDY"},
	0x4a: Opcode{0x4a, 1, Acc, "LSR"},
	0x46: Opcode{0x46, 2, Zpa, "LSR"},
	0x56: Opcode{0x56, 2, Zpx, "LSR"},
	0x4e: Opcode{0x4e, 3, Abs, "LSR"},
	0x5e: Opcode{0x5e, 3, Abx, "LSR"},
	0xea: Opcode{0xea, 1, Imp, "NOP"},
	0x09: Opcode{0x09, 2, Imm, "ORA"},
	0x05: Opcode{0x05, 2, Zpa, "ORA"},
	0x15: Opcode{0x15, 2, Zpx, "ORA"},
	0x0d: Opcode{0x0d, 3, Abs, "ORA"},
	0x1d: Opcode{0x1d, 3, Abx, "ORA"},
	0x19: Opcode{0x19, 3, Aby, "ORA"},
	0x01: Opcode{0x01, 2, Inx, "ORA"},
	0x11: Opcode{0x11, 2, Iny, "ORA"},
	0xaa: Opcode{0xaa, 1, Imp, "TAX"},
	0x8a: Opcode{0x8a, 1, Imp, "TXA"},
	0xca: Opcode{0xca, 1, Imp, "DEX"},
	0xe8: Opcode{0xe8, 1, Imp, "Inx"},
	0xa8: Opcode{0xa8, 1, Imp, "TAY"},
	0x98: Opcode{0x98, 1, Imp, "TYA"},
	0x88: Opcode{0x88, 1, Imp, "DEY"},
	0xc8: Opcode{0xc8, 1, Imp, "Iny"},
	0x2a: Opcode{0x2a, 1, Acc, "ROL"},
	0x26: Opcode{0x26, 2, Zpa, "ROL"},
	0x36: Opcode{0x36, 2, Zpx, "ROL"},
	0x2e: Opcode{0x2e, 3, Abs, "ROL"},
	0x3e: Opcode{0x3e, 3, Abx, "ROL"},
	0x6a: Opcode{0x6a, 1, Acc, "ROR"},
	0x66: Opcode{0x66, 2, Zpa, "ROR"},
	0x76: Opcode{0x76, 2, Zpx, "ROR"},
	0x6e: Opcode{0x6e, 3, Abs, "ROR"},
	0x7e: Opcode{0x7e, 3, Abx, "ROR"},
	0x40: Opcode{0x40, 1, Imp, "RTI"},
	0x60: Opcode{0x60, 1, Imp, "RTS"},
	0xe9: Opcode{0xe9, 2, Imm, "SBC"},
	0xe5: Opcode{0xe5, 2, Zpa, "SBC"},
	0xf5: Opcode{0xf5, 2, Zpx, "SBC"},
	0xed: Opcode{0xed, 3, Abs, "SBC"},
	0xfd: Opcode{0xfd, 3, Abx, "SBC"},
	0xf9: Opcode{0xf9, 3, Aby, "SBC"},
	0xe1: Opcode{0xe1, 2, Inx, "SBC"},
	0xf1: Opcode{0xf1, 2, Iny, "SBC"},
	0x85: Opcode{0x85, 2, Zpa, "STA"},
	0x95: Opcode{0x95, 2, Zpx, "STA"},
	0x8d: Opcode{0x8d, 3, Abs, "STA"},
	0x9d: Opcode{0x9d, 3, Abx, "STA"},
	0x99: Opcode{0x99, 3, Aby, "STA"},
	0x81: Opcode{0x81, 2, Inx, "STA"},
	0x91: Opcode{0x91, 2, Iny, "STA"},
	0x9a: Opcode{0x9a, 1, Imp, "TXS"},
	0xba: Opcode{0xba, 1, Imp, "TSX"},
	0x48: Opcode{0x48, 1, Imp, "PHA"},
	0x68: Opcode{0x68, 1, Imp, "PLA"},
	0x08: Opcode{0x08, 1, Imp, "PHP"},
	0x28: Opcode{0x28, 1, Imp, "PLP"},
	0x86: Opcode{0x86, 2, Zpa, "STX"},
	0x96: Opcode{0x96, 2, Zpy, "STX"},
	0x8e: Opcode{0x8e, 3, Abs, "STX"},
	0x84: Opcode{0x84, 2, Zpa, "STY"},
	0x94: Opcode{0x94, 2, Zpx, "STY"},
	0x8c: Opcode{0x8c, 3, Abs, "STY"},
}

// the status registers
const (
	StatusC = byte(1 << 0) // carry
	StatusZ = byte(1 << 1) // zero
	StatusI = byte(1 << 2) // interrupt
	StatusD = byte(1 << 3) // decimal
	StatusB = byte(1 << 4) // break
	StatusU = byte(1 << 5) // unused
	StatusV = byte(1 << 6) // overflow
	StatusN = byte(1 << 7) // negative
)

type instruction struct {
	code     int
	length   int
	mode     int
	ops      []byte
	mnemonic string
}

// CPU virtual processor + memory type
// call AttachMem() before use
type CPU struct {
	a      byte
	x      byte
	y      byte
	status byte
	sp     byte
	pc     int
	ins    instruction
	mem    *[0x10000]byte
	stack  []byte
}

// AttachMem attaches memory to the CPU
func (c *CPU) AttachMem(m *[MaxMem]byte) {
	c.mem = m
	c.stack = c.mem[0x100:0x200]
}

// Next sets the program counter to the next instruction
func (c *CPU) Next() int {
	c.pc += c.ins.length
	return c.pc
}

// SetPC sets the program counter
func (c *CPU) SetPC(pc int) {
	c.pc = pc
}

// PC gets the program counter
func (c *CPU) PC() int {
	return c.pc
}

// InsLen returns the length of the instruction at the PC
func (c *CPU) InsLen() int {
	return c.ins.length
}

// Mode returns the mode of the instruction at the PC
func (c *CPU) Mode() int {
	return c.ins.mode
}

// Opcode returns the Opcode at the PC
func (c *CPU) Opcode() int {
	return c.ins.code
}

// Mnemonic returns the mnemonic of the current Opcode
func (c *CPU) Mnemonic() string {
	return c.ins.mnemonic
}

// AddPC adds n to the current PC
func (c *CPU) AddPC(n int) {
	c.pc += n
}

// Op returns operand o as a byte
func (c *CPU) Op(o int) byte {
	return c.ins.ops[0%2]
}

// // Op8 returns the 8-bit signed Operand as an int
// func (c *CPU) Op8() int8 {
// 	return int8(c.ins.ops[0])
// }

// // OpU8 returns the 8-bit unsigned Operand as an int
// func (c *CPU) OpU8() byte {
// 	return c.ins.ops[0]
// }

// OpU16 returns the 16-bit unsigned Operand as an int
func (c *CPU) OpU16() int {
	return int(c.ins.ops[0]) | int(c.ins.ops[1])<<8
}

// A returns the a register
func (c *CPU) A() byte {
	return c.a
}

// SetA sets the a register
func (c *CPU) SetA(b byte) {
	c.a = b
}

// Pa returns a pointer to the a register
func (c *CPU) Pa() *byte {
	return &c.a
}

// X returns the x register
func (c *CPU) X() byte {
	return c.x
}

// SetX sets the x register
func (c *CPU) SetX(b byte) {
	c.x = b
}

// IncX incriment the x register
func (c *CPU) IncX() {
	c.x++
}

// DecX decriment the x register
func (c *CPU) DecX() {
	c.x--
}

// Y returns the y register
func (c *CPU) Y() byte {
	return c.y
}

// SetY sets the y register
func (c *CPU) SetY(b byte) {
	c.y = b
}

// IncY incriment the x register
func (c *CPU) IncY() {
	c.y++
}

// DecY decriment the y register
func (c *CPU) DecY() {
	c.y--
}

// Status returns the status register
func (c *CPU) Status() byte {
	return c.status
}

// SetStatus returns the status register
func (c *CPU) SetStatus(s byte) {
	c.status |= s
}

// ClearStatus returns the status register
func (c *CPU) ClearStatus(s byte) {
	c.status &^= s
}

// Mem slice overlaying the memory
func (c *CPU) Mem() []byte {
	return c.mem[:]
}

// Mem16 returns the 16-bit unsigned value from memory as an int
func (c *CPU) Mem16(a int) int {
	return int(c.mem[a]) | int(c.mem[a])<<8
}

// Push16 pushes a 16-bit value onto the stack
func (c *CPU) Push16(i int16) {
	c.sp += 2
	c.stack[0xff-c.sp] = byte(i)
	c.stack[(0xff-c.sp)+1] = byte(i >> 8)
}

// Pop16 pops a 16-bit value from the stack
func (c *CPU) Pop16() int {
	i := int(c.stack[0xff-c.sp]) | int(c.stack[(0xff-c.sp)+1])<<8
	c.sp -= 2
	return i
}

// BranchAddr returns the destination of a branch instruction
func (c *CPU) BranchAddr() int {
	return c.pc + int(int8(c.ins.ops[0])) + 2
}

// AbsJumpAddr returns the destination of an absolute  jmp/jsr instruction
func (c *CPU) AbsJumpAddr() int {
	return int(c.ins.ops[0]) | int(c.ins.ops[1])<<8
}

// TODO indirect jump

// FetchInstr fetches a CPU instuction
func (c *CPU) FetchInstr() {
	c.ins.code = int(c.mem[c.pc])
	c.ins.mnemonic = opcodes[c.ins.code].mnemonic
	c.ins.length = opcodes[c.ins.code].length

	c.ins.ops = c.mem[c.pc+1 : c.pc+3]
	c.ins.mode = opcodes[c.ins.code].mode
	//c.next = c.pc + c.ins.length
}
