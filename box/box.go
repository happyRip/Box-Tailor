package box

import (
	"errors"
	"path"
	"sort"

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
func (p product) CalculateSize() (float64, float64) {
	return 0, 0
}

func (p *product) SetName(name string) error {
	p.name = name
	return nil
}

func (p *product) SetNameFromFilepath(filepath string) error {
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

func (p product) Name() string {
	return p.name
}

func (p product) Size() (u.Triad, error) {
	return p.size, nil
}

func (p product) SizeX() float64 {
	x, _, _, _ := p.size.GetValues()
	return x
}

func (p product) SizeY() float64 {
	_, y, _, _ := p.size.GetValues()
	return y
}

func (p product) SizeZ() float64 {
	_, _, z, _ := p.size.GetValues()
	return z
}

func (p product) BoxVariant() string {
	return p.boxVariant
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

func (b board) Size() (float64, float64, error) {
	return b.size.GetValues()
}

func (b board) SizeX() float64 {
	x, _, _ := b.size.GetValues()
	return x
}

func (b board) SizeY() float64 {
	_, y, _ := b.size.GetValues()
	return y
}

func (b board) Margin() (float64, float64, error) {
	return b.margin.GetValues()
}

type shelfParams struct {
	width, height float64
	margin        u.Pair
}

func (s *shelfParams) SetDimensions(width, height float64) error {
	s.width = width
	s.height = height
	return nil
}

func (s *shelfParams) SetMargin(x, y float64) error {
	err := s.margin.SetValues(x, y)
	if err != nil {
		return err
	}
	return nil
}

func (s shelfParams) Width() float64 {
	return s.width
}

func (s shelfParams) Height() float64 {
	return s.height
}

func (s shelfParams) GetDimensions() (float64, float64, error) {
	return s.width, s.height, nil
}

func (s shelfParams) Margin() (u.Pair, error) {
	return s.margin, nil
}

type rack struct {
	productList     []product
	shelfList       [][]product
	shelfParameters shelfParams
}

func (r *rack) ShelfPack() error {
	if w, h, _ := r.shelfParameters.GetDimensions(); u.AnyLessThanZero(w, h) {
		return errors.New("rackParameters dimensions not positive")
	}
	sort.SliceStable(r.productList,
		func(i, j int) bool {
			iV := r.productList[i].SizeY()
			jV := r.productList[j].SizeY()
			return iV < jV
		},
	)

	products := r.productList
	var (
		shelf   []product
		currPos float64
	)
	for len(products) > 0 {
		i := LessOrEqual('x', r.Width()-currPos, products...)

		if i == -1 {
			r.AppendShelf(shelf...)
			shelf = []product{}
			currPos = 0
			i = 0
		}

		shelf = append(shelf, products[i])
		products = RemoveFromProductSlice(i, products...)
		currPos += shelf[len(shelf)-1].SizeX()

		if len(products) == 0 {
			r.AppendShelf(shelf...)
		}
	}
	return nil
}

func (r *rack) AppendShelf(products ...product) {
	r.shelfList = append(r.shelfList, products)
}

func (r rack) Width() float64 {
	return r.shelfParameters.Width()
}

func (r rack) Height() float64 {
	return r.shelfParameters.Height()
}

func RemoveFromProductSlice(i int, products ...product) []product {
	if i >= len(products)-1 {
		return products[:i]
	}
	return append(products[:i], products[i+1:]...)
}

// TODO
func LessOrEqual(axis rune, value float64, args ...product) int {
	return 0
}
