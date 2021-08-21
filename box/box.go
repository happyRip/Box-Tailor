package box

import (
	"errors"
	"path"
)

type Draft interface {
	Draw() string
	CalculateSize() (float64, float64)
}

type product struct {
	name string
	size Triad
}

func NewEmptyProduct() product {
	return product{}
}

func NewProduct(name string, size Triad) (product, error) {
	p := product{}
	p.SetName(name)
	err := p.SetSize(size.GetValues())
	return p, err
}

func (p *product) SetName(name string) error {
	p.name = name
	return nil
}

func (p *product) SetNameFromFilepath(filepath string) error {
	filename := path.Base(filepath)
	p.name = trimExtension(filename)
	return nil
}

func (p *product) SetSize(x, y, z float64) error {
	if anyNotPositive(x, y, z) {
		return errors.New("dimensions cannot be negative")
	}
	p.size.SetValues(x, y, z)
	return nil
}

func (p product) Name() string {
	return p.name
}

func (p product) Size() (Triad, error) {
	return p.size, nil
}
