package plotter

import (
	"fmt"
	"math"
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
		strconv.FormatFloat(math.Round(x*u.UNIT), 'f', -1, 64),
		strconv.FormatFloat(math.Round(y*u.UNIT), 'f', -1, 64),
	)
	got := ConstructCommand(command, x*u.UNIT, y*u.UNIT)

	if got != want {
		t.Errorf("input: (%f, %f)\ngot: %q,\nwanted: %q\n",
			x, y, got, want,
		)
	}
}
