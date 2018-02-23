package cpu65

import (
	"fmt"
)

func (i *instruction) String() string {
	switch i.Mode {
	case imm:
		return fmt.Sprintf("%s #$%02x", i.Mnemonic, i.Ops[0])
	case zpa:
		return fmt.Sprintf("%s $%02x", i.Mnemonic, i.Ops[0])
	case zpx:
		return fmt.Sprintf("%s $%02x,X", i.Mnemonic, i.Ops[0])
	case abs:
		return fmt.Sprintf("%s $%02x%02x", i.Mnemonic, i.Ops[1], i.Ops[0])
	case abx:
		return fmt.Sprintf("%s $%02x%02x,X", i.Mnemonic, i.Ops[1], i.Ops[0])
	case aby:
		return fmt.Sprintf("%s $%02x%02x,Y", i.Mnemonic, i.Ops[1], i.Ops[0])
	case ind:
		return fmt.Sprintf("%s ($%02x%02x)", i.Mnemonic, i.Ops[1], i.Ops[0])
	case inx:
		return fmt.Sprintf("%s ($%02x,X)", i.Mnemonic, i.Ops[0])
	case iny:
		return fmt.Sprintf("%s ($%02x),Y", i.Mnemonic, i.Ops[0])
	case acc:
		return fmt.Sprintf("%s", i.Mnemonic)
	case rel:
		return fmt.Sprintf("%s $%02x", i.Mnemonic, i.Ops[0])
		// return fmt.Sprintf("%s $%02x ($%04x)", i.Mnemonic, i.Ops[0],
		// c.PC+i.Length+int(int8(i.Ops[0])))
	case imp:
		return fmt.Sprintf("%s", i.Mnemonic)
	default:
		return fmt.Sprintf("%s", i.Mnemonic)
	}
}

type CpuStatus byte

func (c CpuStatus) String() string {
	st := [8]byte{'-', '-', '-', '-', '-', '-', '-', '-'}
	s := [8]byte{'N', 'V', 'U', 'B', 'D', 'I', 'Z', 'C'}
	b := CpuStatus(1 << 7)
	for i := 0; i < 8; i++ {
		if c&b != 0 {
			st[i] = s[i]
		}
		b >>= 1
	}
	return string(st[:])
}
