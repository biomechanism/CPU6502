package core

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
	if value < 0 {
		cpu.p |= 1 << 7
	}
}
