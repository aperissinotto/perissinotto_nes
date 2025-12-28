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
	b.Mem[0x8000] = 0xA9
	b.Mem[0x8001] = 0x42

	// Vetor de reset
	b.Mem[0xFFFC] = 0x00
	b.Mem[0xFFFD] = 0x80

	c.Reset()
	c.Step()

	fmt.Printf("Registrador A = %X\n", c.A)
}
