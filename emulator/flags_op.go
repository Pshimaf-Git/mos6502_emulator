package emulator

// SEC: Set flag carry on true
func (cpu *CPU) sec(_ modeFunc) {
	cpu.setStatusFlag(FlagCarry, true)
}

func (cpu *CPU) clc(_ modeFunc) {
	cpu.setStatusFlag(FlagCarry, false)
}

func (cpu *CPU) sed(_ modeFunc) {
	cpu.setStatusFlag(FlagDecimal, true)
}

func (cpu *CPU) cld(_ modeFunc) {
	cpu.setStatusFlag(FlagDecimal, false)
}

func (cpu *CPU) sei(_ modeFunc) {
	cpu.setStatusFlag(FlagInterrupt, true)
}

func (cpu *CPU) cli(_ modeFunc) {
	cpu.setStatusFlag(FlagInterrupt, false)
}

func (cpu *CPU) clv(_ modeFunc) {
	cpu.setStatusFlag(FlagOverflow, false)
}
