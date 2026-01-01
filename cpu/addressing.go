package cpu

func (c *CPU) addrZeroPage() uint16 {
	addr := c.Bus.Read(c.PC)
	c.PC++
	return uint16(addr)
}

func (c *CPU) addrZeroPageX() uint16 {
	addr := c.Bus.Read(c.PC)
	c.PC++
	addr = (addr + c.X) & 0xFF
	return uint16(addr)
}

func (c *CPU) addrAbsolute() uint16 {
	lowByte := c.Bus.Read(c.PC)
	highByte := c.Bus.Read(c.PC + 1)
	c.PC += 2
	return uint16(highByte)<<8 | uint16(lowByte)
}

func (c *CPU) addrAbsoluteX() (bool, uint16) {
	lowByte := c.Bus.Read(c.PC)
	highByte := c.Bus.Read(c.PC + 1)
	c.PC += 2
	base := uint16(highByte)<<8 | uint16(lowByte)
	addr := base + uint16(c.X)

	pageCrossed := (base & 0xFF00) != (addr & 0xFF00)

	return pageCrossed, addr
}

func (c *CPU) addrAbsoluteY() (bool, uint16) {
	lowByte := c.Bus.Read(c.PC)
	highByte := c.Bus.Read(c.PC + 1)
	c.PC += 2
	base := uint16(highByte)<<8 | uint16(lowByte)
	addr := base + uint16(c.Y)

	pageCrossed := (base & 0xFF00) != (addr & 0xFF00)

	return pageCrossed, addr
}

func (c *CPU) addrIndirectX() uint16 {
	addr := c.addrZeroPageX()
	lowByte := c.Bus.Read(addr)
	highByte := c.Bus.Read((addr + 1) & 0xFF)
	return uint16(highByte)<<8 | uint16(lowByte)
}

func (c *CPU) addrIndirectY() (bool, uint16) {
	oper := c.Bus.Read(c.PC)
	c.PC++
	lowByte := c.Bus.Read(uint16(oper))
	highByte := c.Bus.Read(uint16((oper + 1) & 0xFF))
	base := uint16(highByte)<<8 | uint16(lowByte)
	addr := base + uint16(c.Y)

	pageCrossed := (base & 0xFF00) != (addr & 0xFF00)

	return pageCrossed, addr
}
