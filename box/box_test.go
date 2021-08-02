package box

import (
	"math/rand"
	"path"
	"testing"

	u "github.com/happyRip/Box-Tailor/box/utility"
)

func TestProduct(t *testing.T) {
	name := "Some name"
	x, y, z :=
		rand.Float64()*100, rand.Float64()*100, rand.Float64()*100
	boxVariant := "variant"

	p := NewProduct()
	p.SetName(name)
	p.SetSize(x, y, z)
	p.SetVariant(boxVariant)

	if got := p.Name(); got != name {
		t.Errorf("input: (%s)\ngot: %s\nwant: %s\n",
			name, got, name,
		)
	}
	size := u.Triad{}
	size.SetValues(x, y, z)
	if got, _ := p.Size(); got != size {
		t.Errorf("input: (%f, %f, %f)\ngot: %v\nwant: %v\n",
			x, y, z, got, size,
		)
	}
	if got := p.BoxVariant(); got != boxVariant {
		t.Errorf("input: (%s)\ngot: %s\nwant: %s\n",
			boxVariant, got, boxVariant,
		)
	}
	filepath := "some/file/path/to/filename.ext"
	filename := u.TrimExtension(path.Base(filepath))
	p.SetNameFromFilepath(filepath)
	if got := p.Name(); got != filename {
		t.Errorf("input: (%s)\ngot: %s\nwant: %s\n",
			filepath, got, filename,
		)
	}

	// TODO
	// test Draw()
	// test CalculateSize()
}
