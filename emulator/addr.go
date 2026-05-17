package emulator

func (cpu *CPU) addrImplied() Word {
	return 0
}

func (cpu *CPU) addrRelative() Word {
	offset := int8(cpu.FetchByte())

	return Word(int32(cpu.Regs.PC) + int32(offset))
}

func (cpu *CPU) addrImmediate() Word {
	addr := cpu.Regs.PC
	cpu.incPC()
	return addr
}

func (cpu *CPU) addrZeroPage() Word {
	addr := Word(cpu.FetchByte())
	return addr
}

func (cpu *CPU) addrAbsolute() Word {
	return cpu.FetchWord()
}

func (cpu *CPU) addrAbsoluteX() Word {
	low := Word(cpu.Read(cpu.Regs.PC))
	high := Word(cpu.Read(cpu.Regs.PC + 1))
	cpu.incPC()
	cpu.incPC()
	return ((high << 8) | low) + Word(cpu.Regs.X)
}

func (cpu *CPU) addrZeroPageX() Word {
	addr := Word(cpu.FetchByte() + cpu.Regs.X)
	return addr & 0xFF
}

func (cpu *CPU) addrZeroPageY() Word {
	addr := Word(cpu.FetchByte() + cpu.Regs.Y)
	return addr & 0xFF
}

func (cpu *CPU) addrAbsoluteY() Word {
	low := Word(cpu.FetchByte())
	high := Word(cpu.FetchByte())
	return ((high << 8) | low) + Word(cpu.Regs.Y)
}

func (cpu *CPU) addrIndirect() Word {
	ptr := cpu.FetchWord()
	low := Word(cpu.Read(ptr))

	high := Word(cpu.Read((ptr & 0xFF00) | ((ptr + 1) & 0x00FF)))
	return (high << 8) | low
}

func (cpu *CPU) addrIndirectX() Word {
	zp := cpu.FetchByte()
	ptr := Word(zp+cpu.Regs.X) & 0x00FF
	low := Word(cpu.Read(ptr))
	high := Word(cpu.Read((ptr + 1) & 0x00FF))
	return (high << 8) | low
}

func (cpu *CPU) addrIndirectY() Word {
	zp := Word(cpu.FetchByte())
	low := Word(cpu.Read(zp))
	high := Word(cpu.Read((zp + 1) & 0x00FF))
	base := (high << 8) | low
	return base + Word(cpu.Regs.Y)
}
