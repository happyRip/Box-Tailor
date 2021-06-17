package plotter

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	u "github.com/happyRip/Box-Tailor/plotter/utility"
)

func TestConstructCommand(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	command := "AB"
	x, y :=
		rand.Float64()*float64(rand.Intn(150)),
		rand.Float64()*float64(rand.Intn(150))

	want := fmt.Sprintf("%s:%s,%s;\n",
		command,
		strconv.FormatFloat(x*u.UNIT, 'f', -1, 64),
		strconv.FormatFloat(y*u.UNIT, 'f', -1, 64),
	)
	got := ConstructCommand(command, x, y)

	if got != want {
		t.Errorf("got: %q, wanted: %q\n", got, want)
	}
}
