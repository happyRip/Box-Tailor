package lidded

import (
	"fmt"

	u "github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/plotter"
)

type product struct {
	name string
	size u.Triad
}

func NewProduct(name string, size u.Triad) (product, error) {
	var p product
	p.SetName(name)
	p.SetSize(size.X(), size.Y(), size.Z())
	return p, nil
}

func NewEmptyProduct() product {
	return product{}
}

func (p *product) SetName(name string) error {
	p.name = name
	return nil
}

func (p *product) SetSize(width, depth, height float64) error {
	err := p.size.SetValues(width, depth, height)
	if err != nil {
		return err
	}
	return nil
}

func (p *product) ProcessUserInput() {
	var (
		name    string
		x, y, z float64
		size    u.Triad
	)

	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Print("Dimensions\n    width: ")
	fmt.Scan(&x)
	fmt.Print("    depth: ")
	fmt.Scan(&y)
	fmt.Print("   height: ")
	fmt.Scan(&z)

	size.SetValues(x, y, z)

	p.name = name
	p.size = size
}

func (p product) Name() string {
	return p.name
}

func (p product) Size() (float64, float64, float64, error) {
	return p.size.GetValues()
}

func (p product) Width() float64 {
	return p.size.X()
}

func (p product) Depth() float64 {
	return p.size.Y()
}

func (p product) Height() float64 {
	return p.size.Z()
}

type box struct {
	content        product
	buffer         u.Triad
	variant        string
	boardThickness float64
}

func (b box) Draw() string {
	var out string

	x, y, z, _ := b.InternalSize()
	// thk := b.BoardThickness()
	pen := plotter.NewPen()

	for i := 0; i < 2; i++ {
		yHeight := /* 2*thk + */ x + y
		xFlap := 0.9 * z
		xCover := 0.9 * y
		out += pen.MoveRelative(0, -yHeight)
		out += pen.MoveRelative(xFlap, 0)
		out += pen.MoveRelative(y-xFlap, 0.5*y)
		out += pen.MoveRelative(-0.05*y, -0.5*y)
		out += pen.MoveRelative(xCover, 0)
		out += pen.MoveRelative(0.05*y, 0.5*y)
		out += pen.MoveRelative(-(y - xFlap), -0.5*y)
		out += pen.MoveRelative(xFlap, 0)
		x, y, z = -x, -y, -z
	}

	return out
}

func (b box) Name() string {
	return b.content.Name()
}

func (b box) ContentSize() (float64, float64, float64, error) {
	return b.content.Size()
}

func (b box) MarginSize() (float64, float64, float64, error) {
	return b.buffer.GetValues()
}

func (b box) InternalSize() (float64, float64, float64, error) {
	x, y, z, err := b.ContentSize()
	if err != nil {
		return -1, -1, -1, err
	}

	m, n, o, err := b.MarginSize()
	if err != nil {
		return -1, -1, -1, err
	}

	return x + m, y + n, z + o, nil
}

func (b box) Variant() string {
	return b.variant
}

func (b box) BoardThickness() float64 {
	return b.boardThickness
}
