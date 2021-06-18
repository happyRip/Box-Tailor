package plotter

import (
	"bufio"
	"errors"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	u "github.com/happyRip/Box-Tailor/plotter/utility"
)

type pen struct {
	X, Y int // current position
}

func NewPen() pen {
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
		r := math.Round(f)
		command += strconv.FormatFloat(r, 'f', -1, 64)
		if i < len(args)-1 {
			command += ","
		}
	}
	command += ";\n"
	return command
}

type floatPair struct {
	x, y float64
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
		x: float64(x.max-x.min) / u.UNIT,
		y: float64(y.max-y.min) / u.UNIT,
	}

	err = file.Close()
	if err != nil {
		return empty, err
	}
	return dimensions, nil
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
