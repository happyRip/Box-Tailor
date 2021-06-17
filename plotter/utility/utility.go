package utility

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
