package cpu

// LDA ($nn,X) -> Opcode $A1
func (c *CPU) opLDA_INDX() bool {
	addr := c.addrIndirectX()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	return false
}

// LDX #$nn -> Opcode $A2
func (c *CPU) opLDX_IMM() bool {
	value := c.Bus.Read(c.PC)
	c.PC++

	c.X = value
	c.setZN(c.X)
	return false
}

// LDA $nn -> Opcode $A5
func (c *CPU) opLDA_ZP() bool {
	addr := c.addrZeroPage()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	return false
}

// LDX $nn -> Opcode $A6
func (c *CPU) opLDX_ZP() bool {
	addr := c.addrZeroPage()

	c.X = c.Bus.Read(uint16(addr))
	c.setZN(c.X)
	return false
}

// LDA #$nn -> Opcode $A9
func (c *CPU) opLDA_IMM() bool {
	value := c.Bus.Read(c.PC)
	c.PC++

	c.A = value
	c.setZN(c.A)
	return false
}

// LDA $nnnn -> Opcode $AD
func (c *CPU) opLDA_ABS() bool {
	addr := c.addrAbsolute()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	return false
}

// LDA ($nn),Y -> Opcode $B1
func (c *CPU) opLDA_INDY() bool {
	pageCrossed, addr := c.addrIndirectY()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)

	return pageCrossed
}

// LDA $nn,X -> Opcode $B5
func (c *CPU) opLDA_ZPX() bool {
	addr := c.addrZeroPageX()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	return false
}

// LDA $nnnn,Y -> Opcode $B9
func (c *CPU) opLDA_ABSY() bool {
	pageCrossed, addr := c.addrAbsoluteY()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)

	return pageCrossed
}

// LDA $nnnn,X -> Opcode $BD
func (c *CPU) opLDA_ABSX() bool {
	pageCrossed, addr := c.addrAbsoluteX()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)

	return pageCrossed
}

// NOP -> Opcode $EA
func (c *CPU) opNOP() bool {
	return false
}
