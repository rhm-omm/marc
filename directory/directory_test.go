package directory

import (
	"bytes"
	"log"
	"os"
	"testing"
)

const fname = "c:\\users\\rhmoo\\OneDrive\\Documents\\RMoon.mrc"

func TestFldLen(t *testing.T) {
	data, err := os.ReadFile(fname) // Read in whole file
	if err != nil {
		log.Fatal(err)
	}
	i := bytes.IndexByte(data, 0x1d) // Locate first record terminator (RT)
	r := data[:i+1]                  // First record, including RT
	dir := DirFrom(r)                // Build directory

	lenarr001 := dir.FldLen(001)
	if len(lenarr001) != 1 {
		t.Errorf("Nr of 001 flds: expected %d, got %d", 1, len(lenarr001))
	}
	if len(lenarr001) == 1 && lenarr001[0] != 17 {
		t.Errorf("Expected %d, got %d", 17, lenarr001[0])
	}

	lenarr830 := dir.FldLen(830)
	if len(lenarr830) != 1 {
		t.Errorf("Nr of 830 flds: expected %d, got %d", 1, len(lenarr830))
	}
	if len(lenarr830) == 1 && lenarr830[0] != 36 {
		t.Errorf("Expected %d, got %d", 36, lenarr830[0])
	}
}

func TestFldOfs(t *testing.T) {
	data, err := os.ReadFile(fname) // Read in whole file
	if err != nil {
		log.Fatal(err)
	}
	i := bytes.IndexByte(data, 0x1d) // Locate first record terminator (RT)
	r := data[:i+1]                  // First record, including RT
	dir := DirFrom(r)                // Build directory

	ofsarr001 := dir.FldOfs(001)
	if len(ofsarr001) != 1 {
		t.Errorf("Nr of 001 flds: expected %d, got %d", 1, len(ofsarr001))
	}
	if len(ofsarr001) == 1 && ofsarr001[0] != 0 {
		t.Errorf("Expected %d, got %d", 0, ofsarr001[0])
	}

	ofsarr830 := dir.FldOfs(830)
	if len(ofsarr830) != 1 {
		t.Errorf("Nr of 830 flds: expected %d, got %d", 1, len(ofsarr830))
	}
	if len(ofsarr830) == 1 && ofsarr830[0] != 3482 {
		t.Errorf("Expected %d, got %d", 3482, ofsarr830[0])
	}
}
