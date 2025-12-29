package cpu

func (c *CPU) opNOP() {
	c.cycles++
}

// LDA #$nn
func (c *CPU) opLDA_IMM() {
	value := c.Bus.Read(c.PC)
	c.PC++

	c.A = value
	c.setZN(c.A)
	c.cycles += 2
}

// LDA $nn
func (c *CPU) opLDA_ZP() {
	addr := c.addrZeroPage()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	c.cycles += 3
}

// LDA $nn,X
func (c *CPU) opLDA_ZPX() {
	addr := c.addrZeroPageX()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	c.cycles += 3
}

func (c *CPU) opLDA_ABS() {
	addr := c.addrAbsolute()

	c.A = c.Bus.Read(uint16(addr))
	c.setZN(c.A)
	c.cycles += 3
}

// LDX #$nn
func (c *CPU) opLDX_IMM() {
	value := c.Bus.Read(c.PC)
	c.PC++

	c.X = value
	c.setZN(c.X)
	c.cycles += 2
}

// LDX $nn
func (c *CPU) opLDX_ZP() {
	addr := c.addrZeroPage()

	c.X = c.Bus.Read(uint16(addr))
	c.setZN(c.X)
	c.cycles += 3
}
