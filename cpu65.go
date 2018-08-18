// Package cpu65 a 6502 disassembler/assembler
package cpu65

// MaxMem is the maximum memory size for or CPU
const (
	MaxMem = 0x10000
	Modes  = imp + 1
)

// CPU modes
const (
	inx = 0x00 // indirect x
	zpa = 0x01 // zero page
	imm = 0x02 // immediate mode
	abs = 0x03 // absolute
	iny = 0x04 // indirect y
	zpx = 0x05 // zero page x
	aby = 0x06 // absolute y
	abx = 0x07 // absolute x
	ind = 0x08 // indirect
	acc = 0x0a // accumulator
	rel = 0x0b // relative
	imp = 0x0c // implied
	zpy = 0x0d // zero page y
	unk = 0x0f
)

const (
	none = iota
	preg
	areg
	xreg
	yreg
	memm
	addr
)

type opcode struct {
	code     int // the opcode
	length   int
	mode     int
	target   int // the thing that gets the result
	mnemonic string
}

var opcodes = [256]opcode{
	0x69: opcode{0x69, 2, imm, areg, "ADC"},
	0x65: opcode{0x65, 2, zpa, areg, "ADC"},
	0x75: opcode{0x75, 2, zpx, areg, "ADC"},
	0x6d: opcode{0x6d, 3, abs, areg, "ADC"},
	0x7d: opcode{0x7d, 3, abx, areg, "ADC"},
	0x79: opcode{0x79, 3, aby, areg, "ADC"},
	0x61: opcode{0x61, 2, inx, areg, "ADC"},
	0x71: opcode{0x71, 2, iny, areg, "ADC"},
	0x29: opcode{0x29, 2, imm, areg, "AND"},
	0x25: opcode{0x25, 2, zpa, areg, "AND"},
	0x35: opcode{0x35, 2, zpx, areg, "AND"},
	0x2d: opcode{0x2d, 3, abs, areg, "AND"},
	0x3d: opcode{0x3d, 3, abx, areg, "AND"},
	0x39: opcode{0x39, 3, aby, areg, "AND"},
	0x21: opcode{0x21, 2, inx, areg, "AND"},
	0x31: opcode{0x31, 2, iny, areg, "AND"},
	0x0a: opcode{0x0a, 1, acc, areg, "ASL"},
	0x06: opcode{0x06, 2, zpa, memm, "ASL"},
	0x16: opcode{0x16, 2, zpx, memm, "ASL"},
	0x0e: opcode{0x0e, 3, abs, memm, "ASL"},
	0x1e: opcode{0x1e, 3, abx, memm, "ASL"},
	0x24: opcode{0x24, 2, zpa, memm, "BIT"},
	0x2c: opcode{0x2c, 3, abs, memm, "BIT"},
	0x10: opcode{0x10, 2, rel, addr, "BPL"},
	0x30: opcode{0x30, 2, rel, addr, "BMI"},
	0x50: opcode{0x50, 2, rel, addr, "BVC"},
	0x70: opcode{0x70, 2, rel, addr, "BVS"},
	0x90: opcode{0x90, 2, rel, addr, "BCC"},
	0xb0: opcode{0xb0, 2, rel, addr, "BCS"},
	0xd0: opcode{0xd0, 2, rel, addr, "BNE"},
	0xf0: opcode{0xf0, 2, rel, addr, "BEQ"},
	0x00: opcode{0x00, 1, imp, none, "BRK"},
	0xc9: opcode{0xc9, 2, imm, none, "CMP"},
	0xc5: opcode{0xc5, 2, zpa, none, "CMP"},
	0xd5: opcode{0xd5, 2, zpx, none, "CMP"},
	0xcd: opcode{0xcd, 3, abs, none, "CMP"},
	0xdd: opcode{0xdd, 3, abx, none, "CMP"},
	0xd9: opcode{0xd9, 3, aby, none, "CMP"},
	0xc1: opcode{0xc1, 2, inx, none, "CMP"},
	0xd1: opcode{0xd1, 2, iny, none, "CMP"},
	0xe0: opcode{0xe0, 2, imm, none, "CPX"},
	0xe4: opcode{0xe4, 2, zpa, none, "CPX"},
	0xec: opcode{0xec, 3, abs, none, "CPX"},
	0xc0: opcode{0xc0, 2, imm, none, "CPY"},
	0xc4: opcode{0xc4, 2, zpa, none, "CPY"},
	0xcc: opcode{0xcc, 3, abs, none, "CPY"},
	0xc6: opcode{0xc6, 2, zpa, memm, "DEC"},
	0xd6: opcode{0xd6, 2, zpx, memm, "DEC"},
	0xce: opcode{0xce, 3, abs, memm, "DEC"},
	0xde: opcode{0xde, 3, abx, memm, "DEC"},
	0x49: opcode{0x49, 2, imm, areg, "EOR"},
	0x45: opcode{0x45, 2, zpa, areg, "EOR"},
	0x55: opcode{0x55, 2, zpx, areg, "EOR"},
	0x4d: opcode{0x4d, 3, abs, areg, "EOR"},
	0x5d: opcode{0x5d, 3, abx, areg, "EOR"},
	0x59: opcode{0x59, 3, aby, areg, "EOR"},
	0x41: opcode{0x41, 2, inx, areg, "EOR"},
	0x51: opcode{0x51, 2, iny, areg, "EOR"},
	0x18: opcode{0x18, 1, imp, none, "CLC"},
	0x38: opcode{0x38, 1, imp, none, "SEC"},
	0x58: opcode{0x58, 1, imp, none, "CLI"},
	0x78: opcode{0x78, 1, imp, none, "SEI"},
	0xb8: opcode{0xb8, 1, imp, none, "CLV"},
	0xd8: opcode{0xd8, 1, imp, none, "CLD"},
	0xf8: opcode{0xf8, 1, imp, none, "SED"},
	0xe6: opcode{0xe6, 2, zpa, memm, "INC"},
	0xf6: opcode{0xf6, 2, zpx, memm, "INC"},
	0xee: opcode{0xee, 3, abs, memm, "INC"},
	0xfe: opcode{0xfe, 3, abx, memm, "INC"},
	0x4c: opcode{0x4c, 3, abs, addr, "JMP"},
	0x6c: opcode{0x6c, 3, ind, addr, "JMP"},
	0x20: opcode{0x20, 3, abs, addr, "JSR"},
	0xa9: opcode{0xa9, 2, imm, areg, "LDA"},
	0xa5: opcode{0xa5, 2, zpa, areg, "LDA"},
	0xb5: opcode{0xb5, 2, zpx, areg, "LDA"},
	0xad: opcode{0xad, 3, abs, areg, "LDA"},
	0xbd: opcode{0xbd, 3, abx, areg, "LDA"},
	0xb9: opcode{0xb9, 3, aby, areg, "LDA"},
	0xa1: opcode{0xa1, 2, inx, areg, "LDA"},
	0xb1: opcode{0xb1, 2, iny, areg, "LDA"},
	0xa2: opcode{0xa2, 2, imm, xreg, "LDX"},
	0xa6: opcode{0xa6, 2, zpa, xreg, "LDX"},
	0xb6: opcode{0xb6, 2, zpy, xreg, "LDX"},
	0xae: opcode{0xae, 3, abs, xreg, "LDX"},
	0xbe: opcode{0xbe, 3, aby, xreg, "LDX"},
	0xa0: opcode{0xa0, 2, imm, yreg, "LDY"},
	0xa4: opcode{0xa4, 2, zpa, yreg, "LDY"},
	0xb4: opcode{0xb4, 2, zpx, yreg, "LDY"},
	0xac: opcode{0xac, 3, abs, yreg, "LDY"},
	0xbc: opcode{0xbc, 3, abx, yreg, "LDY"},
	0x4a: opcode{0x4a, 1, acc, areg, "LSR"},
	0x46: opcode{0x46, 2, zpa, memm, "LSR"},
	0x56: opcode{0x56, 2, zpx, memm, "LSR"},
	0x4e: opcode{0x4e, 3, abs, memm, "LSR"},
	0x5e: opcode{0x5e, 3, abx, memm, "LSR"},
	0xea: opcode{0xea, 1, imp, none, "NOP"},
	0x09: opcode{0x09, 2, imm, areg, "ORA"},
	0x05: opcode{0x05, 2, zpa, areg, "ORA"},
	0x15: opcode{0x15, 2, zpx, areg, "ORA"},
	0x0d: opcode{0x0d, 3, abs, areg, "ORA"},
	0x1d: opcode{0x1d, 3, abx, areg, "ORA"},
	0x19: opcode{0x19, 3, aby, areg, "ORA"},
	0x01: opcode{0x01, 2, inx, areg, "ORA"},
	0x11: opcode{0x11, 2, iny, areg, "ORA"},
	0xaa: opcode{0xaa, 1, imp, xreg, "TAX"},
	0x8a: opcode{0x8a, 1, imp, areg, "TXA"},
	0xca: opcode{0xca, 1, imp, xreg, "DEX"},
	0xe8: opcode{0xe8, 1, imp, xreg, "INX"},
	0xa8: opcode{0xa8, 1, imp, yreg, "TAY"},
	0x98: opcode{0x98, 1, imp, areg, "TYA"},
	0x88: opcode{0x88, 1, imp, yreg, "DEY"},
	0xc8: opcode{0xc8, 1, imp, yreg, "INY"},
	0x2a: opcode{0x2a, 1, acc, areg, "ROL"},
	0x26: opcode{0x26, 2, zpa, memm, "ROL"},
	0x36: opcode{0x36, 2, zpx, memm, "ROL"},
	0x2e: opcode{0x2e, 3, abs, memm, "ROL"},
	0x3e: opcode{0x3e, 3, abx, memm, "ROL"},
	0x6a: opcode{0x6a, 1, acc, areg, "ROR"},
	0x66: opcode{0x66, 2, zpa, memm, "ROR"},
	0x76: opcode{0x76, 2, zpx, memm, "ROR"},
	0x6e: opcode{0x6e, 3, abs, memm, "ROR"},
	0x7e: opcode{0x7e, 3, abx, memm, "ROR"},
	0x40: opcode{0x40, 1, imp, none, "RTI"},
	0x60: opcode{0x60, 1, imp, none, "RTS"},
	0xe9: opcode{0xe9, 2, imm, areg, "SBC"},
	0xe5: opcode{0xe5, 2, zpa, areg, "SBC"},
	0xf5: opcode{0xf5, 2, zpx, areg, "SBC"},
	0xed: opcode{0xed, 3, abs, areg, "SBC"},
	0xfd: opcode{0xfd, 3, abx, areg, "SBC"},
	0xf9: opcode{0xf9, 3, aby, areg, "SBC"},
	0xe1: opcode{0xe1, 2, inx, areg, "SBC"},
	0xf1: opcode{0xf1, 2, iny, areg, "SBC"},
	0x85: opcode{0x85, 2, zpa, memm, "STA"},
	0x95: opcode{0x95, 2, zpx, memm, "STA"},
	0x8d: opcode{0x8d, 3, abs, memm, "STA"},
	0x9d: opcode{0x9d, 3, abx, memm, "STA"},
	0x99: opcode{0x99, 3, aby, memm, "STA"},
	0x81: opcode{0x81, 2, inx, memm, "STA"},
	0x91: opcode{0x91, 2, iny, memm, "STA"},
	0x9a: opcode{0x9a, 1, imp, none, "TXS"},
	0xba: opcode{0xba, 1, imp, none, "TSX"},
	0x48: opcode{0x48, 1, imp, none, "PHA"},
	0x68: opcode{0x68, 1, imp, none, "PLA"},
	0x08: opcode{0x08, 1, imp, none, "PHP"},
	0x28: opcode{0x28, 1, imp, none, "PLP"},
	0x86: opcode{0x86, 2, zpa, memm, "STX"},
	0x96: opcode{0x96, 2, zpy, memm, "STX"},
	0x8e: opcode{0x8e, 3, abs, memm, "STX"},
	0x84: opcode{0x84, 2, zpa, memm, "STY"},
	0x94: opcode{0x94, 2, zpx, memm, "STY"},
	0x8c: opcode{0x8c, 3, abs, memm, "STY"},
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
	opcode   byte
	Length   int
	Mode     int
	Op       []byte
	Mnemonic string
}

// CPU virtual processor + memory type
// call AttachMem() before use
type CPU struct {
	A      byte
	X      byte
	Y      byte
	Status byte
	sp     byte
	PC     int
	Instr  instruction
	Mem    *[0x10000]byte
	stack  []byte
}

// Opcode returns the opcode at the current PC
func (c *CPU) Opcode() byte {
	return c.Instr.opcode
}

// AttachMem attaches memory to the CPU
func (c *CPU) AttachMem(m *[MaxMem]byte) {
	c.Mem = m
	c.stack = c.Mem[0x100:0x200]
	c.sp = 0xff
}

// Next sets the program counter to the next instruction
func (c *CPU) Next() int {
	c.PC += c.Instr.Length
	return c.PC
}

// OpU16 returns the 16-bit unsigned Operand as an int
func (c *CPU) OpU16() int {
	return int(c.Instr.Op[0]) | int(c.Instr.Op[1])<<8
}

// Mem16 returns the 16-bit unsigned value from memory as an int
func (c *CPU) Mem16(a int) int {
	return int(c.Mem[a]) | int(c.Mem[a+1])<<8
}

// Push16 pushes a 16-bit value onto the stack
func (c *CPU) Push16(i int16) {
	c.stack[c.sp-1] = byte(i)
	c.stack[c.sp] = byte(i >> 8)
	c.sp -= 2
}

// Pop16 pOp a 16-bit value from the stack
func (c *CPU) Pop16() int {
	i := int(c.stack[c.sp+1]) | int(c.stack[c.sp+2])<<8
	c.sp += 2
	return i
}

// BranchAddr returns the destination of a branch instruction
func (c *CPU) BranchAddr() int {
	return c.PC + int(int8(c.Instr.Op[0])) + 2
}

// absJumpAddr returns the destination of an absolute jmp/jsr instruction
func (c *CPU) absJumpAddr() int {
	return int(c.Instr.Op[0]) | int(c.Instr.Op[1])<<8
}

// indJumpAddr returns the destination of an imdrect jmp instruction
func (c *CPU) imdJumpAddr() int {
	addr := int(c.Instr.Op[0]) | int(c.Instr.Op[1])<<8
	return int(c.Mem[addr]) | int(c.Mem[addr+1])<<8
}

// FetchInstr fetches a CPU instuction
func (c *CPU) FetchInstr() {
	c.Instr.opcode = c.Mem[c.PC]
	c.Instr.Mnemonic = opcodes[c.Instr.opcode].mnemonic
	c.Instr.Length = opcodes[c.Instr.opcode].length
	c.Instr.Op = c.Mem[c.PC+1 : c.PC+3]
	c.Instr.Mode = opcodes[c.Instr.opcode].mode
}
