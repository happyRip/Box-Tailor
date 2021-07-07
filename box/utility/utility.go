package utility

type triad struct {
	x, y, z float64
}

func NewTriad() triad {
	return triad{}
}

func (t *triad) SetValues(x, y, z float64) error {
	t.x = x
	t.y = y
	t.z = z
	return nil
}

func (t triad) GetValues() (float64, float64, float64) {
	return t.x, t.y, t.x
}
