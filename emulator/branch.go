package emulator

// JMP: Jump at address
func (cpu *CPU) jmp(mode modeFunc) {
	addr := mode()
	cpu.Regs.PC = addr
}

func (cpu *CPU) jsr(mode func() Word) {
	addr := mode()
	retAddr := cpu.Regs.PC - 1
	cpu.Push(byte(retAddr >> 8))
	cpu.Push(byte(retAddr & 0xFF))
	cpu.Regs.PC = addr
}

func (cpu *CPU) rts(mode func() Word) {
	low := Word(cpu.Pop())
	high := Word(cpu.Pop())
	cpu.Regs.PC = (high<<8 | low) + 1
}

func (cpu *CPU) rti(mode modeFunc) {
	status := cpu.Pop()
	cpu.Regs.P = (status & 0xEF) | 0x20
	low := Word(cpu.Pop())
	high := Word(cpu.Pop())
	cpu.Regs.PC = (high << 8) | low
}

func (cpu *CPU) bne(mode func() Word) {
	addr := mode()
	if !cpu.GetFlag(FlagZero) {
		cpu.Regs.PC = addr
	}
}

func (cpu *CPU) beq(mode func() Word) {
	addr := mode()
	if cpu.GetFlag(FlagZero) {
		cpu.Regs.PC = addr
	}
}

func (cpu *CPU) bvs(mode func() Word) {
	addr := mode()
	if cpu.GetFlag(FlagOverflow) {
		cpu.Regs.PC = addr
	}
}

func (cpu *CPU) bvc(mode modeFunc) {
	addr := mode()
	if !cpu.GetFlag(FlagOverflow) {
		cpu.Regs.PC = addr
	}
}

func (cpu *CPU) bcc(mode modeFunc) {
	addr := mode()
	if !cpu.GetFlag(FlagCarry) {
		cpu.Regs.PC = addr
	}
}

func (cpu *CPU) bcs(mode modeFunc) {
	addr := mode()
	if cpu.GetFlag(FlagCarry) {
		cpu.Regs.PC = addr
	}
}

func (cpu *CPU) bpl(mode modeFunc) {
	addr := mode()
	if !cpu.GetFlag(FlagNegative) {
		cpu.Regs.PC = addr
	}
}

func (cpu *CPU) bmi(mode modeFunc) {
	addr := mode()
	if cpu.GetFlag(FlagNegative) {
		cpu.Regs.PC = addr
	}
}
