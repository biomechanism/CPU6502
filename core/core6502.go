package core

import "fmt"

type Cpu struct {
	a                   byte   //Accumulator
	x, y                byte   //Index Registers
	p                   byte   //Processor flags
	n, v, b, d, i, z, c bool   //Status flags
	s                   byte   //Stack pointer
	pc                  uint16 //Program counter
	mem                 []byte
}

//NewCPU instantiates a new instance of the Cpu
func NewCPU(memory []byte) *Cpu {
	var cpu = Cpu{}
	cpu.mem = memory
	return &cpu
}

func (cpu *Cpu) GetA() byte {
	return cpu.a
}

func (cpu *Cpu) GetX() byte {
	return cpu.x
}

func (cpu *Cpu) GetY() byte {
	return cpu.y
}

// func (cpu *Cpu) GetP() byte {
// 	return cpu.p
// }

func (cpu *Cpu) GetS() byte {
	return cpu.s
}

func (cpu *Cpu) GetPC() uint16 {
	return cpu.pc
}

func (cpu *Cpu) SetCarry() {
	cpu.c = true
}

func (cpu *Cpu) ClearCarry() {
	cpu.c = false
}

func (cpu *Cpu) SetNZStatus(value byte) {
	// if int8(value) < 0 {
	// 	cpu.p |= 1 << 7
	// } else {
	// 	cpu.p &= ^byte(1 << 6)
	// }

}

func (cpu *Cpu) Execute() {
	//op := cpu.mem[cpu.pc]
	//Call op function

}

func (cpu *Cpu) setNegativeStatus(value byte) {
	if int8(value) < 0 {
		cpu.n = true
	} else {
		cpu.n = false
	}
}

func (cpu *Cpu) setZeroStatus(value byte) {
	if value == 0 {
		cpu.z = true
	} else {
		cpu.z = false
	}
}

func (cpu *Cpu) setOverflowStatus(val1, val2, result byte) {
	v1 := int8(val1)
	v2 := int8(val2)
	r := int8(result)

	fmt.Printf("(Unsigned) V1: %d, V2: %d, R: %d\n", val1, val2, result)
	fmt.Printf("(Signed) V1: %d, V2: %d, R: %d\n", v1, v2, r)

	if v1 >= 0 && v2 >= 0 && r < 0 {
		fmt.Println("1. Setting Overflow")
		cpu.v = true
	} else if v1 < 0 && v2 < 0 && r > 0 {
		fmt.Println("2. Setting Overflow")
		//TODO: Check if r should be >0 or >= 0
		cpu.v = true
	} else {
		fmt.Println("Clearing Overflow")
		cpu.v = false
	}

}

func (cpu *Cpu) setCarryStatus(val1, val2, result byte) {

	val3 := val1 + val2

	if val3 < val1 || val3 < val2 {
		fmt.Println("SETTING CARRY")
		cpu.c = true
	} else {
		cpu.c = false
	}

}

func (cpu *Cpu) clearOverflowStatus() {
	cpu.v = false
}

func (cpu *Cpu) isOverflow() bool {
	return cpu.v
}

func (cpu *Cpu) isCarry() bool {
	return cpu.c
}

func (cpu *Cpu) getCarry() byte {

	if cpu.isCarry() {
		return 1
	}

	return 0

}

func (cpu *Cpu) isNegative() bool {
	return cpu.n
}

func (cpu *Cpu) isZero() bool {
	return cpu.z
}

func (cpu *Cpu) readImm(loc uint16) byte {
	return cpu.mem[loc]
}

func (cpu *Cpu) readZp(loc uint16) byte {
	fmt.Printf("[ZeroPage] LOC: %d\n", loc)
	fmt.Printf("[ZeroPage] LOC Val: %d\n", cpu.mem[loc])
	v := cpu.mem[cpu.mem[loc]]
	fmt.Printf("[ZeroPage] Final Val: %d\n", v)
	return cpu.mem[cpu.mem[loc]]
}

func (cpu *Cpu) readZpX(loc uint16) byte {
	fmt.Printf("[ZeroPageX] LOC: %d\n", loc)
	fmt.Printf("[ZeroPageX] LOC Val: %d\n", cpu.mem[loc])
	v := cpu.addWithCarry(cpu.mem[loc], cpu.x)
	fmt.Printf("[ZeroPageX] Final Val: %d\n", cpu.mem[v])
	return cpu.mem[v]
}

func (cpu *Cpu) readAbs(loc uint16) byte {
	fmt.Printf("[Absolute] LOC: %d\n", loc)
	fmt.Printf("[Absolute] LOC Val: %d\n", cpu.mem[loc])
	v1 := cpu.mem[loc]
	v2 := cpu.mem[loc+1]
	var addr uint16
	addr = uint16(v2)
	addr = addr << 8
	addr = addr | uint16(v1)
	fmt.Printf("[ZeroPageX] Final Val: %d\n", cpu.mem[addr])
	return cpu.mem[addr]
}

func (cpu *Cpu) readAbsX(loc uint16) byte {
	v1 := cpu.mem[loc]
	v2 := cpu.mem[loc+1]
	var addr uint16
	addr = uint16(v2) << 8
	addr |= uint16(v1)
	addr += uint16(cpu.x)
	return cpu.mem[addr]
}

func (cpu *Cpu) readAbsY(loc uint16) byte {
	v1 := cpu.mem[loc]
	v2 := cpu.mem[loc+1]
	var addr uint16
	addr = uint16(v2) << 8
	addr |= uint16(v1)
	addr += uint16(cpu.y)
	return cpu.mem[addr]
}

func (cpu *Cpu) readIndX(loc uint16) byte {
	zpIndex := cpu.mem[loc]
	zpIndex += cpu.x
	lowByte := cpu.mem[zpIndex]
	hiByte := cpu.mem[zpIndex+1]
	var addr = (uint16(hiByte) << 8) | uint16(lowByte)
	return cpu.mem[addr]
}

func (cpu *Cpu) readIndY(loc uint16) byte {
	var zpAddr uint16
	v1 := cpu.mem[loc]
	zpAddr = uint16(v1)
	lowByte := cpu.mem[zpAddr]
	hiByte := cpu.mem[zpAddr+1]
	var addr = (uint16(hiByte) << 8) | uint16(lowByte)
	addr += uint16(cpu.y)
	return cpu.mem[addr]
}

func (cpu *Cpu) writeImm(loc uint16, value byte) {
	cpu.mem[loc] = value
}

func (cpu *Cpu) writeZp(loc uint16, value byte) {
	cpu.mem[cpu.mem[loc]] = value
}

//Possibly don't need Carry?
func (cpu *Cpu) writeZpX(loc uint16, value byte) {
	v := cpu.addWithCarry(cpu.mem[loc], cpu.x)
	cpu.mem[v] = value
}

func (cpu *Cpu) writeAbs(loc uint16, value byte) {
	v1 := cpu.mem[loc]
	v2 := cpu.mem[loc+1]
	var addr uint16
	addr = uint16(v2)
	addr = addr << 8
	addr = addr | uint16(v1)
	cpu.mem[addr] = value
}

func (cpu *Cpu) writeAbsX(loc uint16, value byte) {
	v1 := cpu.mem[loc]
	v2 := cpu.mem[loc+1]
	var addr uint16
	addr = uint16(v2) << 8
	addr |= uint16(v1)
	addr += uint16(cpu.x)
	cpu.mem[addr] = value
}

func (cpu *Cpu) pushStatustToStack() {
	//N 	V 	- 	B 	D 	I 	Z 	C
	reg := cpu.b2i(cpu.n)
	reg = reg << 1
	reg |= cpu.b2i(cpu.v)
	reg = reg << 2
	reg |= cpu.b2i(cpu.b)
	reg = reg << 1
	reg |= cpu.b2i(cpu.d)
	reg = reg << 1
	reg |= cpu.b2i(cpu.i)
	reg = reg << 1
	reg |= cpu.b2i(cpu.z)
	reg = reg << 1
	reg |= cpu.b2i(cpu.c)

	loc := uint16(01<<4 | cpu.s)
	cpu.writeZp(loc, reg)
	cpu.s--
}

func (cpu *Cpu) popStatusFromStack() {
	cpu.s++
	loc := uint16(01<<4 | cpu.s)
	val := cpu.readZp(loc)
	cpu.n = cpu.i2b(val & 0x80)
	cpu.v = cpu.i2b(val & 0x40)
	cpu.b = cpu.i2b(val & 0x20)
	cpu.d = cpu.i2b(val & 0x08)
	cpu.i = cpu.i2b(val & 0x04)
	cpu.z = cpu.i2b(val & 0x02)
	cpu.c = cpu.i2b(val & 0x01)
}

func (cpu *Cpu) push(val byte) {
	loc := uint16(01<<4 | cpu.s)
	cpu.writeZp(loc, val)
	cpu.s--
}

func (cpu *Cpu) pop() byte {
	cpu.s++
	loc := uint16(01<<4 | cpu.s)
	return cpu.readZp(loc)
}

func (cpu *Cpu) b2i(val bool) byte {
	if val == true {
		return 1
	}

	return 0
}

func (cpu *Cpu) i2b(val byte) bool {
	if val > 0 {
		return true
	}

	return false
}
