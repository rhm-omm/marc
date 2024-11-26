package ctrlfld

// MARC control field. Lacks indicators and subfields
type CtrlFld struct {
	tag   int
	value []byte // Stripped of field terminator
}

func (cf CtrlFld) Tag() int {
	return cf.tag
}

func (cf CtrlFld) Value() string {
	return string(cf.value)
}

var CfMap = make(map[int]string)

func ValueOf(tag int) string {
	return CfMap[tag]
}
