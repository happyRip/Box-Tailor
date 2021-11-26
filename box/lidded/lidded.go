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
	Kerf           float64
}

func (b Box) Draw() []string {
	// draw cut lines
	cuts := []string{plotter.SelectPen(1)}
	cuts = append(cuts,
		b.drawCutLines()...,
	)

	//draw fold lines
	folds := []string{plotter.SelectPen(3)}
	folds = append(folds,
		b.drawFoldLines()...,
	)

	shape := []string{plotter.SelectPen(5)}
	shape = append(shape,
		b.drawNoOffset()...,
	)
	cuts = append(cuts, shape...)

	return append(cuts, folds...)
}

func (b Box) drawCutLines() []string {
	x, y, z := b.InternalSize()
	t := b.BoardThickness
	o := b.Kerf

	add := 4.
	sep, diff := add+t/2-2*o, 0.

	var (
		pen plotter.Pen
		out []string
	)
	if b.Origin.X != 0 || b.Origin.Y != 0 {
		out = append(out, pen.MoveAbsolute(b.Origin.X, b.Origin.Y))
	}
	for i := 0; i < 2; i++ {
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, -(y + 2*(z+2*t+o))},
					{z + 2*o - diff - t/2, 0},
					{0, z + 1.5*t},
					{sep, 0},
					{0, -(z + 1.5*t)},
					{x + 2*o - 2*diff, 0},
					{0, z + 1.5*t},
					{sep, 0},
					{0, -(z + 1.5*t)},
					{z + 2*o - diff - t/2, 0},
				}...,
			)...,
		)
		x, y, z = -x, -y, -z
		t, o = -t, -o
		sep, diff = -sep, -diff
	}

	return out
}

func (b Box) drawFoldLines() []string {
	x, y, z := b.InternalSize()
	t := b.BoardThickness
	o := b.Kerf

	add := 4.
	sep := add + t/2

	var (
		pen plotter.Pen
		out []string
	)
	if b.Origin.X != 0 || b.Origin.Y != 0 {
		out = append(out, pen.MoveAbsolute(b.Origin.X, b.Origin.Y))
	}
	out = append(out,
		pen.MoveRelative(z+sep+o-t, -(z+t/2)-o),
		pen.DrawRectangle(x+t, -(y+3*t)),
		pen.MoveRelative(-z-0.25*sep, -t),
		pen.Line(z, 0),
		pen.MoveRelative(-z, -(y+t)),
		pen.Line(z, 0),
		pen.MoveRelative(x+2*sep+t, 0),
		pen.Line(z, 0),
		pen.MoveRelative(-z, y+t),
		pen.Line(z, 0),
	)
	return out
}

func (b Box) drawNoOffset() []string {
	x, y, z := b.InternalSize()
	t := b.BoardThickness
	o := b.Kerf

	add := 4.
	sep, diff := add + t/2, 0.

	var (
		pen plotter.Pen
		out []string
	)
	if b.Origin.X != 0 || b.Origin.Y != 0 {
		out = append(out, pen.MoveAbsolute(b.Origin.X, b.Origin.Y))
	}
	out = append(out,
		pen.MoveRelative(o, -o),
	)
	o = 0.
	for i := 0; i < 2; i++ {
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, -(y + 2*(z + 2*t))},
					{z + 2*o - diff - t/2, 0},
					{0, z + 1.5*t},
					{sep, 0},
					{0, -(z + 1.5*t)},
					{x + 2*o - 2*diff, 0},
					{0, z + 1.5*t},
					{sep, 0},
					{0, -(z + 1.5*t)},
					{z + 2*o - diff - t/2, 0},
				}...,
			)...,
		)
		x, y, z = -x, -y, -z
		t, o = -t, -o
		sep, diff = -sep, -diff
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
