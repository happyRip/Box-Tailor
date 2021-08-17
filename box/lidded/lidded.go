package lidded

import (
	"github.com/happyRip/Box-Tailor/box"
	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/plotter"
)

type Box struct {
	Content        box.Product
	Margin         utility.Triad
	Origin         utility.Pair
	BoardThickness float64
}

func (b Box) Draw() []string {
	x, y, z := b.InternalSize()
	thk := b.BoardThickness
	var pen plotter.Pen

	// draw cut lines
	out := []string{plotter.SelectPen(1)}
	if b.Origin.X != 0 || b.Origin.Y != 0 {
		out = append(out, pen.MoveAbsolute(b.Origin.X, b.Origin.Y))
	}
	for i := 0; i < 2; i++ {
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, -(2*(thk+z) + y)},
					{z, 0},
					{0, z + thk},
					{0.5 * thk, 0},
					{0, -(z + thk)},
					{x + thk, 0},
					{0, z + thk},
					{0.5 * thk, 0},
					{0, -(z + thk)},
					{z, 0},
				}...,
			)...,
		)
		x, y, z = -x, -y, -z
		thk = -thk
	}

	//draw fold lines
	out = append(out,
		plotter.SelectPen(2),
		pen.MoveRelative(z+0.5*thk, -(z+0.5*thk)),
		pen.DrawRectangle(x+thk, -(y+thk)),
		pen.MoveAbsolute(b.Origin.X, -(z+thk)),
		pen.Line(z, 0),
		pen.MoveRelative(-z, -y),
		pen.Line(z, 0),
		pen.MoveRelative(2*thk+x, 0),
		pen.Line(z, 0),
		pen.MoveRelative(-z, y),
		pen.Line(z, 0),
	)

	return out
}

func (b Box) CalculateSize() (float64, float64) {
	x, y, z := b.InternalSize()
	thk := b.BoardThickness
	return 2*(z+thk) + x, 2*(z+thk) + y
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
