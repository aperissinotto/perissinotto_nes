package cpu

import "perissinotto_nes/bus"

type CPU struct {
	A, X, Y byte
	PC      uint16
	SP      byte
	Status  byte

	Bus    *bus.Bus
	cycles int

	opcodeTable [256]func(*CPU)
}

func NewCPU(bus *bus.Bus) *CPU {
	c := &CPU{
		Bus: bus,
		SP:  0xFD,
	}
	c.Reset()
	c.initOpcodes()
	return c
}

func (c *CPU) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.SP = 0xFD
	c.Status = 0x24

	lo := c.Bus.Read(0xFFFC)
	hi := c.Bus.Read(0xFFFD)
	c.PC = uint16(hi)<<8 | uint16(lo)

	c.cycles = 7
}

func (c *CPU) Step() {
	opcode := c.Bus.Read(c.PC)
	c.PC++
	c.opcodeTable[opcode](c)
}
