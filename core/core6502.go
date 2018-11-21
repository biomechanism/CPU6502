package core

import "fmt"

type Cpu struct {
	a    byte   //Accumulator
	x, y byte   //Index Registers
	p    byte   //Processor flags
	s    byte   //Stack pointer
	pc   uint16 //Program counter
	mem  []byte
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

func (cpu *Cpu) GetP() byte {
	return cpu.p
}

func (cpu *Cpu) GetS() byte {
	return cpu.s
}

func (cpu *Cpu) GetPC() uint16 {
	return cpu.pc
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
		cpu.p |= 1 << 7
	} else {
		cpu.p &= ^byte(1 << 6)
	}
}

func (cpu *Cpu) setOverflowStatus(val1, val2, result uint8) {
	v1 := int(val1)
	v2 := int(val2)
	r := int(result)

	//fmt.Printf("(Unsigned) V1: %d, V2: %d, R: %d\n", val1, val2, result)
	//fmt.Printf("(Signed) V1: %d, V2: %d, R: %d\n", v1, v2, r)

	if v1 >= 0 && v2 >= 0 && r < 0 {
		fmt.Println("1. Setting Overflow")
		cpu.p |= 1 << 6
	} else if v1 < 0 && v2 < 0 && r > 0 {
		fmt.Println("2. Setting Overflow")
		//TODO: Check if r should be >0 or >= 0
		cpu.p |= 1 << 6
	} else {
		fmt.Println("Clearing Overflow")
		cpu.p &= ^byte(1 << 6)
	}

}

func (cpu *Cpu) clearOverflowStatus() {
	cpu.p &= ^byte(1 << 6)
}
