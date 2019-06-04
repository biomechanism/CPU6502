# 6502 CPU Emulator

### Status

The current emulator supports all official instructions. Support for the illegal opcodes has not yet been implemented, but are planned. 
As it is on its own it is not hughly useful, but the intention is for it to be part of a larger project. 
For the curious it may be useful to look into CPU emulation, to that extent I hope the code is mostly straightforward enough to understand. 
The intention however is not to produce an efficient emulator, rather something resonable to understand and functional. 
Many improvements could no doubt be made to its operation, as it is, it likely creates quite a bit of garbage for the GC to collect.

### Language

 - Go

### References

https://www.geeksforgeeks.org/check-for-integer-overflow/

http://www.classic-games.com/commodore64/64doc.html

http://www.6502.org/tutorials/vflag.html

https://stackoverflow.com/questions/19301498/carry-flag-auxiliary-flag-and-overflow-flag-in-assembly

http://graphics.stanford.edu/~seander/bithacks.html

https://www.masswerk.at/6502/6502_instruction_set.html#AND

http://www.emulator101.com/6502-addressing-modes.html

http://www.oxyron.de/html/opcodes02.html

https://retrocomputing.stackexchange.com/questions/145/why-does-6502-indexed-lda-take-an-extra-cycle-at-page-boundaries

https://en.wikibooks.org/wiki/6502_Assembly

http://www.6502.org/tutorials/compare_beyond.html

https://stackoverflow.com/questions/32917880/why-does-conditional-branching-in-asm-6502-have-limit-of-128-bytes

https://www.c64-wiki.com/wiki/Interrupt

https://www.c64-wiki.com/wiki/Reset_(Process)

https://www.pagetable.com/?p=410

https://wiki.nesdev.com/w/images/7/76/Programmanual.pdf

http://archive.6502.org/publications/dr_dobbs_journal_selected_articles/sbc_tsx_txs_instructions.pdf

https://skilldrick.github.io/easy6502/

