package directory

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rhm-omm/marc/ldr"
)

type Entry struct {
	Tag    int
	Fldlen int
	Fldofs int
}

type Directory []Entry

func DirFrom(rawMARC []byte) Directory {
	// Get base address from leader
	l := ldr.LdrFrom(rawMARC)
	baseAddr := l.BaseAddr()
	// Create the slice that contains the directory
	dirSrc := rawMARC[24:baseAddr]
	// Create Directory vbl
	var dir Directory
	// Loop through the bytes, filling out the entry struct
	for i := 0; i < len(dirSrc)-1; i += 12 {
		var e Entry
		t, err := strconv.Atoi(string(dirSrc[i : i+3]))
		if err != nil {
			fmt.Println("Tag not a number")
			os.Exit(1)
		}
		e.Tag = t
		l, err := strconv.Atoi(string(dirSrc[i+3 : i+7]))
		if err != nil {
			fmt.Println("Length not a number")
			os.Exit(1)
		}
		e.Fldlen = l
		s, err := strconv.Atoi(string(dirSrc[i+7 : i+12]))
		if err != nil {
			fmt.Println("Starting position not a number")
			os.Exit(1)
		}
		e.Fldofs = s
		// Update map
		v, ok := fldMap[e.Tag]
		// If no value for tag, make one and add it
		if !ok {
			ea := make([]Entry, 0)
			ea = append(ea, e)
			fldMap[e.Tag] = ea
		} else {
			v = append(v, e)
			fldMap[e.Tag] = v
		}
		dir = append(dir, e)
	}
	return dir
}

var fldMap = make(map[int][]Entry)

func (d Directory) EntryFor(tag int) []Entry {
	return fldMap[tag]
}

func (d Directory) Fldlen(tag int) []int {
	entries := d.EntryFor(tag)
	var lenarr = make([]int, len(entries))
	if len(entries) > 0 {
		for i := 0; i < len(entries); i++ {
			lenarr[i] = entries[i].Fldlen
		}
	}
	return lenarr
}

func (d Directory) FldOfs(tag int) []int {
	entries := d.EntryFor(tag)
	var ofsarr = make([]int, len(entries))
	if len(entries) > 0 {
		for i := 0; i < len(entries); i++ {
			ofsarr[i] = entries[i].Fldofs
		}
	}
	return ofsarr
}
