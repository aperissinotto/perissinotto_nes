package main

import (
	"fmt"
	"perissinotto_nes/bus"
	"perissinotto_nes/cpu"
)

func main() {
	b := &bus.Bus{}
	c := cpu.NewCPU(b)

	// Programa: LDA #$42
	//	b.ram[0x8000] = 0xA9
	//	b.ram[0x8001] = 0x42

	b.Write(0x8000, 0xA2)
	b.Write(0x8001, 0x10)
	b.Write(0x8002, 0xB5)
	b.Write(0x8003, 0x80)
	b.Write(0x0090, 0x55)
	b.Write(0x0100, 0x40)
	b.Write(0x8004, 0xAD)
	b.Write(0x8005, 0x00)
	b.Write(0x8006, 0x01)

	// Vetor de reset
	b.Write(0xFFFC, 0x00)
	b.Write(0xFFFD, 0x80)

	c.Reset()
	c.Step()
	c.Step()
	c.Step()

	fmt.Printf("Registrador A = %X\n", c.A)
}
