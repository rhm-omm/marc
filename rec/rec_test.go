package rec

import (
	"log"
	"os"
	"testing"

	"github.com/rhm-omm/marc/ldr"
)

const fname = "c:\\users\\rhmoo\\OneDrive\\Documents\\RMoon.mrc"

func TestNext(t *testing.T) {
	var f, err = os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	r := Next(f, 0)
	l := ldr.LdrFrom(r)
	recLen := l.RecLen()
	if recLen != 4108 {
		t.Errorf("4108 expected, got %d", recLen)
	}

}
