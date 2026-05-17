package emulator

// LDA: Load Accumulator
func (cpu *CPU) lda(mode modeFunc) {
	addr := mode()
	cpu.Regs.Accumulator = cpu.Read(addr)
	cpu.updateZN(cpu.Regs.Accumulator)
}

// STA: Save accumulator at memory
func (cpu *CPU) sta(mode modeFunc) {
	addr := mode()
	cpu.Write(addr, cpu.Regs.Accumulator)
}

func (cpu *CPU) stx(mode modeFunc) {
	addr := mode()
	cpu.Write(addr, cpu.Regs.X)
}

func (cpu *CPU) sty(mode modeFunc) {
	addr := mode()
	cpu.Write(addr, cpu.Regs.Y)
}

func (cpu *CPU) ldx(mode modeFunc) {
	addr := mode()
	cpu.Regs.X = cpu.Read(addr)
	cpu.updateZN(cpu.Regs.X)
}

func (cpu *CPU) ldy(mode modeFunc) {
	addr := mode()
	cpu.Regs.Y = cpu.Read(addr)
	cpu.updateZN(cpu.Regs.Y)
}

func (cpu *CPU) tax(mode func() Word) {
	cpu.Regs.X = cpu.Regs.Accumulator
	cpu.updateZN(cpu.Regs.X)
}

func (cpu *CPU) tay(mode func() Word) {
	cpu.Regs.Y = cpu.Regs.Accumulator
	cpu.updateZN(cpu.Regs.Y)
}

func (cpu *CPU) txa(mode modeFunc) {
	cpu.Regs.Accumulator = cpu.Regs.X
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) tya(mode modeFunc) {
	cpu.Regs.Accumulator = cpu.Regs.Y
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) txs(_ modeFunc) {
	cpu.Regs.SP = cpu.Regs.X
}

func (cpu *CPU) tsx(_ modeFunc) {
	cpu.Regs.X = cpu.Regs.SP
	cpu.updateZN(cpu.Regs.X)
}
