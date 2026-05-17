package emulator

func (cpu *CPU) asl(mode modeFunc) {
	addr := mode()
	val := cpu.Read(addr)
	cpu.setStatusFlag(FlagCarry, (val&0x80) != 0)
	val <<= 1
	cpu.Write(addr, val)
	cpu.updateZN(val)
}

func (cpu *CPU) aslAcc(mode modeFunc) {
	cpu.setStatusFlag(FlagCarry, (cpu.Regs.Accumulator&0x80) != 0)
	cpu.Regs.Accumulator <<= 1
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) lsr(mode modeFunc) {
	addr := mode()
	val := cpu.Read(addr)
	cpu.setStatusFlag(FlagCarry, (val&1) != 0)
	val >>= 1
	cpu.Write(addr, val)
	cpu.updateZN(val)
}

func (cpu *CPU) lsrAcc(mode func() Word) {
	cpu.setStatusFlag(FlagCarry, (cpu.Regs.Accumulator&0x01) != 0)
	cpu.Regs.Accumulator >>= 1
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) rol(mode modeFunc) {
	addr := mode()
	val := cpu.Read(addr)
	oldCarry := uint8(0)
	if cpu.GetFlag(FlagCarry) {
		oldCarry = 1
	}
	cpu.setStatusFlag(FlagCarry, (val&0x80) != 0)
	val = (val << 1) | oldCarry
	cpu.Write(addr, val)
	cpu.updateZN(val)
}

func (cpu *CPU) rolAcc(_ modeFunc) {
	oldCarry := uint8(0)
	if cpu.GetFlag(FlagCarry) {
		oldCarry = 1
	}
	cpu.setStatusFlag(FlagCarry, (cpu.Regs.Accumulator&0x80) != 0)
	cpu.Regs.Accumulator = (cpu.Regs.Accumulator << 1) | oldCarry
	cpu.updateZN(cpu.Regs.Accumulator)
}

func (cpu *CPU) ror(mode modeFunc) {
	addr := mode()
	val := cpu.Read(addr)
	oldCarry := uint8(0)
	if cpu.GetFlag(FlagCarry) {
		oldCarry = 0x80
	}
	cpu.setStatusFlag(FlagCarry, (val&0x01) != 0)
	val = (val >> 1) | oldCarry
	cpu.Write(addr, val)
	cpu.updateZN(val)
}

func (cpu *CPU) rorAcc(_ modeFunc) {
	oldCarry := uint8(0)
	if cpu.GetFlag(FlagCarry) {
		oldCarry = 0x80
	}
	cpu.setStatusFlag(FlagCarry, (cpu.Regs.Accumulator&0x01) != 0)
	cpu.Regs.Accumulator = (cpu.Regs.Accumulator >> 1) | oldCarry
	cpu.updateZN(cpu.Regs.Accumulator)
}
