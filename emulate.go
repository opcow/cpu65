package cpu65

import "fmt"

// ExecuteInstr is an array holding function which emulate 6502 instructions
var ExecuteInstr = [256]func(c *CPU){
	0x00: (*CPU).emuBrk,
	0xa9: (*CPU).emuLda,
	0xa5: (*CPU).emuLda,
	0xb5: (*CPU).emuLda,
	0xad: (*CPU).emuLda,
	0xbd: (*CPU).emuLda,
	0xb9: (*CPU).emuLda,
	0xa1: (*CPU).emuLda,
	0xb1: (*CPU).emuLda,
	0xa2: (*CPU).emuLdx,
	0xa6: (*CPU).emuLdx,
	0xb6: (*CPU).emuLdx,
	0xae: (*CPU).emuLdx,
	0xbe: (*CPU).emuLdx,
	0xa0: (*CPU).emuLdy,
	0xa4: (*CPU).emuLdy,
	0xb4: (*CPU).emuLdy,
	0xac: (*CPU).emuLdy,
	0xbc: (*CPU).emuLdy,
	0x85: (*CPU).emuSta,
	0x95: (*CPU).emuSta,
	0x8d: (*CPU).emuSta,
	0x9d: (*CPU).emuSta,
	0x99: (*CPU).emuSta,
	0x81: (*CPU).emuSta,
	0x91: (*CPU).emuSta,
	0x10: (*CPU).emuBra,
	0x30: (*CPU).emuBra,
	0x50: (*CPU).emuBra,
	0x70: (*CPU).emuBra,
	0x90: (*CPU).emuBra,
	0xb0: (*CPU).emuBra,
	0xd0: (*CPU).emuBra,
	0xf0: (*CPU).emuBra,
	0xd8: (*CPU).emuCld,
	0x20: (*CPU).emuJsr,
	0x24: (*CPU).emuBit,
	0x2c: (*CPU).emuBit,
	0x4a: (*CPU).emuLsr,
	0x46: (*CPU).emuLsr,
	0x56: (*CPU).emuLsr,
	0x4e: (*CPU).emuLsr,
	0x5e: (*CPU).emuLsr,
	0x18: (*CPU).emuClc,
	0x69: (*CPU).emuAdc,
	0x65: (*CPU).emuAdc,
	0x75: (*CPU).emuAdc,
	0x6d: (*CPU).emuAdc,
	0x7d: (*CPU).emuAdc,
	0x79: (*CPU).emuAdc,
	0x61: (*CPU).emuAdc,
	0x71: (*CPU).emuAdc,
	0x0a: (*CPU).emuAsl,
	0x06: (*CPU).emuAsl,
	0x16: (*CPU).emuAsl,
	0x0e: (*CPU).emuAsl,
	0x1e: (*CPU).emuAsl,
	0x2a: (*CPU).emuRol,
	0x26: (*CPU).emuRol,
	0x36: (*CPU).emuRol,
	0x2e: (*CPU).emuRol,
	0x3e: (*CPU).emuRol,
	0xe9: (*CPU).emuSbc,
	0xe5: (*CPU).emuSbc,
	0xf5: (*CPU).emuSbc,
	0xed: (*CPU).emuSbc,
	0xfd: (*CPU).emuSbc,
	0xf9: (*CPU).emuSbc,
	0xe1: (*CPU).emuSbc,
	0xf1: (*CPU).emuSbc,
	0x60: (*CPU).emuRts,
	0x6a: (*CPU).emuRor,
	0x66: (*CPU).emuRor,
	0x76: (*CPU).emuRor,
	0x6e: (*CPU).emuRor,
	0x7e: (*CPU).emuRor,
	0x9a: (*CPU).emuTxs,
	0xba: (*CPU).emuTsx,
	0x48: (*CPU).emuPha,
	0x68: (*CPU).emuPla,
	0x08: (*CPU).emuPhp,
	0x28: (*CPU).emuPlp,
	0xea: (*CPU).emuNop,
	0x29: (*CPU).emuAnd,
	0x25: (*CPU).emuAnd,
	0x35: (*CPU).emuAnd,
	0x2d: (*CPU).emuAnd,
	0x3d: (*CPU).emuAnd,
	0x39: (*CPU).emuAnd,
	0x21: (*CPU).emuAnd,
	0x31: (*CPU).emuAnd,
	0x09: (*CPU).emuOra,
	0x05: (*CPU).emuOra,
	0x15: (*CPU).emuOra,
	0x0d: (*CPU).emuOra,
	0x1d: (*CPU).emuOra,
	0x19: (*CPU).emuOra,
	0x01: (*CPU).emuOra,
	0x11: (*CPU).emuOra,
	0x49: (*CPU).emuEor,
	0x45: (*CPU).emuEor,
	0x55: (*CPU).emuEor,
	0x4d: (*CPU).emuEor,
	0x5d: (*CPU).emuEor,
	0x59: (*CPU).emuEor,
	0x41: (*CPU).emuEor,
	0x51: (*CPU).emuEor,
	0x86: (*CPU).emuStx,
	0x96: (*CPU).emuStx,
	0x8e: (*CPU).emuStx,
	0x84: (*CPU).emuSty,
	0x94: (*CPU).emuSty,
	0x8c: (*CPU).emuSty,
	0xaa: (*CPU).emuTax,
	0xa8: (*CPU).emuTay,
	0x8a: (*CPU).emuTxa,
	0x98: (*CPU).emuTya,
	0xc9: (*CPU).emuCmp,
	0xc5: (*CPU).emuCmp,
	0xd5: (*CPU).emuCmp,
	0xcd: (*CPU).emuCmp,
	0xdd: (*CPU).emuCmp,
	0xd9: (*CPU).emuCmp,
	0xc1: (*CPU).emuCmp,
	0xd1: (*CPU).emuCmp,
	0xe0: (*CPU).emuCpx,
	0xe4: (*CPU).emuCpx,
	0xec: (*CPU).emuCpx,
	0xc0: (*CPU).emuCpy,
	0xc4: (*CPU).emuCpy,
	0xcc: (*CPU).emuCpy,
	0xe6: (*CPU).emuInc,
	0xf6: (*CPU).emuInc,
	0xee: (*CPU).emuInc,
	0xfe: (*CPU).emuInc,
	0xc6: (*CPU).emuDec,
	0xd6: (*CPU).emuDec,
	0xce: (*CPU).emuDec,
	0xde: (*CPU).emuDec,
	0xe8: (*CPU).emuInx,
	0xc8: (*CPU).emuIny,
	0xca: (*CPU).emuDex,
	0x88: (*CPU).emuDey,
	0x4c: (*CPU).emuJmp,
	0x6c: (*CPU).emuJmp,
}

// Execute execute the current CPU instruction
func (c *CPU) Execute() (byte, error) {
	if ExecuteInstr[c.Instr.opcode] == nil {
		return c.Instr.opcode, fmt.Errorf("looking up emulation for instruction: %s", c.Instr.Mnemonic)
	}
	ExecuteInstr[c.Instr.opcode](c)
	return c.Instr.opcode, nil
}

func (c *CPU) getEffAddr() *byte {
	mode := c.Instr.Mode
	switch mode {
	case imm:
		return &c.Mem[c.PC+1]
	case zpa:
		return &c.Mem[c.Instr.Op[0]]
	case zpx:
		return &c.Mem[c.Instr.Op[0]+c.X]
	case abs:
		return &c.Mem[c.OpU16()]
	case abx:
		return &c.Mem[c.OpU16()+int(c.X)]
	case aby:
		return &c.Mem[c.OpU16()+int(c.Y)]
	case inx:
		return &c.Mem[c.Mem16(int(c.Instr.Op[0]+c.X))]
	case iny:
		return &c.Mem[c.OpU16()+int(c.Y)]
	case acc:
		return &c.A
	}
	return nil
}

func (c *CPU) setNZReg(reg byte) {
	if reg == 0 {
		c.Status |= StatusZ
		c.Status &^= StatusN
	} else {
		c.Status &^= StatusZ
		if reg&(1<<7) != 0 {
			c.Status |= StatusN
		} else {
			c.Status &^= StatusN
		}
	}
}

// storage
func (c *CPU) emuLda() {
	c.A = *c.getEffAddr()
	c.setNZReg(c.A)
	c.PC += c.Instr.Length
}

func (c *CPU) emuLdx() {
	c.X = *c.getEffAddr()
	c.setNZReg(c.X)
	c.PC += c.Instr.Length
}

func (c *CPU) emuLdy() {
	c.Y = *c.getEffAddr()
	c.setNZReg(c.Y)
	c.PC += c.Instr.Length
}

func (c *CPU) emuSta() {
	*c.getEffAddr() = c.A
	c.PC += c.Instr.Length
}

func (c *CPU) emuStx() {
	*c.getEffAddr() = c.X
	c.PC += c.Instr.Length
}

func (c *CPU) emuSty() {
	*c.getEffAddr() = c.Y
	c.PC += c.Instr.Length
}

func (c *CPU) emuTax() {
	c.X = c.A
	c.setNZReg(c.X)
	c.PC += c.Instr.Length
}

func (c *CPU) emuTay() {
	c.Y = c.A
	c.setNZReg(c.Y)
	c.PC += c.Instr.Length
}

func (c *CPU) emuTxa() {
	c.A = c.X
	c.setNZReg(c.A)
	c.PC += c.Instr.Length
}

func (c *CPU) emuTya() {
	c.A = c.Y
	c.setNZReg(c.A)
	c.PC += c.Instr.Length
}

// registers
func (c *CPU) emuCld() {
	c.Status &^= StatusD
	c.PC += c.Instr.Length
}

func (c *CPU) emuClc() {
	c.Status &^= StatusC
	c.PC += c.Instr.Length
}

func (c *CPU) emuBra() {
	var status bool
	o := c.Instr.opcode
	switch o {
	case 0x10:
		status = (c.Status & StatusN) == 0
	case 0x30:
		status = (c.Status & StatusN) != 0
	case 0x50:
		status = (c.Status & StatusV) == 0
	case 0x70:
		status = (c.Status & StatusV) != 0
	case 0x90:
		status = (c.Status & StatusC) == 0
	case 0xb0:
		status = (c.Status & StatusC) != 0
	case 0xd0:
		status = (c.Status & StatusZ) == 0
	case 0xf0:
		status = (c.Status & StatusZ) != 0
	}
	if status {
		c.PC += c.Instr.Length + int(int8(c.Instr.Op[0]))
	} else {
		c.PC += c.Instr.Length
	}
}

func (c *CPU) emuJsr() {
	c.Push16(int16(c.PC + 2))
	c.PC = c.OpU16()
}

func (c *CPU) emuJmp() {
	if c.Instr.Mode == abs {
		c.PC = c.OpU16()
	} else {
		a := c.OpU16()
		c.PC = int(c.Mem[a]) + int(c.Mem[a])<<8
	}
}

func (c *CPU) emuBit() {
	var n byte
	if c.Instr.opcode == 0x24 {
		n = c.Mem[c.Instr.Op[0]]
	} else {
		n = c.Mem[c.OpU16()]
	}
	c.Status |= n & 0xc0
	if c.A&n == 0 {
		c.Status |= StatusZ
	} else {
		c.Status &^= StatusZ
	}
	c.PC += c.Instr.Length
}

// add and subtract TODO: BCD
func (c *CPU) emuAdc() {
	n := int(c.A) + int(*c.getEffAddr())
	if int8(n) < -128 || int8(n) > 127 {
		c.Status |= StatusV
	} else {
		c.Status &^= StatusV
	}
	if n > 255 { // fixme see if adc and sbc can be consolidated
		c.Status |= StatusC
	} else {
		c.Status &^= StatusC
	}
	c.A = byte(n)
	c.PC += c.Instr.Length
}

func (c *CPU) emuSbc() {
	n := int(c.A) - int(*c.getEffAddr())
	if int8(n) < -128 || int8(n) > 127 {
		c.Status |= StatusV
	} else {
		c.Status &^= StatusV
	}
	if n > -0 {
		c.Status |= StatusC
	} else {
		c.Status &^= StatusC
	}
	c.A = byte(n)
	c.PC += c.Instr.Length
}

func (c *CPU) emuInc() {
	m := c.getEffAddr()
	*m++
	c.setNZReg(*m)
	c.PC += c.Instr.Length
}

func (c *CPU) emuDec() {
	m := c.getEffAddr()
	*m--
	c.setNZReg(*m)
	c.PC += c.Instr.Length
}

func (c *CPU) emuInx() {
	c.X++
	c.setNZReg(c.X)
	c.PC += c.Instr.Length
}

func (c *CPU) emuIny() {
	c.Y++
	c.setNZReg(c.Y)
	c.PC += c.Instr.Length
}

func (c *CPU) emuDex() {
	c.X--
	c.setNZReg(c.X)
	c.PC += c.Instr.Length
}

func (c *CPU) emuDey() {
	c.Y--
	c.setNZReg(c.Y)
	c.PC += c.Instr.Length
}

// compare
func (c *CPU) emuCmp() {
	r := c.A - *c.getEffAddr()
	c.setNZReg(r)
	if r >= 0 {
		c.Status |= StatusC
	} else {
		c.Status &^= StatusC
	}
	c.PC += c.Instr.Length
}

func (c *CPU) emuCpx() {
	r := c.X - *c.getEffAddr()
	c.setNZReg(r)
	if r >= 0 {
		c.Status |= StatusC
	} else {
		c.Status &^= StatusC
	}
	c.PC += c.Instr.Length
}

func (c *CPU) emuCpy() {
	r := c.Y - *c.getEffAddr()
	c.setNZReg(r)
	if r >= 0 {
		c.Status |= StatusC
	} else {
		c.Status &^= StatusC
	}
	c.PC += c.Instr.Length
}

// bitwise ops
func (c *CPU) emuAnd() {
	c.A &= *c.getEffAddr()
	c.setNZReg(c.A)
	c.PC += c.Instr.Length
}

func (c *CPU) emuOra() {
	c.A |= *c.getEffAddr()
	c.setNZReg(c.A)
	c.PC += c.Instr.Length
}

func (c *CPU) emuEor() {
	c.A ^= *c.getEffAddr()
	c.setNZReg(c.A)
	c.PC += c.Instr.Length
}

func (c *CPU) emuLsr() {
	var bit0 byte
	addr := c.getEffAddr()
	bit0 = *addr & 1
	*addr >>= 1
	c.Status |= bit0
	c.PC += c.Instr.Length
}

func (c *CPU) emuAsl() {
	var bit7 byte
	addr := c.getEffAddr()
	bit7 = *addr & (1 << 7)
	*addr <<= 1
	c.Status |= bit7 >> 7
	c.PC += c.Instr.Length
}

func (c *CPU) emuRol() {
	addr := c.getEffAddr()
	bit7 := *addr & (1 << 7)
	*addr <<= 1
	c.Status |= bit7 >> 7
	*addr |= c.Status & StatusC
	c.PC += c.Instr.Length
}

func (c *CPU) emuRor() {
	addr := c.getEffAddr()
	bit0 := *addr & (1 << 0)
	*addr >>= 1
	c.Status |= bit0
	*addr |= c.Status & (StatusC << 7)
	c.PC += c.Instr.Length
}

// stack instructions
func (c *CPU) emuPha() {
	c.stack[c.sp] = c.A
	c.sp--
	c.PC += c.Instr.Length
}

func (c *CPU) emuPla() {
	c.sp++
	c.A = c.stack[c.sp]
	c.setNZReg(c.A)
	c.PC += c.Instr.Length
}

func (c *CPU) emuTxs() {
	c.sp = c.X
	c.PC += c.Instr.Length
}

func (c *CPU) emuTsx() {
	c.X = c.sp
	c.setNZReg(c.X)
	c.PC += c.Instr.Length
}

func (c *CPU) emuPhp() {
	c.stack[c.sp] = c.Status
	c.sp--
	c.PC += c.Instr.Length
}

func (c *CPU) emuPlp() {
	c.sp++
	c.Status = c.stack[c.sp]
	c.PC += c.Instr.Length
}

func (c *CPU) emuRts() {
	c.PC = c.Pop16() + 1
}

func (c *CPU) emuRti() {
	c.Status = c.stack[c.sp+1]
	c.sp++
	c.PC = c.Pop16() + 1
}

func (c *CPU) emuBrk() {
	c.PC += 2
}

func (c *CPU) emuNop() {
	c.PC += c.Instr.Length
}
