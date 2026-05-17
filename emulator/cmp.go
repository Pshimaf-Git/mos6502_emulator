package emulator

func (cpu *CPU) cmp(mode func() Word) {
	cpu.compare(cpu.Regs.Accumulator, mode)
}

func (cpu *CPU) cpx(mode func() Word) {
	cpu.compare(cpu.Regs.X, mode)
}

func (cpu *CPU) cpy(mode func() Word) {
	cpu.compare(cpu.Regs.Y, mode)
}

func (cpu *CPU) compare(reg byte, mode func() Word) {
	addr := mode()
	data := cpu.Read(addr)
	cpu.setStatusFlag(FlagCarry, reg >= data)
	cpu.updateZN(reg - data)
}
