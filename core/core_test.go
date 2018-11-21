package core

import (
	"fmt"
	"testing"
)

func init() {

}

func TestLoad(t *testing.T) {

}

func TestANDImmediate(t *testing.T) {
	cpu := newCpu()

	cpu.a = 2
	cpu.mem[0] = andImm
	cpu.mem[1] = 6

	inst := cpu.Decode()
	inst()

	fmt.Printf("OVERFLOW: %d\n", cpu.p&1<<6)

	if cpu.a != 2 {
		t.Errorf("Expexted %d, Actual %d\n", 2, cpu.a)
	}

	if cpu.pc != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.pc)
	}
}

func TestANDOverflowCleared(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10
	cpu.mem[0] = andImm
	cpu.mem[1] = 6

	cpu.p |= (1 << 6)

	inst := cpu.Decode()
	inst()

	overflow := cpu.p & (1 << 6)

	fmt.Printf("OVERFLOW: %d\n", overflow)

	if overflow != 0 {
		t.Errorf("Expexted %d, Actual %d\n", 0, overflow)
	}

}

func TestANDZeroPage(t *testing.T) {
	cpu := newCpu()

	pcStart := cpu.pc
	cpu.a = 3
	cpu.mem[0] = andZp
	cpu.mem[1] = 4
	cpu.mem[4] = 7
	inst := cpu.Decode()
	inst()

	if cpu.a != 3 {
		t.Errorf("Expexted %d, Actual %d\n", 3, cpu.a)
	}

	expected := uint16(infoArray[andZp][Size]) + pcStart

	if cpu.pc != expected {
		t.Errorf("Expected %d, Actual %d\n", expected, cpu.pc)
	}

}

func TestANDZeroPageX(t *testing.T) {
	cpu := newCpu()
	cpu.a = 9
	cpu.x = 2

	cpu.mem[0] = andZpX
	cpu.mem[1] = 8

	inst := cpu.Decode()
	inst()

	if cpu.a != 8 {
		t.Errorf("Expected %d, Actual %d\n", 8, cpu.a)
	}

}

func TestANDAbs(t *testing.T) {
	cpu := newCpu()
	cpu.a = 10

	cpu.mem[0] = andAbs
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03

	cpu.mem[1000] = 3

	inst := cpu.Decode()
	inst()

	if cpu.a != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.a)
	}

}

func TestANDAbsX(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = andAbsX
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03
	cpu.x = 2

	cpu.mem[1002] = 3

	inst := cpu.Decode()
	inst()

	if cpu.a != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.a)
	}

}

func TestANDAbsY(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = andAbsY
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03
	cpu.y = 2

	cpu.mem[1002] = 3

	inst := cpu.Decode()
	inst()

	if cpu.a != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.a)
	}
}

func TestANDIndX(t *testing.T) {

	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = andIndX
	cpu.mem[1] = 4 //Zero Page Index

	cpu.mem[4] = 0xe8
	cpu.mem[5] = 0x03

	cpu.mem[1000] = 3

	cpu.x = 0

	inst := cpu.Decode()
	inst()

	if cpu.a != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.a)
	}

}

func TestANDIndY(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10
	cpu.y = 1

	cpu.mem[0] = andIndY
	cpu.mem[1] = 4 //Zero Page Index

	cpu.mem[4] = 0xe8
	cpu.mem[5] = 0x03

	cpu.mem[1001] = 9

	inst := cpu.Decode()
	inst()

	if cpu.a != 8 {
		t.Errorf("Expected %d, Actual %d\n", 8, cpu.a)
	}

}

func newCpu() *Cpu {
	cpu := NewCPU(make([]byte, 1024*16))
	cpu.pc = 0
	return cpu
}
