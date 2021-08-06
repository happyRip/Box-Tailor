package box

import (
	"errors"
	"fmt"
	"path"
	"sort"

	"github.com/happyRip/Box-Tailor/box/utility"
	u "github.com/happyRip/Box-Tailor/box/utility"
)

type Drafter interface {
	Draw() []string
	CalculateSize() (float64, float64)
}

type Product struct {
	Name       string
	Size       u.Triad
	BoxVariant string
}

// TODO
func (p Product) Draw() string {
	return ""
}

// TODO
func (p Product) CalculateSize() (float64, float64) {
	return 0, 0
}

// lidded main
func (p *Product) ProcessUserInput() {
	var name string
	fmt.Print("Podaj nazwę pliku wyjściowego: ")
	fmt.Scanln(&name)
	fmt.Print("Podaj wymiary zawartości pudełka:\n      długość (x) [mm]: ")

	var x, y, z float64
	fmt.Scan(&x)
	fmt.Print("    szerokość (y) [mm]: ")
	fmt.Scan(&y)
	fmt.Print("     wysokość (z) [mm]: ")
	fmt.Scan(&z)

	var size utility.Triad
	size.SetValues(x, y, z)

	p.Name = name
	p.Size = size
}

func (p *Product) SetNameFromFilepath(filepath string) error {
	filename := path.Base(filepath)
	p.Name = u.TrimExtension(filename)
	return nil
}

func (p *Product) SetSize(x, y, z float64) error {
	if u.AnyNotPositive(x, y, z) {
		return errors.New("dimensions cannot be negative")
	}
	p.Size.SetValues(x, y, z)
	return nil
}

func (p Product) SizeX() float64 {
	x := p.Size.X
	return x
}

func (p Product) SizeY() float64 {
	y := p.Size.Y
	return y
}

func (p Product) SizeZ() float64 {
	z := p.Size.Z
	return z
}

type Board struct {
	Size   u.Pair
	Margin u.Pair
}

func (b *Board) SetSize(x, y float64) error {
	if u.AnyNotPositive(x, y) {
		return errors.New("dimension not positive")
	}
	b.Size.SetValues(x, y)
	return nil
}

func (b *Board) SetMargin(x, y float64) error {
	if u.AnyLessThanZero(x, y) {
		return errors.New("dimension less than zero")
	}
	b.Margin.SetValues(x, y)
	return nil
}

func (b Board) SizeX() float64 {
	x := b.Size.X
	return x
}

func (b Board) SizeY() float64 {
	y := b.Size.Y
	return y
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
	productList     []Product
	shelfList       [][]Product
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
		shelf   []Product
		currPos float64
	)
	for len(products) > 0 {
		i := LessOrEqual('x', r.Width()-currPos, products...)

		if i == -1 {
			r.AppendShelf(shelf...)
			shelf = []Product{}
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

func (r *rack) AppendShelf(products ...Product) {
	r.shelfList = append(r.shelfList, products)
}

func (r rack) Width() float64 {
	return r.shelfParameters.Width()
}

func (r rack) Height() float64 {
	return r.shelfParameters.Height()
}

func RemoveFromProductSlice(i int, products ...Product) []Product {
	if i >= len(products)-1 {
		return products[:i]
	}
	return append(products[:i], products[i+1:]...)
}

// TODO
func LessOrEqual(axis rune, value float64, args ...Product) int {
	return 0
}
