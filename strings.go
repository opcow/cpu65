package cpu65

import (
	"fmt"
)

func (i *instruction) String() string {
	switch i.Mode {
	case imm:
		return fmt.Sprintf("%s #$%02x", i.Mnemonic, i.Op[0])
	case zpa:
		return fmt.Sprintf("%s $%02x", i.Mnemonic, i.Op[0])
	case zpx:
		return fmt.Sprintf("%s $%02x,X", i.Mnemonic, i.Op[0])
	case abs:
		return fmt.Sprintf("%s $%02x%02x", i.Mnemonic, i.Op[1], i.Op[0])
	case abx:
		return fmt.Sprintf("%s $%02x%02x,X", i.Mnemonic, i.Op[1], i.Op[0])
	case aby:
		return fmt.Sprintf("%s $%02x%02x,Y", i.Mnemonic, i.Op[1], i.Op[0])
	case ind:
		return fmt.Sprintf("%s ($%02x%02x)", i.Mnemonic, i.Op[1], i.Op[0])
	case inx:
		return fmt.Sprintf("%s ($%02x,X)", i.Mnemonic, i.Op[0])
	case iny:
		return fmt.Sprintf("%s ($%02x),Y", i.Mnemonic, i.Op[0])
	case acc:
		return fmt.Sprintf("%s", i.Mnemonic)
	case rel:
		return fmt.Sprintf("%s $%02x", i.Mnemonic, i.Op[0])
		// return fmt.Sprintf("%s $%02x ($%04x)", i.Mnemonic, i.Op[0],
		// c.PC+i.Length+int(int8(i.Op[0])))
	case imp:
		return fmt.Sprintf("%s", i.Mnemonic)
	default:
		return fmt.Sprintf("%s", i.Mnemonic)
	}
}

// Result returns a string showing the results of the instruction's execution
// or the destination of a jump
func (c *CPU) Result() string {
	switch c.Instr.Mode {
	// case imm:
	// 	return fmt.Sprintf("($%02x)", *c.getEffAddr())
	// case zpa:
	// 	return fmt.Sprintf("z($%02x)", *c.getEffAddr())
	// case zpx:
	// 	return fmt.Sprintf("($%02x)", *c.getEffAddr())
	case rel:
		return fmt.Sprintf("(%04X)", c.PC+c.Instr.Length+int(int8(c.Instr.Op[0])))
	default:
		return ""
	}
}

// CPUStatus used to represent the status register as a string
type CPUStatus byte

func (c CPUStatus) String() string {
	st := [8]byte{'-', '-', '-', '-', '-', '-', '-', '-'}
	s := [8]byte{'N', 'V', 'U', 'B', 'D', 'I', 'Z', 'C'}
	b := CPUStatus(1 << 7)
	for i := 0; i < 8; i++ {
		if c&b != 0 {
			st[i] = s[i]
		}
		b >>= 1
	}
	return string(st[:])
}
