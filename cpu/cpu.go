package cpu

import "perissinotto_nes/bus"

type CPU struct {
	A, X, Y byte
	PC      uint16
	SP      byte
	Status  byte

	Bus    *bus.Bus
	cycles uint64

	opcodes [256]Instruction
}

func NewCPU(bus *bus.Bus) *CPU {
	c := &CPU{
		Bus: bus,
	}
	c.Reset()
	c.initOpcodes()
	return c
}

func (c *CPU) Reset() {
	c.SP = 0xFD
	c.Status = 0x34

	lo := c.Bus.Read(0xFFFC)
	hi := c.Bus.Read(0xFFFD)
	c.PC = uint16(hi)<<8 | uint16(lo)

	c.cycles = 7
}

func (c *CPU) Step() {
	opcode := c.Bus.Read(c.PC)
	c.PC++

	inst := c.opcodes[opcode]

	extraCycle := inst.exec()

	c.cycles += uint64(inst.cycles)

	if extraCycle {
		c.cycles++
	}
}
