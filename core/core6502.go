package core

type Cpu struct {
	a    uint8  //Accumulator
	x, y uint8  //Index Registers
	p    uint8  //Processor flags
	s    uint8  //Stack pointer
	pc   uint16 //Program counter
	mem  []uint8
}

//NewCPU instantiates a new instance of the Cpu
func NewCPU(memory []uint8) *Cpu {
	var cpu = Cpu{}
	cpu.mem = memory
	return &cpu
}

func (cpu *Cpu) GetA() uint8 {
	return cpu.a
}

func (cpu *Cpu) GetX() uint8 {
	return cpu.x
}

func (cpu *Cpu) GetY() uint8 {
	return cpu.y
}

func (cpu *Cpu) GetP() uint8 {
	return cpu.p
}

func (cpu *Cpu) GetS() uint8 {
	return cpu.s
}

func (cpu *Cpu) GetPC() uint16 {
	return cpu.pc
}
