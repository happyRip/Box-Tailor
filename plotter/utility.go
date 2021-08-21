package plotter

import (
	"bufio"
	"errors"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const UNIT = 40 // HPGL uses resolution of 40 units per mm

func floatToIntTimesTen(f float64) int {
	return int(math.Round(f * 10))
}

func toStringSingleDecimal(f float64) string {
	reduced := math.Round(f*10) / 10
	return strconv.FormatFloat(reduced, 'f', -1, 64)
}

func intSingleDecimalToFloat(i int) float64 {
	return float64(i) / 10
}

func toStringUnits(i int) string {
	var f float64 = intSingleDecimalToFloat(i)
	return strconv.FormatFloat(f*UNIT, 'f', -1, 64)
}

func getNumbers(s string) []string {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	return re.FindAllString(s, -1)
}

func GetDimensionsFromFile(source string) (floatPair, error) {
	empty := floatPair{}

	if extension := filepath.Ext(source); extension != ".plt" {
		return empty, errors.New("incorrect file type")
	}

	file, err := os.Open(source)
	if err != nil {
		return empty, err
	}

	x, y := extremes{}, extremes{}
	x.init()
	y.init()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:2] == "PD" {
			stringSlice := getNumbers(scanner.Text())

			for i, v := range stringSlice {
				v, err := strconv.Atoi(v)
				if err != nil {
					return empty, err
				}

				switch i % 2 {
				case 0:
					x.getExtremes(v)
				case 1:
					y.getExtremes(v)
				}

			}
		}
	}

	dimensions := floatPair{
		x: float64(x.max-x.min) / UNIT,
		y: float64(y.max-y.min) / UNIT,
	}

	err = file.Close()
	if err != nil {
		return empty, err
	}
	return dimensions, nil
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

type floatPair struct {
	x, y float64
}
