package plotter

import (
	"math"
	"strconv"
)

const UNIT = 40 // HPGL uses resolution of 40 units per mm

func FloatToIntTimesTen(f float64) int {
	return int(math.Round(f * 10))
}

func ToStringSingleDecimal(f float64) string {
	reduced := math.Round(f*10) / 10
	return strconv.FormatFloat(reduced, 'f', -1, 64)
}

func IntSingleDecimalToFloat(i int) float64 {
	return float64(i) / 10
}

func ToStringUnits(i int) string {
	var f float64 = IntSingleDecimalToFloat(i)
	return strconv.FormatFloat(f*UNIT, 'f', -1, 64)
}

func getNumbers(s string) []string {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	return re.FindAllString(s, -1)
}

type extremes struct {
	min, max int
}

func (e *extremes) init() {
	e.min, e.max = math.MaxInt64, math.MinInt64
}

func (e *extremes) getExtremes(i int) {
	e.setMin(i)
	e.setMax(i)
}

func (e *extremes) setMin(i int) {
	if e.min > i {
		e.min = i
	}
}

func (e *extremes) setMax(i int) {
	if e.max < i {
		e.max = i
	}
}
