package datafld

import (
	"fmt"
	"os"
)

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

func FieldWithTag(tag int) DataFld {
	if tag < 100 {
		fmt.Printf("Not a data field tag")
		os.Exit(1)
	}
	return dfMap[tag]
}
