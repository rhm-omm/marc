package datafld

import (
	"fmt"
	"os"
)

const SUBFLD_DELIMITER byte = 0x1F

type SubFld struct {
	tag     byte
	content []byte
}

type DataFld struct {
	tag       int
	indicator [2]byte
	subfields []SubFld
}

func (df DataFld) Tag() int {
	return df.tag
}

var dfMap = make(map[int]DataFld)

func FldWithTag(tag int) DataFld {
	if tag < 100 {
		fmt.Printf("Not a data field tag")
		os.Exit(1)
	}
	return dfMap[tag]
}
