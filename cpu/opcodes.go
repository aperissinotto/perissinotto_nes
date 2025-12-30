package cpu

func (c *CPU) initOpcodes() {
	for i := 0; i < 256; i++ {
		c.opcodeTable[i] = (*CPU).opNOP
	}

	c.opcodeTable[0xA9] = (*CPU).opLDA_IMM  // LDA #$nn
	c.opcodeTable[0xA5] = (*CPU).opLDA_ZP   // LDA $nn
	c.opcodeTable[0xB5] = (*CPU).opLDA_ZPX  // LDA $nn,X
	c.opcodeTable[0xAD] = (*CPU).opLDA_ABS  // LDA $nnnn
	c.opcodeTable[0xBD] = (*CPU).opLDA_ABSX // LDA $nnnn,X
	c.opcodeTable[0xA2] = (*CPU).opLDX_IMM  // LDX #$nn
	c.opcodeTable[0xA6] = (*CPU).opLDX_ZP   // LDX $nn
}
