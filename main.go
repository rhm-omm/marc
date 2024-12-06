package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const fname = "c:\\users\\rhmoo\\OneDrive\\Documents\\RMoon.mrc"

func main() {
	data, err := os.ReadFile(fname) // Read in whole file
	if err != nil {
		log.Fatal(err)
	}

	i := bytes.IndexByte(data, 0x1d) // Locate first record terminator (RT)
	r := data[:i]                    // First record, w/o RT
	s := string(r)
	ra := []rune(s)
	fmt.Printf("Length of byte array = %d\n", len(r))
	fmt.Printf("Length of string = %d\n", len(s))
	fmt.Printf("Length of rune array = %d\n", len(ra))
}
