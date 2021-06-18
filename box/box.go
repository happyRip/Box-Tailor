package box

import "errors"

type triad struct {
	x, y, z float64
}

type product struct {
	name string
	size triad
}

func NewProduct() product {
	return product{}
}

func (p *product) SetName(name string) error {
	p.name = name
	return nil
}

func (p *product) SetSize(dimensions triad) error {
	d := []float64{dimensions.x, dimensions.y, dimensions.z}
	for v := range d {
		if v < 0 {
			return errors.New("dimension cannot be negative")
		}
	}
	p.size = dimensions
	return nil
}

func (p product) GetName() string {
	return p.name
}

func (p product) GetSize() triad {
	return p.size
}
