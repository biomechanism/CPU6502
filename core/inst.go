package core

func (cpu *Cpu) addWithCarry(val1, val2 byte) byte {
	result := val1 + val2

	if result < val1 {
		cpu.p |= 1
	} else {
		cpu.p &= ^byte(1)
	}

	cpu.setOverflowStatus(val1, val2, result)

	return result
}

func (cpu *Cpu) ADC() {
	opcode := cpu.mem[cpu.pc]
	var v byte
	switch opcode {
	case adcImm:
		v = cpu.readImm(cpu.pc + 1)
	case adcZp:
		v = cpu.readZp(cpu.pc + 1)
	case adcZpX:
		//TODO: CHECK, Not sure whether the carry needs to be handled when adding the X index
		//to the base or not.
		v = cpu.readZpX(cpu.pc + 1)
	case adcAbs:
		v = cpu.readAbs(cpu.pc + 1)
	case adcAbsX:
		//TODO: CHECK, should addredd calculations from and index have a carry
		//check?
		v = cpu.readAbsX(cpu.pc + 1)
	case adcAbsY:
		v = cpu.readAbsY(cpu.pc + 1)
	case adcIndX:
		v = cpu.readIndX(cpu.pc + 1)
	case adcIndY:
		v = cpu.readIndY(cpu.pc + 1)
	}

	cpu.a = cpu.addWithCarry(cpu.addWithCarry(v, cpu.getCarry()), cpu.a)
	cpu.setNegativeStatus(cpu.a)

}

func (cpu *Cpu) AND() {
	opcode := cpu.mem[cpu.pc]
	var v byte
	switch opcode {
	case andImm:
		v = cpu.readImm(cpu.pc + 1)
	case andZp:
		v = cpu.readZp(cpu.pc + 1)
	case andZpX:
		v = cpu.readZpX(cpu.pc + 1)
	case andAbs:
		v = cpu.readAbs(cpu.pc + 1)
	case andAbsX:
		v = cpu.readAbsX(cpu.pc + 1)
	case andAbsY:
		v = cpu.readAbsY(cpu.pc + 1)
	case andIndX:
		v = cpu.readIndX(cpu.pc + 1)
	case andIndY:
		v = cpu.readIndY(cpu.pc + 1)
	}

	cpu.a &= v
	cpu.setNegativeStatus(cpu.a)
	cpu.setZeroStatus(cpu.a)
}

func (cpu *Cpu) ASL() {
	
}
