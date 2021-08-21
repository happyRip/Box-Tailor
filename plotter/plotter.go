package plotter

import (
	"math"
	"strconv"
)

type Pen struct {
	x, y int // current position
}

func NewPen(x, y float64) Pen {
	p := Pen{}
	p.SetPosition(x, y)
	return p
}

func (p *Pen) MoveAbsolute(x, y float64) string {
	p.x = floatToIntTimesTen(x)
	p.y = floatToIntTimesTen(y)
	return ConstructCommand("PU", x, y)
}

func (p *Pen) MoveRelative(x, y float64) string {
	p.x += floatToIntTimesTen(x)
	p.y += floatToIntTimesTen(y)
	return ConstructCommand(
		"PU",
		intSingleDecimalToFloat(p.x),
		intSingleDecimalToFloat(p.y),
	)
}

func (p *Pen) Line(x, y float64) string {
	p.x = floatToIntTimesTen(x)
	p.y = floatToIntTimesTen(y)
	return ConstructCommand("PD", x, y)
}

func (p Pen) Rectangle(width, height float64) string {
	var rect string
	for i := 0; i < 2; i++ {
		rect += p.Line(width, 0)
		rect += p.Line(0, height)
		width *= -1
		height *= -1
	}
	return rect
}

func (p *Pen) SetPosition(x, y float64) {
	p.x = floatToIntTimesTen(x)
	p.y = floatToIntTimesTen(y)
}

func (p *Pen) SetX(x float64) {
	p.x = floatToIntTimesTen(x)
}

func (p *Pen) SetY(y float64) {
	p.y = floatToIntTimesTen(y)
}

func (p Pen) X() float64 {
	return intSingleDecimalToFloat(p.x)
}

func (p Pen) Y() float64 {
	return intSingleDecimalToFloat(p.y)
}

func SelectPen(i int) string {
	return ConstructCommand("SP", float64(i))
}

func ConstructCommand(command string, args ...float64) string {
	command += ":"
	for i, f := range args {
		r := math.Round(f)
		command += strconv.FormatFloat(r, 'f', -1, 64)
		if i < len(args)-1 {
			command += ","
		}
	}
	command += ";\n"
	return command
}
