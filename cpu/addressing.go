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

func (c *CPU) addrAbsoluteX() uint16 {
	lowByte := c.Bus.Read(c.PC)
	highByte := c.Bus.Read(c.PC + 1)
	c.PC += 2
	return (uint16(highByte)<<8 | uint16(lowByte)) + uint16(c.X)
}
