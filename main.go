package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/rhm-omm/marc/directory"
	"github.com/rhm-omm/marc/ldr"
)

const fname = "c:\\users\\rhmoo\\OneDrive\\Documents\\RMoon.mrc"

func main() {

	data, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	i := bytes.IndexByte(data, 0x1d)
	fmt.Printf("%d\n", i)
	var r = data[:i+1] // Include RT
	fmt.Printf("%x\n", r[len(r)-1])
	l := ldr.LdrFrom(r)
	fmt.Println(l.BaseAddr())

	dir := directory.DirFrom(r)
	fmt.Println(dir.FldStart(245))
	fmt.Println(dir.FldLen(245))

}
