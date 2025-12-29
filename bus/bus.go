package bus

type Bus struct {
	ram [65536]byte
}

func (b *Bus) Read(addr uint16) byte {
	return b.ram[addr]
}

func (b *Bus) Write(addr uint16, data byte) {
	b.ram[addr] = data
}
