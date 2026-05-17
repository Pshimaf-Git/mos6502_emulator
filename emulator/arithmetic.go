package emulator

func (cpu *CPU) adc(mode modeFunc) {
	addr := mode()
	data := cpu.Read(addr)
	cpu.addWithCarry(data)
}

func (cpu *CPU) sbc(mode modeFunc) {
	addr := mode()
	data := cpu.Read(addr)
	invertedData := ^data
	cpu.addWithCarry(invertedData)
}

// ISB: Increment memory by 1, then subtract the new value from accumulator (SBC)
// NON DOC
func (cpu *CPU) isb(mode modeFunc) {
	addr := mode()

	old := cpu.Read(addr)

	newVal := old + 1
	cpu.Write(addr, newVal)

	inverted := ^newVal
	cpu.addWithCarry(inverted)
}

// INC: Increment
func (cpu *CPU) inc(mode modeFunc) {
	addr := mode()
	val := cpu.Read(addr) + 1
	cpu.Write(addr, val)
	cpu.updateZN(val)
}

// DEC: Decrement
func (cpu *CPU) dec(mode modeFunc) {
	addr := mode()
	val := cpu.Read(addr) - 1
	cpu.Write(addr, val)
	cpu.updateZN(val)
}

func (cpu *CPU) dex(mode modeFunc) {
	cpu.Regs.X--
	cpu.updateZN(cpu.Regs.X)
}

func (cpu *CPU) dey(mode modeFunc) {
	cpu.Regs.Y--
	cpu.updateZN(cpu.Regs.Y)
}

func (cpu *CPU) inx(mode modeFunc) {
	cpu.Regs.X++
	cpu.updateZN(cpu.Regs.X)
}

func (cpu *CPU) iny(mode modeFunc) {
	cpu.Regs.Y++
	cpu.updateZN(cpu.Regs.Y)
}

// Helper
func (cpu *CPU) addWithCarry(data byte) {
	carry := uint16(0)
	if cpu.GetFlag(FlagCarry) {
		carry = 1
	}
	a := uint16(cpu.Regs.Accumulator)
	b := uint16(data)
	sum := a + b + carry
	result := byte(sum & 0xFF)

	cpu.setStatusFlag(FlagCarry, sum > 0xFF)

	operandsSameSign := ((a ^ b) & 0x80) == 0
	resultSignDiffers := ((a ^ uint16(result)) & 0x80) != 0
	cpu.setStatusFlag(FlagOverflow, operandsSameSign && resultSignDiffers)

	cpu.Regs.Accumulator = result
	cpu.updateZN(cpu.Regs.Accumulator)
}
