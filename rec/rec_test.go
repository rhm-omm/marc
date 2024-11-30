package rec

import (
	"log"
	"os"
	"testing"

	"github.com/rhm-omm/marc/rec"
)

var fname = "c:\\users\\rhmoo\\OneDrive\\Documents\\RMoon.mrc"

func TestNext(t *testing.T) {
	var f, err = os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	r := rec.Next(f, 0)
}
