package box

import (
	"errors"
	"path"

	u "github.com/happyRip/Box-Tailor/box/utility"
)

type Draft interface {
	Draw() string
	CalculateSize() (float64, float64)
}

type product struct {
	name       string
	size       u.Triad
	boxVariant string
}

func NewProduct() product {
	return product{}
}

// TODO
func (p product) Draw() string {
	return ""
}

// TODO
func (p *product) CalculateSize() (float64, float64) {
	return 0, 0
}

func (p *product) SetName(name string) error {
	p.name = name
	return nil
}

func (p *product) SetNameFromFilename(filepath string) error {
	filename := path.Base(filepath)
	p.name = u.TrimExtension(filename)
	return nil
}

func (p *product) SetSize(x, y, z float64) error {
	if u.AnyNotPositive(x, y, z) {
		return errors.New("dimensions cannot be negative")
	}
	p.size.SetValues(x, y, z)
	return nil
}

func (p *product) SetVariant(variant string) error {
	p.boxVariant = variant
	return nil
}

func (p product) GetName() (string, error) {
	return p.name, nil
}

func (p product) GetSize() (u.Triad, error) {
	return p.size, nil
}

func (p product) GetVariant() (string, error) {
	return p.boxVariant, nil
}

type board struct {
	size   u.Pair
	margin u.Pair
}

func (b *board) SetSize(x, y float64) error {
	if u.AnyNotPositive(x, y) {
		return errors.New("dimension not positive")
	}
	b.size.SetValues(x, y)
	return nil
}

func (b *board) SetMargin(x, y float64) error {
	if u.AnyLessThanZero(x, y) {
		return errors.New("dimension less than zero")
	}
	b.margin.SetValues(x, y)
	return nil
}

func (b board) GetSize() (float64, float64, error) {
	return b.size.GetValues()
}

func (b board) GetMargin() (float64, float64, error) {
	return b.margin.GetValues()
}

type sortingAlgorithm interface {
	Sort() error
	GetSorted() ([][]product, error)
}

type rackParams struct {
	width, height float64
	margin        u.Pair
}

type shelf struct {
	toSort         []product
	sortedRacks    [][]product
	rackParameters rackParams
}

// TODO
func (s *shelf) ShelfPack() ([]product, error) {
	return nil, nil
}
