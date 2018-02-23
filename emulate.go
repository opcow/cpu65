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
	0xca: (*CPU).emuDex,
	0x88: (*CPU).emuDey,
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
}

// Execute execute the current CPU instruction
func (c *CPU) Execute() (byte, error) {
	if ExecuteInstr[c.Instr.opcode] == nil {
		return c.Instr.opcode, fmt.Errorf("looking up opcode: %0x", c.Instr.opcode)
	}
	ExecuteInstr[c.Instr.opcode](c)
	c.PC += c.Instr.Length
	return c.Instr.opcode, nil
}

func (c *CPU) getEffAddr() *byte {
	mode := c.Instr.Mode
	switch mode {
	case imm:
		return &c.Mem[c.PC+1]
	case zpa:
		return &c.Mem[c.Instr.Ops[0]]
	case zpx:
		return &c.Mem[c.Instr.Ops[0]+c.X]
	case abs:
		return &c.Mem[c.OpU16()]
	case abx:
		return &c.Mem[c.OpU16()+int(c.X)]
	case aby:
		return &c.Mem[c.OpU16()+int(c.Y)]
	case inx:
		return &c.Mem[c.Mem16(int(c.Instr.Ops[0]+c.X))]
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

func (c *CPU) emuLda() {
	c.A = *c.getEffAddr()
	c.setNZReg(c.A)
}

func (c *CPU) emuLdx() {
	c.X = *c.getEffAddr()
	c.setNZReg(c.X)
}

func (c *CPU) emuLdy() {
	c.Y = *c.getEffAddr()
	c.setNZReg(c.Y)
}

func (c *CPU) emuSta() {
	*c.getEffAddr() = c.A
}

func (c *CPU) emuCld() {
	c.Status &^= StatusD
}

func (c *CPU) emuClc() {
	c.Status &^= StatusC
}

func (c *CPU) emuDex() {
	c.X--
	c.setNZReg(c.X)
}

func (c *CPU) emuDey() {
	c.Y--
	c.setNZReg(c.Y)
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
		c.PC += int(int8(c.Instr.Ops[0]))
	}
}

func (c *CPU) emuJsr() {
	c.Push16(int16(c.PC + 2))
	c.PC = c.OpU16() - c.Instr.Length
}

func (c *CPU) emuBit() {
	var n byte
	if c.Instr.opcode == 0x24 {
		n = c.Mem[c.Instr.Ops[0]]
	} else {
		n = c.Mem[c.OpU16()]
	}
	c.Status |= n & 0xc0
	if c.A&n == 0 {
		c.Status |= StatusZ
	} else {
		c.Status &^= StatusZ
	}
}

func (c *CPU) emuLsr() {
	var bit0 byte
	addr := c.getEffAddr()
	bit0 = *addr & 1
	*addr >>= 1
	c.Status |= bit0
}

func (c *CPU) emuAsl() {
	var bit7 byte
	addr := c.getEffAddr()
	bit7 = *addr & (1 << 7)
	*addr <<= 1
	c.Status |= bit7 >> 7
}

func (c *CPU) emuRol() {
	addr := c.getEffAddr()
	bit7 := *addr & (1 << 7)
	*addr <<= 1
	c.Status |= bit7 >> 7
	*addr |= c.Status & StatusC
}

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
}

func (c *CPU) emuRts() {
	c.PC = c.Pop16() + 1
}

func (c *CPU) emuBrk() {
}
