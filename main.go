package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/rhm-omm/marc/directory"
)

const fname = "c:\\users\\rhmoo\\OneDrive\\Documents\\RMoon.mrc"

func main() {
	data, err := os.ReadFile(fname) // Read in whole file
	if err != nil {
		log.Fatal(err)
	}
	i := bytes.IndexByte(data, 0x1d) // Locate first record terminator (RT)
	r := data[:i+1]                  // First record, including RT
	dir := directory.DirFrom(r)      // Build directory
	fmt.Println(len(dir))

	lenarr001 := dir.FldLen(001)
	fmt.Printf("Nr of 001 flds: %d\n", len(lenarr001))
	fmt.Println(lenarr001[0])
	ofsarr001 := dir.FldOfs(001)
	fmt.Println(ofsarr001[0])
	lenarr650 := dir.FldLen(650)
	fmt.Printf("Nr of 650 flds: %d\n", len(lenarr650))
	fmt.Println(lenarr650)
	fmt.Println(dir)
}
