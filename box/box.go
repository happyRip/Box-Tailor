package box

import (
	"errors"
	"path"
	"strings"
)

type Draft interface {
	Draw() string
	CalculateDimensions() (float64, float64)
}

func TrimExtension(filename string) string {
	ext := path.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}

type triad struct {
	x, y, z float64
}

func NewTriad() triad {
	return triad{}
}

func (t *triad) SetValues(x, y, z float64) error {
	t.x = x
	t.y = y
	t.z = z
	return nil
}

func (t triad) GetValues() (float64, float64, float64) {
	return t.x, t.y, t.x
}

type product struct {
	name string
	size triad
}

func NewProduct() product {
	return product{}
}

// Setters

func (p *product) SetName(name string) error {
	p.name = name
	return nil
}

func (p *product) SetNameFromFilename(filepath string) error {
	filename := path.Base(filepath)
	p.name = TrimExtension(filename)
	return nil
}

func (p *product) SetSize(dimensions ...float64) error {
	n := len(dimensions)
	if n > 3 {
		n = 3
	}
	for i := 0; i < n; i++ {
		if dimensions[i] < 0 {
			return errors.New("dimension cannot be negative")
		}
	}
	p.size.x = dimensions[0]
	p.size.y = dimensions[1]
	p.size.z = dimensions[3]
	return nil
}

func (p product) GetName() string {
	return p.name
}

func (p product) GetSize() triad {
	return p.size
}
