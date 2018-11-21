package core

import "errors"

const (
	ExpectedOpcode = iota
	Size
	Cycles
)

var infoArray = [][3]int{
	{brk, 1, 7},
	{oraIndX, 2, 6},
	{fe02, 0, 0},
	{fe03, 0, 0},
	{fe04, 0, 0},
	{oraZp, 2, 3},
	{aslZp, 2, 5},
	{fe07, 0, 0},

	{php, 1, 3},
	{oraImm, 2, 2}, //One Extra cycle to be added if accessing across page boundary
	{aslAcc, 1, 2},
	{fe0b, 0, 0},
	{fe0c, 0, 0},
	{oraAbs, 3, 4},
	{aslAbs, 3, 6},
	{fe0f, 0, 0},

	{bpl, 2, 2}, //+1 Cycle if branch to same page; +2 if branch to different page
	{oraIndY, 2, 5},
	{fe12, 0, 0},
	{fe13, 0, 0},
	{fe14, 0, 0},
	{oraZpX, 2, 4},
	{aslZpX, 2, 6},
	{fe17, 0, 0},

	{clc, 1, 2},
	{oraAbsY, 3, 4}, //+1 cycle for crossing page boundary
	{fe1a, 0, 0},
	{fe1b, 0, 0},
	{fe1c, 0, 0},
	{oraAbsX, 3, 4}, //+1 cycle for crossing page boundary
	{aslAbsX, 3, 7},
	{fe1f, 0, 0},

	{jsr, 3, 6},
	{andIndX, 2, 6},
	{fe22, 0, 0},
	{fe23, 0, 0},
	{bitZp, 2, 3},
	{andZp, 2, 3},
	{rolZp, 2, 5},
	{fe27, 0, 0},

	{plp, 1, 4},
	{andImm, 2, 2},
	{rolAcc, 1, 2},
	{fe2b, 0, 0},
	{bitAbs, 3, 4},
	{andAbs, 3, 4},
	{rolAbs, 3, 6},
	{fe2f, 0, 0},

	{bmi, 2, 2}, //+1 if branch to same page, +2 otherwise
	{andIndY, 2, 5},
	{fe32, 0, 0},
	{fe33, 0, 0},
	{fe34, 0, 0},
	{andZpX, 2, 4},
	{rolZpX, 2, 6},
	{fe37, 0, 0},

	{sec, 1, 2},
	{andAbsY, 3, 4}, //+1 if page boundary is crossed.
	{fe3a, 0, 0},
	{fe3b, 0, 0},
	{fe3c, 0, 0},
	{andAbsX, 3, 4}, //+1 if page boundary is crossed
	{rolAbsX, 3, 7},
	{fe3f, 0, 0},

	{rti, 1, 6},
	{eorIndX, 2, 6},
	{fe42, 0, 0},
	{fe43, 0, 0},
	{fe44, 0, 0},
	{eorZp, 2, 3},
	{lsrZp, 2, 5},
	{fe47, 0, 0},

	{pha, 1, 3},
	{eorImm, 2, 2},
	{lsrAcc, 1, 2},
	{fe4b, 0, 0},
	{jmpAbs, 3, 3},
	{eorAbs, 3, 4},
	{lsrAbs, 3, 6},
	{fe4f, 0, 0},

	{bvc, 2, 2},     // +1 if branch to same page, +2 otherwise
	{eorIndY, 2, 5}, // +1 if page boundary is crossed
	{fe52, 0, 0},
	{fe53, 0, 0},
	{fe54, 0, 0},
	{eorZpX, 2, 4},
	{lsrZpX, 2, 6},
	{fe57, 0, 0},

	{cli, 1, 2},
	{eorAbsY, 3, 4}, // +1 if page boundary crossed
	{fe5a, 0, 0},
	{fe5b, 0, 0},
	{fe5c, 0, 0},
	{eorAbsX, 3, 4}, // +1 if page boundary crossed
	{lsrAbsX, 3, 7},
	{fe5f, 0, 0},

	{rts, 1, 6},
	{adcIndX, 2, 6},
	{fe62, 0, 0},
	{fe63, 0, 0},
	{fe64, 0, 0},
	{adcZp, 2, 3},
	{rorZp, 2, 5},
	{fe67, 0, 0},

	{pla, 1, 4},
	{adcImm, 2, 2},
	{rorAcc, 1, 2},
	{fe6b, 0, 0},
	{jmpInd, 3, 5},
	{adcAbs, 3, 4},
	{rorAbs, 3, 6},
	{fe6f, 0, 0},

	{bvs, 2, 2}, // +1 if branch to same a page, +2 otherwise
	{adcIndY, 2, 6},
	{fe72, 0, 0},
	{fe73, 0, 0},
	{fe74, 0, 0},
	{adcZpX, 2, 4},
	{rorZpX, 2, 6},
	{fe77, 0, 0},

	{sei, 1, 2},
	{adcAbsY, 3, 4}, // +1 if page boundary is crossed
	{fe7a, 0, 0},
	{fe7b, 0, 0},
	{fe7c, 0, 0},
	{adcAbsX, 3, 4}, //+1 if page boundary crossed
	{rorAbsX, 3, 7},
	{fe7f, 0, 0},

	{fe80, 0, 0},
	{staIndX, 2, 6},
	{fe82, 0, 0},
	{fe83, 0, 0},
	{styZp, 2, 3},
	{staZp, 2, 3},
	{stxZp, 2, 3},
	{fe87, 0, 0},

	{dey, 1, 2},
	{fe89, 0, 0},
	{txa, 1, 2},
	{fe8b, 0, 0},
	{styAbs, 3, 4},
	{staAbs, 3, 4},
	{stxAbs, 3, 4},
	{fe8f, 0, 0},

	{bcc, 2, 2}, // +1 if branch to same page, +2 otherwise
	{staIndY, 2, 6},
	{fe92, 0, 0},
	{fe93, 0, 0},
	{styZpX, 2, 4},
	{staZpX, 2, 4},
	{stxZpY, 2, 4},
	{fe97, 0, 0},

	{tya, 1, 2},
	{staAbsY, 3, 5},
	{txs, 1, 2},
	{fe9b, 0, 0},
	{fe9c, 0, 0},
	{staAbsX, 3, 5},
	{fe9e, 0, 0},
	{fe9f, 0, 0},

	{ldyImm, 2, 2},
	{ldaIndX, 2, 6},
	{ldxImm, 2, 2},
	{fea3, 0, 0},
	{ldyZp, 2, 3},
	{ldaZp, 2, 3},
	{ldxZp, 2, 3},
	{fea7, 0, 0},

	{tay, 1, 2},
	{ldaImm, 2, 2},
	{tax, 1, 2},
	{feab, 0, 0},
	{ldyAbs, 3, 4},
	{ldaAbs, 3, 4},
	{ldxAbs, 3, 4},
	{feaf, 0, 0},

	{bcs, 2, 2},     // +1 if branch to same page, +2 otherwise
	{ldaIndY, 2, 5}, // +1 if page boundary is crossed
	{feb2, 0, 0},
	{feb3, 0, 0},
	{ldyZpX, 2, 4},
	{ldaZpX, 2, 4},
	{ldxZpY, 2, 4},
	{feb7, 0, 0},

	{clv, 1, 2},
	{ldaAbsY, 3, 4}, // +1 if page boundary is crossed.
	{tsx, 1, 2},
	{febb, 0, 0},
	{ldyAbsX, 3, 4}, // +1 if page boundary is crossed.
	{ldaAbsX, 3, 4}, // +1 if page boundary is crossed.
	{ldxAbsY, 3, 4}, // +1 if page boundary is crossed.
	{febf, 0, 0},

	{cpyImm, 2, 2},
	{cmpIndX, 2, 6},
	{fec2, 0, 0},
	{fec3, 0, 0},
	{cpyZp, 2, 3},
	{cmpZp, 2, 4},
	{decZp, 2, 5},
	{fec7, 0, 0},

	{iny, 1, 2},
	{cmpImm, 2, 2},
	{dex, 1, 2},
	{fecb, 0, 0},
	{cpyAbs, 3, 4},
	{cmpAbs, 3, 4},
	{decAbs, 3, 6},
	{fecf, 0, 0},

	{bne, 2, 2}, // +1 if branch to same page, +2 otherwise
	{cmpIndY, 2, 6},
	{fed2, 0, 0},
	{fed3, 0, 0},
	{fed4, 0, 0},
	{cmpZpX, 2, 4},
	{decZpX, 2, 6},
	{fed7, 0, 0},

	{cld, 1, 2},
	{cmpAbsY, 3, 4}, // +1 if page boundary is crossed
	{feda, 0, 0},
	{fedb, 0, 0},
	{fedc, 0, 0},
	{cmpAbsX, 3, 4}, // +1 if page boundary crossed
	{decAbsX, 3, 7},
	{fedf, 0, 0},

	{cpxImm, 2, 2},
	{sbcIndX, 2, 6},
	{fee2, 0, 0},
	{fee3, 0, 0},
	{cpxZp, 2, 3},
	{sbcZp, 2, 3},
	{incZp, 2, 5},
	{fee7, 0, 0},

	{inx, 1, 2},
	{sbcImm, 2, 2},
	{nop, 1, 2},
	{feeb, 0, 0},
	{cpxAbs, 3, 4},
	{sbcAbs, 3, 4},
	{incAbs, 3, 6},
	{feef, 0, 0},

	{beq, 2, 2},     // +1 for branch to same page, +2 otherwise
	{sbcIndY, 2, 5}, // +1 if page boundary is crossed.
	{fef2, 0, 0},
	{fef3, 0, 0},
	{fef4, 0, 0},
	{sbcZpX, 2, 4},
	{incZpX, 2, 6},
	{fef7, 0, 0},

	{sed, 1, 2},
	{sbcAbsY, 3, 4}, // +1 if page boundary crossed
	{fefa, 0, 0},
	{fefb, 0, 0},
	{fefc, 0, 0},
	{sbcAbsX, 3, 4}, // +1 if page boundary crossed.
	{incAbsX, 3, 7},
	{feff, 0, 0},
}

func InstInfo(opcode int) (size, cycles int, err error) {

	data := infoArray[opcode]

	if opcode == data[ExpectedOpcode] {
		return data[Size], data[Cycles], nil
	}

	return 0, 0, errors.New("The opcode and size/cycle lookup are out of sync")

}
