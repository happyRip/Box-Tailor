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
	x, y, z := l.InternalSize()
	thk := l.BoardThickness
	tabHeight := 1.1 * thk
	pen := plotter.Pen{}

	// draw cut lines
	out := []string{
		plotter.SelectPen(1),
		pen.MoveAbsolute(
			tabHeight+2*z+2.5*thk,
			-(tabHeight + 2*z + 2.5*thk),
		),
	}
	for i := 0; i < 2; i++ {
		xTabCount := 2
		yTabCount := 2
		xTab := x / (2*float64(xTabCount) + 1)
		yTab := y / (2*float64(yTabCount) + 1)
		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, 2*z + 2.5*thk},
				}...,
			)...,
		)

		for j := 0; j < xTabCount; j++ {
			out = append(out,
				pen.LineShape(
					[][2]float64{
						{xTab, 0},
						{0, tabHeight},
						{xTab, 0},
						{0, -tabHeight},
					}...,
				)...,
			)
		}
		out = append(out, pen.Line(xTab, 0))

		out = append(out,
			pen.LineShape(
				[][2]float64{
					{0, -(2*z + 2.5*thk)},
					{thk, 0},
					{0, 0.5 * (x - thk)},
					{z, 0},
					{0, -0.5 * (x - thk)},
					{2*thk + z, 0},
				}...,
			)...,
		)

		for j := 0; j < yTabCount; j++ {
			out = append(out,
				pen.LineShape(
					[][2]float64{
						{0, -yTab},
						{tabHeight, 0},
						{0, -yTab},
						{-tabHeight, 0},
					}...,
				)...,
			)
		}
		out = append(out, pen.Line(0, -yTab))

		out = append(out,
			pen.LineShape(
				[][2]float64{
					{-(2*thk + z), 0},
					{0, -0.5 * (x - thk)},
					{-z, 0},
					{0, 0.5 * (x - thk)},
					{-thk, 0},
				}...,
			)...,
		)

		x, y, z = -x, -y, -z
		thk = -thk
		tabHeight *= -1
	}

	// draw fold lines
	out = append(out,
		plotter.SelectPen(2),
		pen.MoveAbsolute(
			tabHeight+z+1.5*thk,
			-(tabHeight+2*z+2*thk),
		),
		// pen.Line(2*z+thk+x, 0),
		pen.Line(z, 0),
		pen.MoveRelative(thk, 0),
		pen.Line(x, 0),
		pen.MoveRelative(thk, 0),
		pen.Line(z, 0),
		pen.MoveRelative(0.5*thk, -0.5*thk),
		pen.Line(0, -y),
		pen.MoveRelative(thk, 0),
		pen.Line(0, y),
		pen.MoveRelative(-(2.5*thk+z), z+thk),
		pen.Line(-x, 0),
		pen.MoveRelative(0, thk),
		pen.Line(x, 0),
		pen.MoveRelative(thk+z, -(2.5*thk+z+y)),
		// pen.Line(-(2*z+thk+x), 0),
		pen.Line(-z, 0),
		pen.MoveRelative(-thk, 0),
		pen.Line(-x, 0),
		pen.MoveRelative(-thk, 0),
		pen.Line(-z, 0),
		pen.MoveRelative(-0.5*thk, 0.5*thk),
		pen.Line(0, y),
		pen.MoveRelative(-thk, 0),
		pen.Line(0, -y),
		pen.MoveRelative(2.5*thk+z, -(thk+z)),
		pen.Line(x, 0),
		pen.MoveRelative(0, -thk),
		pen.Line(-x, 0),
		pen.MoveRelative(-0.5*thk, 2*thk+z),
		pen.Line(0, y),
		pen.MoveRelative(thk+x, 0),
		pen.Line(0, -y),
	)

	return out
}

func (l Lid) CalculateSize() (float64, float64) {
	return 0, 0
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
