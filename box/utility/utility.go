package utility

import (
	"path"
	"strings"
)

func TrimExtension(filename string) string {
	ext := path.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}

func NotPositive(f float64) bool {
	return f <= 0
}

func LessThanZero(f float64) bool {
	return f < 0
}

type Triad struct {
	x, y, z float64
}

func NewTriad() Triad {
	return Triad{}
}

func (t *Triad) SetValues(x, y, z float64) error {
	t.x, t.y, t.z = x, y, z
	return nil
}

func (t Triad) GetValues() (float64, float64, float64, error) {
	return t.x, t.y, t.x, nil
}

type Pair struct {
	x, y float64
}

func (p *Pair) SetValues(x, y float64) error {
	p.x, p.y = x, y
	return nil
}

func (p Pair) GetValues() (float64, float64, error) {
	return p.x, p.y, nil
}
