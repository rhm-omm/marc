package directory

import ""github.com/rhm-omm/marc/ldr""

type entry struct {
	tag      int
	fldlen   int
	startpos int
}

type Directory []entry

func DirFrom(MARCrec []byte) Directory {
	// Get base address from leader
	l := ldr.LdrFrom(MARCrec)
	// Create the slice that defines the directory
	// Create Directory vbl
	// Loop through the bytes, filling out the entry struct
	// Append entry to directory, and repeat
	// Return directory
}

var fldMap = make(map[int]entry)

func entryFor(tag int) entry {
	return fldMap[tag]
}

func FldLen(tag int) int {
	entry := entryFor(tag)
	return entry.fldlen
}

func FldStart(tag int) int {
	entry := entryFor(tag)
	return entry.startpos
}
