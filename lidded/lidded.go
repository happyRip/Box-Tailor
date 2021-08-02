package lidded

import (
	"github.com/happyRip/Box-Tailor/box"
	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/plotter"
)

type Drafter interface {
	Draw() string
}

type Box struct {
	Content        box.Product
	Margin         utility.Triad
	BoardThickness float64
}

func (b Box) Draw() []string {
	var out []string
	x, y, z := b.InternalSize()
	thk := b.BoardThickness
	pen := plotter.Pen{}

	// draw outer box shape
	out = append(out, plotter.SelectPen(1))
	for i := 0; i < 2; i++ {
		yHeight := 2*thk + x + y
		xFlap := 0.9 * z
		xCover := 0.9 * y
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, -yHeight},
					{xFlap, 0},
					{z - xFlap, 0.5 * (y + thk)},
					{thk, 0},
					{0.05 * y, -0.5 * (y + thk)},
					{xCover, 0},
					{0.05 * y, 0.5 * (y + thk)},
					{thk, 0},
					{z - xFlap, -0.5 * (y + thk)},
					{xFlap, 0},
				}...,
			)...,
		)
		x, y, z = -x, -y, -z
		thk = -thk
	}

	// draw fold lines
	out = append(out,
		plotter.SelectPen(2),
		pen.MoveAbsolute(0, -(0.5*(y+thk))),
		pen.Line(2*z+y+2*thk, 0),
		pen.MoveAbsolute(0, -(1.5*thk+x+0.5*y)),
		pen.Line(2*z+y+2*thk, 0),
		pen.MoveAbsolute(z+0.5*thk, -(0.5*(y+thk))),
		pen.Line(0, -(x+thk)),
		pen.MoveAbsolute(z+y+1.5*thk, -(0.5*(y+thk))),
		pen.Line(0, -(x+thk)),
	)
	return out
}

func (b *Box) SetBuffer(x, y, z float64) {
	b.Margin.SetValues(x, y, z)
}

func (b Box) ContentSize() (float64, float64, float64) {
	return b.Content.Size.GetValues()
}

func (b Box) MarginSize() (float64, float64, float64) {
	return b.Margin.GetValues()
}

func (b Box) InternalSize() (float64, float64, float64) {
	x, y, z := b.ContentSize()
	m, n, o := b.MarginSize()
	return x + m, y + n, z + o
}
