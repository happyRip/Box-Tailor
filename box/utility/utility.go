package utility

type Triad struct {
	x, y, z float64
}

func NewTriad() Triad {
	return Triad{}
}

func (t *Triad) SetValues(x, y, z float64) error {
	t.x = x
	t.y = y
	t.z = z
	return nil
}

func (t Triad) GetValues() (float64, float64, float64) {
	return t.x, t.y, t.x
}
