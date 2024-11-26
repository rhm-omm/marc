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

func (l Ldr) BaseAddress() int {
	addrAscii := make([]byte, 5)
	copy(addrAscii, l[12:17])
	val, err := strconv.Atoi(string(addrAscii))
	if err != nil {
		fmt.Printf("Not a number")
		os.Exit(1)
	}
	return val
}

func (l Ldr) RecLen() int {
	addrAscii := make([]byte, 5)
	copy(addrAscii, l[:5])
	val, err := strconv.Atoi(string(addrAscii))
	if err != nil {
		fmt.Printf("Not a number")
		os.Exit(1)
	}
	return val
}
