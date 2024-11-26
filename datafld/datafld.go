package datafld

const SUBFLD_DELIMITER byte = 0x1F

type SubFld struct {
	tag     byte
	content string
}

type DataFld struct {
	tag       int
	indicator [2]byte
	subfields []SubFld
}

var dfMap = make(map[int]DataFld)
