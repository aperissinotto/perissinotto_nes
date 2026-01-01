package cpu

type Instruction struct {
	name   string
	exec   func() bool
	cycles uint8
}

func (c *CPU) initOpcodes() {
	c.opcodes[0xA1] = Instruction{
		name:   "LDA ($nn,X)",
		exec:   c.opLDA_INDX,
		cycles: 6,
	}

	c.opcodes[0xA2] = Instruction{
		name:   "LDX #imm",
		exec:   c.opLDX_IMM,
		cycles: 2,
	}

	c.opcodes[0xA5] = Instruction{
		name:   "LDA $nn",
		exec:   c.opLDA_ZP,
		cycles: 3,
	}

	c.opcodes[0xA6] = Instruction{
		name:   "LDX $nn",
		exec:   c.opLDX_ZP,
		cycles: 3,
	}

	c.opcodes[0xA9] = Instruction{
		name:   "LDA #imm",
		exec:   c.opLDA_IMM,
		cycles: 2,
	}

	c.opcodes[0xAD] = Instruction{
		name:   "LDA $nnnn",
		exec:   c.opLDA_ABS,
		cycles: 4,
	}

	c.opcodes[0xB1] = Instruction{
		name:   "LDA ($nn),Y",
		exec:   c.opLDA_INDY,
		cycles: 5,
	}

	c.opcodes[0xB5] = Instruction{
		name:   "LDA $nn,X",
		exec:   c.opLDA_ZPX,
		cycles: 4,
	}

	c.opcodes[0xB9] = Instruction{
		name:   "LDA $nnnn,Y",
		exec:   c.opLDA_ABSY,
		cycles: 4,
	}

	c.opcodes[0xBD] = Instruction{
		name:   "LDA $nnnn,X",
		exec:   c.opLDA_ABSX,
		cycles: 4,
	}

	c.opcodes[0xEA] = Instruction{
		name:   "NOP",
		exec:   c.opNOP,
		cycles: 2,
	}
}
