package core

import (
	"fmt"
	"testing"
)

func init() {

}

func TestLoad(t *testing.T) {

	cpu := NewCPU(make([]uint8, 1024*16))

	if cpu.A() == 0 {
		fmt.Println("Moo")
	}

}
