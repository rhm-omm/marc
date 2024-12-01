package directory

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rhm-omm/marc/ldr"
)

type entry struct {
	tag    int
	fldlen int
	fldofs int
}

type Directory []entry

func DirFrom(MARCrec []byte) Directory {
	// Get base address from leader
	l := ldr.LdrFrom(MARCrec)
	baseAddr := l.BaseAddr()
	// Create the slice that contains the directory
	dirSrc := MARCrec[24:baseAddr]
	//	fmt.Println(len(dirSrc))
	// Create Directory vbl
	var dir Directory
	// Loop through the bytes, filling out the entry struct
	for i := 0; i < len(dirSrc)-1; i += 12 {
		var e entry
		t, err := strconv.Atoi(string(dirSrc[i : i+3]))
		if err != nil {
			fmt.Println("Tag not a number")
			os.Exit(1)
		}
		e.tag = t
		l, err := strconv.Atoi(string(dirSrc[i+3 : i+7]))
		if err != nil {
			fmt.Println("Length not a number")
			os.Exit(1)
		}
		e.fldlen = l
		s, err := strconv.Atoi(string(dirSrc[i+7 : i+12]))
		if err != nil {
			fmt.Println("Starting position not a number")
			os.Exit(1)
		}
		e.fldofs = s
		// Update map
		fldMap[e.tag] = e
		// Add entry to directory, and repeat
		dir = append(dir, e)
	}
	return dir
}

var fldMap = make(map[int]entry)

func (d Directory) entryFor(tag int) entry {
	return fldMap[tag]
}

func (d Directory) FldLen(tag int) int {
	entry := d.entryFor(tag)
	return entry.fldlen
}

func (d Directory) FldOfs(tag int) int {
	entry := d.entryFor(tag)
	return entry.fldofs
}
