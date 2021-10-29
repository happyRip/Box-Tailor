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
	x, y, z := b.InternalSize()
	t := b.BoardThickness
	o := b.Kerf
	var pen plotter.Pen

	add := 4.
	sep, diff := add+t/2-2*o, 0.
	// if sep < 0 {
	// 	diff = -sep / 2
	// 	sep = 0
	// }

	// draw cut lines
	out := []string{plotter.SelectPen(1)}
	if b.Origin.X != 0 || b.Origin.Y != 0 {
		out = append(out, pen.MoveAbsolute(b.Origin.X, b.Origin.Y))
	}
	for i := 0; i < 2; i++ {
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, -(y + 2*(z+1.5*t+o))},
					{z + 2*o - diff, 0},
					{0, z + t},
					{sep, 0},
					{0, -(z + t)},
					{x + 2*o - 2*diff, 0},
					{0, z + t},
					{sep, 0},
					{0, -(z + t)},
					{z + 2*o - diff, 0},
				}...,
			)...,
		)
		x, y, z = -x, -y, -z
		t, o = -t, -o
		sep, diff = -sep, -diff
	}

	// debug shape without offset
	out = append(out,
		plotter.SelectPen(5),
		pen.MoveAbsolute(b.Origin.X+o, b.Origin.Y-o),
	)
	sep = add + t/2
	for i := 0; i < 2; i++ {
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, -(y + 2*(z+1.5*t))},
					{z, 0},
					{0, z + t},
					{sep, 0},
					{0, -(z + t)},
					{x, 0},
					{0, z + t},
					{sep, 0},
					{0, -(z + t)},
					{z, 0},
				}...,
			)...,
		)
		x, y, z = -x, -y, -z
		t, o = -t, -o
		sep = -sep
	}

	//draw fold lines
	sep = add + t/2
	out = append(out,
		plotter.SelectPen(3),
		pen.MoveRelative(z+sep+o, -(z+0.5*t)-o),
		pen.DrawRectangle(x, -(y+2*t)),
		pen.MoveRelative(-(z+sep), -0.5*t),
		pen.Line(z, 0),
		pen.MoveRelative(-z, -(y+t)),
		pen.Line(z, 0),
		pen.MoveRelative(x+2*sep, 0),
		pen.Line(z, 0),
		pen.MoveRelative(-z, y+t),
		pen.Line(z, 0),
	)

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
