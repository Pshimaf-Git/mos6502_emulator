package emulator

func (cpu *CPU) pha(_ modeFunc) {
	cpu.Push(cpu.Regs.Accumulator)
}

func (cpu *CPU) pla(_ modeFunc) {
	cpu.Regs.Accumulator = cpu.Pop()
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) php(mode modeFunc) {
	status := cpu.Regs.P | 0x30
	cpu.Push(status)
}

func (cpu *CPU) plp(mode modeFunc) {
	status := cpu.Pop()
	cpu.Regs.P = (status & 0xEF) | 0x20
}
