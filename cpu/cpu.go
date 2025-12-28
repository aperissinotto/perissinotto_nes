package cpu

import "perissinotto_nes/bus"

type CPU struct {
	A  byte
	X  byte
	Y  byte
	SP byte
	PC uint16

	Bus *bus.Bus
}

func NewCPU(b *bus.Bus) *CPU {
	c := &CPU{
		Bus: b,
	}
	c.Reset()
	return c
}

func (c *CPU) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.SP = 0xFD
	c.PC = uint16(c.Bus.Read(0xFFFC)) | uint16(c.Bus.Read(0xFFFD))<<8
}

func (c *CPU) Step() {
	opcode := c.Bus.Read(c.PC)
	c.PC++

	switch opcode {
	case 0xA9: // LDA imediato
		value := c.Bus.Read(c.PC)
		c.PC++
		c.A = value
	}
}
