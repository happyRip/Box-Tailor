package lidded

import (
	"math"

	"github.com/happyRip/Box-Tailor/box"
	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/plotter"
)

type Box struct {
	Content        box.Product
	Margin         utility.Triad
	Origin         utility.Pair
	BoardThickness float64
	Kerf           float64
	Debug          bool
}

func (b Box) Draw() []string {
	// draw cut lines
	cuts := []string{plotter.SelectPen(1)}
	cuts = append(cuts,
		b.drawCutLines()...,
	)

	// draw fold lines
	folds := []string{plotter.SelectPen(3)}
	folds = append(folds,
		b.drawFoldLines()...,
	)

	if b.Debug {
		// draw shape without offset
		shape := []string{plotter.SelectPen(5)}
		shape = append(shape,
			b.drawNoOffset()...,
		)
		cuts = append(cuts, shape...)
	}

	return append(cuts, folds...)
}

func (b Box) drawCutLines() []string {
	x, y, z := b.InternalSize()
	t, o := b.BoardThickness, b.Kerf

	var (
		pen plotter.Pen
		out []string
	)
	if b.Origin.X != 0 || b.Origin.Y != 0 {
		out = append(out, pen.MoveAbsolute(b.Origin.X-o, b.Origin.Y-o))
	}
	flap := math.Min(z, x/2)
	out = append(out,
		pen.MoveRelative(0, z+2*t-(flap+t)),
	)
	for i := 0; i < 2; i++ {
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, 2*(flap+t) + y + 2*o},
					{z + 2*o, 0},
					{0, -(flap + t/2)},
					{t - 2*o, 0},
					{0, 1.5*t + z},
					{x + 2*o, 0},
					{0, -(1.5*t + z)},
					{t - 2*o, 0},
					{0, flap + t/2},
					{z + 2*o, 0},
				}...,
			)...,
		)
		x, y, z = -x, -y, -z
		flap = -flap
		t, o = -t, -o
	}

	return out
}

func (b Box) drawFoldLines() []string {
	x, y, z := b.InternalSize()
	t := b.BoardThickness

	var (
		pen plotter.Pen
		out []string
	)
	if b.Origin.X != 0 || b.Origin.Y != 0 {
		out = append(out, pen.MoveAbsolute(b.Origin.X, b.Origin.Y))
	}
	out = append(out,
		pen.MoveRelative(z+t/2, z+t/2),
		pen.DrawRectangle(x+t, y+3*t),
		pen.MoveRelative(-(z+t/2), t),
		pen.Line(z, 0),
		pen.MoveRelative(-z, y+t),
		pen.Line(z, 0),
		pen.MoveRelative(x+2*t, 0),
		pen.Line(z, 0),
		pen.MoveRelative(-z, -(y+t)),
		pen.Line(z, 0),
	)
	return out
}

func (b Box) drawNoOffset() []string {
	x, y, z := b.InternalSize()
	t := b.BoardThickness

	var (
		pen plotter.Pen
		out []string
	)
	if b.Origin.X != 0 || b.Origin.Y != 0 {
		out = append(out, pen.MoveAbsolute(b.Origin.X, b.Origin.Y))
	}
	flap := math.Min(z, x/2)
	out = append(out,
		pen.MoveRelative(0, z+2*t-(flap+t)),
	)
	for i := 0; i < 2; i++ {
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, 2*(flap+t) + y},
					{z, 0},
					{0, -(flap + t/2)},
					{t, 0},
					{0, 1.5*t + z},
					{x, 0},
					{0, -(1.5*t + z)},
					{t, 0},
					{0, flap + t/2},
					{z, 0},
				}...,
			)...,
		)
		x, y, z = -x, -y, -z
		flap = -flap
		t = -t
	}

	return out
}

func (b Box) CalculateSize() (float64, float64) {
	x, y, z := b.InternalSize()
	thk := b.BoardThickness
	return 2*(z+thk+b.Kerf) + thk + x, 2*(z+thk+b.Kerf) + y
}

func (b *Box) SetBuffer(x, y, z float64) {
	b.Margin.SetValues(x, y, z)
}

func (b *Box) SetOrigin(x, y float64) {
	b.Origin.X = x
	b.Origin.Y = y
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
