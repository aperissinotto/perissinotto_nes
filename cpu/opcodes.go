package cpu

type Instruction struct {
	name   string
	exec   func() bool
	cycles uint8
}

func (c *CPU) initOpcodes() {
	c.opcodes[0xA0] = Instruction{
		name:   "LDY #imm",
		exec:   c.opLDY_IMM,
		cycles: 2,
	}

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

	c.opcodes[0xA4] = Instruction{
		name:   "LDY $nn",
		exec:   c.opLDY_ZP,
		cycles: 3,
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

	c.opcodes[0xAA] = Instruction{
		name:   "TAX X=A",
		exec:   c.opTAX,
		cycles: 2,
	}

	c.opcodes[0xAC] = Instruction{
		name:   "LDY $nnnn",
		exec:   c.opLDY_ABS,
		cycles: 4,
	}

	c.opcodes[0xAD] = Instruction{
		name:   "LDA $nnnn",
		exec:   c.opLDA_ABS,
		cycles: 4,
	}

	c.opcodes[0xAE] = Instruction{
		name:   "LDX $nnnn",
		exec:   c.opLDX_ABS,
		cycles: 4,
	}

	c.opcodes[0xB1] = Instruction{
		name:   "LDA ($nn),Y",
		exec:   c.opLDA_INDY,
		cycles: 5,
	}

	c.opcodes[0xB4] = Instruction{
		name:   "LDY $nn,X",
		exec:   c.opLDY_ZPX,
		cycles: 4,
	}

	c.opcodes[0xB5] = Instruction{
		name:   "LDA $nn,X",
		exec:   c.opLDA_ZPX,
		cycles: 4,
	}

	c.opcodes[0xB6] = Instruction{
		name:   "LDX $nn,Y",
		exec:   c.opLDX_ZPY,
		cycles: 4,
	}

	c.opcodes[0xB9] = Instruction{
		name:   "LDA $nnnn,Y",
		exec:   c.opLDA_ABSY,
		cycles: 4,
	}

	c.opcodes[0xBC] = Instruction{
		name:   "LDY $nnnn,X",
		exec:   c.opLDY_ABSX,
		cycles: 4,
	}

	c.opcodes[0xBD] = Instruction{
		name:   "LDA $nnnn,X",
		exec:   c.opLDA_ABSX,
		cycles: 4,
	}

	c.opcodes[0xBE] = Instruction{
		name:   "LDX $nnnn,Y",
		exec:   c.opLDX_ABSY,
		cycles: 4,
	}

	c.opcodes[0xEA] = Instruction{
		name:   "NOP",
		exec:   c.opNOP,
		cycles: 2,
	}
}
