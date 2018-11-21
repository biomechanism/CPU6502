package core

func (cpu *Cpu) Decode() func() {
	opcode := cpu.mem[cpu.pc]
	switch opcode {
	case adcAbs, adcAbsX, adcAbsY, adcImm, adcIndX, adcIndY, adcZp, adcZpX:
		return func() {
			executor(cpu.ADC, cpu)
		}
	case andAbs, andAbsX, andAbsY, andImm, andIndX, andIndY, andZp, andZpX:
		return func() {
			executor(cpu.AND, cpu)
		}
	default:
		//Replace with proper function
		return nil
	}

}

func executor(fn func(), cpu *Cpu) {
	opCode := cpu.mem[cpu.pc]
	opSize := infoArray[opCode][Size]
	fn()
	cpu.pc += uint16(opSize)
}
