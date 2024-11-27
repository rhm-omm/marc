package ldr

import (
	"fmt"
	"os"
	"strconv"
)

// MARC record leader
type Ldr [24]byte

// Extract the leader from a MARC record
func LdrFrom(MARCrec []byte) Ldr {
	if len(MARCrec) < 24 {
		fmt.Println("Illegal argument: len must be >= 24")
		os.Exit(1)
	}
	valAsSlice := make([]byte, 24)
	copy(valAsSlice, MARCrec[:24])
	return Ldr(valAsSlice)
}

func (l Ldr) BaseAddr() int {
	ASCIIaddr := make([]byte, 5)
	copy(ASCIIaddr, l[12:17])
	val, err := strconv.Atoi(string(ASCIIaddr))
	if err != nil {
		fmt.Printf("Not a number")
		os.Exit(1)
	}
	return val
}

func (l Ldr) RecLen() int {
	ASCIIaddr := make([]byte, 5)
	copy(ASCIIaddr, l[:5])
	val, err := strconv.Atoi(string(ASCIIaddr))
	if err != nil {
		fmt.Printf("Not a number")
		os.Exit(1)
	}
	return val
}
