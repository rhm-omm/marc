package ctrlfld

import (
	"github.com/rhm-omm/marc/directory"
	"github.com/rhm-omm/marc/ldr"
)

// MARC control field. Lacks indicators and subfields
type CtrlFld struct {
	Tag   int
	Value string
}

func CtrlFldFrom(MARCrec []byte) []CtrlFld {

	l := ldr.LdrFrom(MARCrec)
	baseAdrr := l.BaseAddr()
	d := directory.DirFrom(MARCrec)
	cfa := make([]CtrlFld, 0, len(d))
	for _, e := range d {
		if e.Tag < 10 {
			var cf CtrlFld
			cf.Tag = e.Tag
			v := MARCrec[baseAdrr+e.Fldofs : baseAdrr+e.Fldofs+e.Fldlen-1]
			cf.Value = string(v)
			cfa = append(cfa, cf)
		}
	}
	return cfa
}

func CtrlFldNrFrom(tag int, MARCrec []byte) []CtrlFld {
	allCf := CtrlFldFrom(MARCrec)
	cfa := make([]CtrlFld, 0, len(allCf))
	for _, cf := range allCf {
		if cf.Tag == tag {
			cfa = append(cfa, cf)
		}
	}
	return cfa
}

// Return a field's tag as an int
func (cf CtrlFld) TagOf() int {
	return cf.Tag
}

// Return a field's value as a string w/ FT
func (cf CtrlFld) ValueOf() string {
	return cf.Value
}
