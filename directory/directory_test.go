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

	len001 := dir.FldLen(001)
	if len001 != 17 {
		t.Errorf("Expected %d, got %d", 17, len001)
	}

	len830 := dir.FldLen(830)
	if len830 != 36 {
		t.Errorf("Expected %d, got %d", 36, len830)
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

	ofs001 := dir.FldOfs(001)
	if ofs001 != 0 {
		t.Errorf("Expected %d, got %d", 0, ofs001)
	}

	ofs830 := dir.FldOfs(830)
	if ofs830 != 3482 {
		t.Errorf("Expected %d, got %d", 3482, ofs830)
	}

}
