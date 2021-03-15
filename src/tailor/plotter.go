package tailor

import (
	"strconv"
)

const unit = 40 // HPGL uses resolution of 40 units per mm

type Pen struct {
	X, Y float32
}

func (p *Pen) Move(x, y float32) string {
	p.X, p.Y = p.X+x, p.Y+y
	return "PU:" + strconv.FormatFloat(float64(p.X*unit), 'f', -1, 32) + "," + strconv.FormatFloat(float64(p.Y*unit), 'f', -1, 32) + ";\n"
}

func (p *Pen) Line(x, y float32) string {
	p.X, p.Y = p.X+x, p.Y+y
	return "PD:" + strconv.FormatFloat(float64(p.X*unit), 'f', -1, 32) + "," + strconv.FormatFloat(float64(p.Y*unit), 'f', -1, 32) + ";\n"
}
