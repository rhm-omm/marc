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
		v, ok := fldMap[e.tag]
		// If no value for tag, make one and add it
		if !ok {
			ea := make([]entry, 0)
			ea = append(ea, e)
			fldMap[e.tag] = ea
		} else {
			v = append(v, e)
			fldMap[e.tag] = v
		}
		dir = append(dir, e)
	}
	return dir
}

var fldMap = make(map[int][]entry)

func (d Directory) entryFor(tag int) []entry {
	return fldMap[tag]
}

func (d Directory) FldLen(tag int) []int {
	entries := d.entryFor(tag)
	var lenarr = make([]int, len(entries))
	if len(entries) > 0 {
		for i := 0; i < len(entries); i++ {
			lenarr[i] = entries[i].fldlen
		}
	}
	return lenarr
}

func (d Directory) FldOfs(tag int) []int {
	entries := d.entryFor(tag)
	var ofsarr = make([]int, len(entries))
	if len(entries) > 0 {
		for i := 0; i < len(entries); i++ {
			ofsarr[i] = entries[i].fldofs
		}
	}
	return ofsarr
}
