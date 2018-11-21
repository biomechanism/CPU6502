package core

func (cpu *Cpu) ADC() {

}

func (cpu *Cpu) AND() {
	opcode := cpu.mem[cpu.pc]
	switch opcode {
	case andImm:
		print("AND Immediate\n")
		v := cpu.mem[cpu.pc+1]
		cpu.a = cpu.a & v
		cpu.SetNZStatus(cpu.a)
	case andZp:
		//Zero Page, addressable $0000 - $00FF
		print("AND Zero Page\n")
		v := cpu.mem[cpu.pc+1]
		cpu.a = cpu.a & cpu.mem[v]
	case andZpX:
		print("AND Zero Page X\n")
		v := cpu.mem[cpu.pc+1]
		base := v + cpu.x
		cpu.a = cpu.a & base
	case andAbs:
		print("AND Absolute\n")
		v1 := cpu.mem[cpu.pc+1]
		v2 := cpu.mem[cpu.pc+2]
		var addr uint16
		addr = uint16(v2)
		addr = addr << 8
		addr = addr | uint16(v1)
		v := cpu.mem[addr]
		cpu.a = cpu.a & v
	case andAbsX:
		print("AND Absolute X\n")
		v1 := cpu.mem[cpu.pc+1]
		v2 := cpu.mem[cpu.pc+2]
		var addr uint16
		addr = uint16(v2) << 8
		addr |= uint16(v1)
		addr += uint16(cpu.x)
		v := cpu.mem[addr]
		cpu.a &= v
	case andAbsY:
		print("AND Absolute Y\n")
		v1 := cpu.mem[cpu.pc+1]
		v2 := cpu.mem[cpu.pc+2]
		var addr uint16
		addr = uint16(v2) << 8
		addr |= uint16(v1)
		addr += uint16(cpu.y)
		v := cpu.mem[addr]
		cpu.a &= v
	case andIndX:
		print("AND Indirect X\n")
		zpIndex := cpu.mem[cpu.pc+1]
		zpIndex += cpu.x
		lowByte := cpu.mem[zpIndex]
		hiByte := cpu.mem[zpIndex+1]
		var addr = (uint16(hiByte) << 8) | uint16(lowByte)
		v := cpu.mem[addr]
		cpu.a &= v
	case andIndY:
		print("AND Indirect Y\n")
		var zpAddr uint16
		v1 := cpu.mem[cpu.pc+1]
		zpAddr = uint16(v1)
		lowByte := cpu.mem[zpAddr]
		hiByte := cpu.mem[zpAddr+1]
		var addr = (uint16(hiByte) << 8) | uint16(lowByte)
		addr += uint16(cpu.y)
		v := cpu.mem[addr]
		cpu.a &= v
	}
}
