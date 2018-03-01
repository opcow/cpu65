package cpu65

import "fmt"

func TestOpcodes() {
	for i := range opcodes {
		if opcodes[i].code != 0 {
			fmt.Printf("%s %08b %08b\n", opcodes[i].mnemonic, byte(opcodes[i].code)&0xe3,
				byte(opcodes[i].mode)&0xe3)
		}
	}
}
