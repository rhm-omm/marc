package ctrlfld

import (
	"bytes"
	"log"
	"os"
	"testing"
)

const fname = "c:\\users\\rhmoo\\OneDrive\\Documents\\RMoon.mrc"

func TestCtrlFldFrom(t *testing.T) {
	data, err := os.ReadFile(fname) // Read in whole file
	if err != nil {
		log.Fatal(err)
	}
	i := bytes.IndexByte(data, 0x1d) // Locate first record terminator (RT)
	r := data[:i+1]                  // First record, including RT
	cfa := CtrlFldFrom(r)
	if len(cfa) != 3 {
		t.Errorf("Nr of ctrl flds: expected 3, got %d", len(cfa))
	}

}

func TestCtrlFldNrFrom(t *testing.T) {
	data, err := os.ReadFile(fname) // Read in whole file
	if err != nil {
		log.Fatal(err)
	}
	i := bytes.IndexByte(data, 0x1d) // Locate first record terminator (RT)
	r := data[:i+1]                  // First record, including RT
	cfa := CtrlFldNrFrom(8, r)
	if len(cfa) != 1 {
		t.Errorf("Nr of 008 flds: expected 1, got %d", len(cfa))
	}
	v008 := cfa[0].ValueOf()
	if len(v008) != 38 {
		t.Errorf("Len of 008 fld w/ FT: expected 38, got %d", len(v008))
	}
}
