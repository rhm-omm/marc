package ctrlfld

import (
	"fmt"
	"os"

	"github.com/rhm-omm/marc/directory"
	"github.com/rhm-omm/marc/ldr"
)

// MARC control field. Lacks indicators and subfields
type CtrlFld struct {
	Tag   int
	Value []byte // Stripped of field terminator
}

func CtrlFldFrom(tag int, MARCrec []byte) []CtrlFld {
	if tag < 0 || tag >= 10 {
		fmt.Println("Not a control field tag")
		os.Exit(1)
	}

	l := ldr.LdrFrom(MARCrec)
	base := l.BaseAddr()
	d := directory.DirFrom(MARCrec)
	ea := d.EntryFor(tag)
	cfa := make([]CtrlFld, len(ea))
	for i := 0; i < len(ea); i++ {
		value := MARCrec[base+ea[i].Fldofs : base+ea[i].Fldofs+ea[i].Fldlen-1] // Omit field terminator
		cfa[i].Tag = tag
		cfa[i].Value = value
	}
	cfMap[tag] = cfa
	return cfa
}

// Return a field's tag as an int
func (cf CtrlFld) TagOf() int {
	return cf.Tag
}

// Return a field's value as a string
func (cf CtrlFld) ValueOf() string {
	return string(cf.Value)
}

// Map tags to values
var cfMap = make(map[int][]CtrlFld)

// Return the control field with a specified tag
func FldWithTag(tag int) []CtrlFld {
	if tag < 0 || tag >= 10 {
		fmt.Println("Not a control field tag")
		os.Exit(1)
	}
	return cfMap[tag]
}
