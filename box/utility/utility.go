package utility

import (
	"path"
	"strings"
)

func TrimExtension(filename string) string {
	ext := path.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}

func AnyNotPositive(args ...float64) bool {
	for _, v := range args {
		if v <= 0 {
			return true
		}
	}
	return false
}

func AnyLessThanZero(args ...float64) bool {
	for _, v := range args {
		if v < 0 {
			return true
		}
	}
	return false
}

type Triad struct {
	X, Y, Z float64
}

func (t *Triad) SetValues(x, y, z float64) error {
	t.X, t.Y, t.Z = x, y, z
	return nil
}

func (t Triad) GetValues() (float64, float64, float64) {
	return t.X, t.Y, t.Z
}

func (t Triad) X() float64 {
	return t.x
}

func (t Triad) Y() float64 {
	return t.y
}

func (t Triad) Z() float64 {
	return t.z
}

type Pair struct {
	X, Y float64
}

func (p *Pair) SetValues(x, y float64) error {
	p.X, p.Y = x, y
	return nil
}

func (p Pair) GetValues() (float64, float64) {
	return p.X, p.Y
}
