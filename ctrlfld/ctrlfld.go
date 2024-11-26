package ctrlfld

import (
	"fmt"
	"os"

	"github.com/rhm-omm/marc/fld"
)

// MARC control field. Lacks indicators and subfields
type CtrlFld struct {
	tag   int
	value []byte // Stripped of field terminator
}

// Return a field's tag as an int
func (cf CtrlFld) Tag() int {
	return cf.tag
}

// Return a field's value as a string
func (cf CtrlFld) Value() string {
	return string(cf.value)
}

// Map tags to values
var CfMap = make(map[int]CtrlFld)

// Return the control field with a specified tag
func FldWithTag(tag int) CtrlFld {
	if tag < 0 || tag >= 100 {
		fmt.Println("Not a control field tag")
		os.Exit(1)
	}
	return CfMap[tag]
}

// Return the MARC representation as an array of bytes (with field terminator)
func (cf CtrlFld) MARCrepr() []byte {
	var repr = make([]byte, len(cf.value)+1)
	copy(repr, cf.value)
	repr = append(repr, fld.FLD_TERMINATOR)
	return repr
}
