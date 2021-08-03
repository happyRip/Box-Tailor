package utility

import (
	"math"
	"regexp"
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

func GetNumbers(s string) []string {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	return re.FindAllString(s, -1)
}

type extremes struct {
	min, max int
}

func NewExtremes() extremes {
	var e extremes
	e.min, e.max = math.MaxInt64, math.MinInt64
	return e
}

func (e *extremes) GetExtremes(i int) {
	e.SetMin(i)
	e.SetMax(i)
}

func (e *extremes) SetMin(i int) {
	if e.min > i {
		e.min = i
	}
}

func (e *extremes) SetMax(i int) {
	if e.max < i {
		e.max = i
	}
}

func (e extremes) Min() int {
	return e.min
}

func (e extremes) Max() int {
	return e.max
}
