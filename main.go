package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/rhm-omm/marc/ctrlfld"
)

const fname = "c:\\users\\rhmoo\\OneDrive\\Documents\\RMoon.mrc"

func main() {
	data, err := os.ReadFile(fname) // Read in whole file
	if err != nil {
		log.Fatal(err)
	}
	i := bytes.IndexByte(data, 0x1d) // Locate first record terminator (RT)
	r := data[:i+1]                  // First record, including RT

	cfa := ctrlfld.CtrlFldFrom(r)
	for _, cf := range cfa {
		fmt.Print(cf.TagOf(), ": ")
		lv := len(cf.ValueOf())
		fmt.Println(cf.ValueOf(), " ", lv)
	}
}
