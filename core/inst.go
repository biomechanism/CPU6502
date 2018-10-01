package core

func (cpu *Cpu) ADC() {

}

func (cpu *Cpu) AND() {
	opcode := cpu.mem[cpu.pc]
	switch opcode {
	case andImm:
		print("AND Immediate")
		v := cpu.mem[cpu.pc+1]
		cpu.a = cpu.a & v
	case andZp:
		print("AND Zero Page")
	case andZpX:
		print("AND Zero Page X")
	case andAbs:
		print("AND Absolute")
	case andAbsX:
		print("AND Absolute X")
	case andAbsY:
		print("AND Absolute Y")
	case andIndX:
		print("AND Indirect X")
	case andIndY:
		print("AND Indirect Y")
	}
}
