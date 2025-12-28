package bus

type Bus struct {
	Mem [65536]byte
}

func (b *Bus) Read(addr uint16) byte {
	return b.Mem[addr]
}

func (b *Bus) Write(addr uint16, value byte) {
	b.Mem[addr] = value
}
