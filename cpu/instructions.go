package cpu

// LDY #$nn -> Opcode $A0
func (c *CPU) opLDY_IMM() bool {
	value := c.Bus.Read(c.PC)
	c.PC++

	c.Y = value
	c.setZN(c.Y)
	return false
}

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

// LDY $nn -> Opcode $A4
func (c *CPU) opLDY_ZP() bool {
	addr := c.addrZeroPage()

	c.Y = c.Bus.Read(uint16(addr))
	c.setZN(c.Y)
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

// TAX X=A -> Opcode $AA
func (c *CPU) opTAX() bool {
	c.X = c.A
	c.setZN(c.X)
	return false
}

// LDY $nnnn -> Opcode $AC
func (c *CPU) opLDY_ABS() bool {
	addr := c.addrAbsolute()

	c.Y = c.Bus.Read(uint16(addr))
	c.setZN(c.Y)
	return false
}

// LDA $nnnn -> Opcode $AD
func (c *CPU) opLDA_ABS() bool {
	addr := c.addrAbsolute()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	return false
}

// LDX $nnnn -> Opcode $AE
func (c *CPU) opLDX_ABS() bool {
	addr := c.addrAbsolute()

	c.X = c.Bus.Read(uint16(addr))
	c.setZN(c.X)
	return false
}

// LDA ($nn),Y -> Opcode $B1
func (c *CPU) opLDA_INDY() bool {
	pageCrossed, addr := c.addrIndirectY()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)

	return pageCrossed
}

// LDY $nn,X -> Opcode $B4
func (c *CPU) opLDY_ZPX() bool {
	addr := c.addrZeroPageX()

	c.Y = c.Bus.Read(uint16(addr))
	c.setZN(c.Y)
	return false
}

// LDA $nn,X -> Opcode $B5
func (c *CPU) opLDA_ZPX() bool {
	addr := c.addrZeroPageX()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	return false
}

// LDX $nn,Y -> Opcode $B6
func (c *CPU) opLDX_ZPY() bool {
	addr := c.addrZeroPageY()

	c.X = c.Bus.Read(uint16(addr))
	c.setZN(c.X)
	return false
}

// LDA $nnnn,Y -> Opcode $B9
func (c *CPU) opLDA_ABSY() bool {
	pageCrossed, addr := c.addrAbsoluteY()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)

	return pageCrossed
}

// LDY $nnnn,X -> Opcode $BC
func (c *CPU) opLDY_ABSX() bool {
	pageCrossed, addr := c.addrAbsoluteX()

	c.Y = c.Bus.Read(uint16(addr))
	c.setZN(c.Y)

	return pageCrossed
}

// LDA $nnnn,X -> Opcode $BD
func (c *CPU) opLDA_ABSX() bool {
	pageCrossed, addr := c.addrAbsoluteX()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)

	return pageCrossed
}

// LDX $nnnn,Y -> Opcode $BE
func (c *CPU) opLDX_ABSY() bool {
	pageCrossed, addr := c.addrAbsoluteY()

	c.X = c.Bus.Read(uint16(addr))
	c.setZN(c.X)

	return pageCrossed
}

// NOP -> Opcode $EA
func (c *CPU) opNOP() bool {
	return false
}
