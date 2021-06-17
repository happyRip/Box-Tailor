package plotter

import (
	u "github.com/happyRip/Box-Tailor/src/plotter/utility"
)

type pen struct {
	X, Y int // current position
}

func GetPen() pen {
	return pen{}
}

func (p *pen) MoveAbsolute(x, y float64) string {
	p.X = u.FloatToIntTimesTen(x)
	p.Y = u.FloatToIntTimesTen(y)
	return ConstructCommand("PU", x, y)
}

func (p *pen) MoveRelative(x, y float64) string {
	p.X += u.FloatToIntTimesTen(x)
	p.Y += u.FloatToIntTimesTen(y)
	return ConstructCommand(
		"PU",
		u.IntSingleDecimalToFloat(p.X),
		u.IntSingleDecimalToFloat(p.Y),
	)
}

func (p *pen) Line(x, y float64) string {
	p.X = u.FloatToIntTimesTen(x)
	p.Y = u.FloatToIntTimesTen(y)
	return ConstructCommand("PD", x, y)
}

func SelectPen(i int) string {
	return ConstructCommand("SP", float64(i))
}

func ConstructCommand(command string, args ...float64) string {
	command += ":"
	for i, f := range args {
		command += u.ToStringSingleDecimal(f)
		if i < len(args)-1 {
			command += ","
		}
	}
	command += ";\n"
	return command
}
