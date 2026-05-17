package emulator

func (cpu *CPU) and(mode func() Word) {
	addr := mode()
	cpu.Regs.Accumulator &= cpu.Read(addr)
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) ora(mode func() Word) {
	addr := mode()
	cpu.Regs.Accumulator |= cpu.Read(addr)
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) eor(mode func() Word) {
	addr := mode()
	cpu.Regs.Accumulator ^= cpu.Read(addr)
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) bit(mode modeFunc) {
	addr := mode()
	data := cpu.Read(addr)
	res := cpu.Regs.Accumulator & data
	cpu.setStatusFlag(FlagZero, res == 0)
	cpu.setStatusFlag(FlagNegative, (data&0x80) != 0)
	cpu.setStatusFlag(FlagOverflow, (data&0x40) != 0)
}
