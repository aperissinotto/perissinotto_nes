package cpu

const (
	CARRY     byte = 1 << 0
	ZERO      byte = 1 << 1
	INTERRUPT byte = 1 << 2
	DECIMAL   byte = 1 << 3
	BREAK     byte = 1 << 4
	UNUSED    byte = 1 << 5
	OVERFLOW  byte = 1 << 6
	NEGATIVE  byte = 1 << 7
)

func (c *CPU) setFlag(flag byte, value bool) {
	if value {
		c.Status |= flag
	} else {
		c.Status &^= flag
	}
}

func (c *CPU) setZN(value byte) {
	c.setFlag(ZERO, value == 0)
	c.setFlag(NEGATIVE, value&0x80 != 0)
}
