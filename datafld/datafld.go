package datafld

import (
	"fmt"
	"os"

	"github.com/rhm-omm/marc/directory"
	"github.com/rhm-omm/marc/fld"
	"github.com/rhm-omm/marc/ldr"
)

const SUBFLD_DELIMITER byte = 0x1F

type SubFld struct {
	tag     byte
	content []byte
}

type DataFld struct {
	tag       int
	indicator []byte
	subfields []SubFld
}

func DataFldFrom(tag int, MARCrec []byte) {
	if tag < 100 || tag >= 1000 {
		fmt.Println("Not a data field tag")
		os.Exit(1)
	}
	var df DataFld
	df.tag = tag

	start := directory.FldStart(tag)
	length := directory.FldLen(tag)
	l := ldr.LdrFrom(MARCrec)
	base := l.BaseAddr()
	value := MARCrec[base+start : base+start+length-1] // All subfields w/o field terminator
	copy(df.indicator, value[: 2])
	for i := 2; value[i] != fld.FLD_TERMINATOR; i++ {
		var sf SubFld
		if value[i] == SUBFLD_DELIMITER {
			if sf.tag != 0 {
				df.subfields = append(df.subfields, sf)
			}
			i++
			sf.tag = value[i]
			continue
		}
		sf.content = append(sf.content, value[i])
	}

	
	cfMap[cf.tag] = cf

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
