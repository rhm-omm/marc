package rec

import (
	"log"
	"os"
	"strconv"
)

type MARCrec []byte

func Next(f *os.File, ofs int64) MARCrec {

	var recLenASCII []byte = make([]byte, 5)

	n1, err := f.ReadAt(recLenASCII, 0)
	if err != nil {
		log.Fatal(err)
	}
	if n1 != 5 {
		log.Fatalf("5 bytes expected, but %d bytes read", n1)
	}

	recLen, err := strconv.Atoi(string(recLenASCII))
	if err != nil {
		log.Fatal(err)
	}

	var rec []byte = make([]byte, recLen)

	n2, err := f.ReadAt(rec, 0)
	if err != nil {
		log.Fatal(err)
	}
	if n2 != recLen {
		log.Fatalf("%d bytes expected, but %d bytes read", recLen, n2)
	}

	return rec
}
