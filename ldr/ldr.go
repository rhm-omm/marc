package ldr

import (
	"fmt"
	"os"
	"strconv"
)

type Ldr [24]byte

func From(ba []byte) Ldr {
	if len(ba) < 24 {
		fmt.Println("Illegal argument: len must be >= 24")
		os.Exit(1)
	}
	valAsSlice := make([]byte, 24)
	copy(valAsSlice, ba[:24])
	return Ldr(valAsSlice)
}

func (l Ldr) BaseAddr() int {
	addrASCII := make([]byte, 5)
	copy(addrASCII, l[12:17])
	val, err := strconv.Atoi(string(addrASCII))
	if err != nil {
		fmt.Printf("Not a number")
		os.Exit(1)
	}
	return val
}

func (l Ldr) RecLen() int {
	addrASCII := make([]byte, 5)
	copy(addrASCII, l[:5])
	val, err := strconv.Atoi(string(addrASCII))
	if err != nil {
		fmt.Printf("Not a number")
		os.Exit(1)
	}
	return val
}
