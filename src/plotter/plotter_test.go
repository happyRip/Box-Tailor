package plotter

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	u "./utility"
)

func TestConstructCommand(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	command := "AB"
	x, y :=
		rand.Float64()*float64(rand.Intn(150)),
		rand.Float64()*float64(rand.Intn(150))

	want := fmt.Sprintf("%s:%s,%s;\n", command, x*u.UNIT, y*u.UNIT)
	got := ConstructCommand(command, x, y)

	if got != want {
		t.Errorf("got: %q, wanted: %q\n", got, want)
	}
}
