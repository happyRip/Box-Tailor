package lidded

import (
	"github.com/happyRip/Box-Tailor/box"
	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/plotter"
)

type Lid struct {
	Content        box.Product
	Margin         utility.Triad
	BoardThickness float64
}

func (l Lid) Draw() []string {
	// x, y, z := l.InternalSize()
	pen := plotter.Pen{}

	// draw cut lines
	out := []string{plotter.SelectPen(1)}
	for i := 0; i < 4; i++ {
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, 0},
				}...,
			)...,
		)
	}

	return out
}

func (l Lid) ContentSize() (float64, float64, float64) {
	return l.Content.Size.GetValues()
}

func (l Lid) MarginSize() (float64, float64, float64) {
	return l.Margin.GetValues()
}

func (l Lid) InternalSize() (float64, float64, float64) {
	x, y, z := l.ContentSize()
	m, n, o := l.MarginSize()
	return x + m, y + n, z + o
}
