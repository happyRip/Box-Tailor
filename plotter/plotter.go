package plotter

import (
	"bufio"
	"errors"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/happyRip/Box-Tailor/box/utility"
	u "github.com/happyRip/Box-Tailor/plotter/utility"
)

const unit = 40 // 40 points per mm

type pen struct {
	x, y int // current position
}

func NewPen() pen {
	return pen{}
}

func (p *pen) MoveAbsolute(x, y float64) string {
	p.SetX(x)
	p.SetY(y)
	return ConstructCommand(
		"PU",
		p.X(),
		p.Y(),
	)
}

func (p *pen) MoveRelative(x, y float64) string {
	p.AddToX(x)
	p.AddToY(y)
	return ConstructCommand(
		"PU",
		p.X(),
		p.Y(),
	)
}

func (p *pen) Line(x, y float64) string {
	p.AddToX(x)
	p.AddToY(y)
	return ConstructCommand("PD", p.X(), p.Y())
}

func (p *pen) LineShape(points ...[2]float64) string {
	var out string
	for _, point := range points {
		x, y := point[0], point[1]
		out += p.Line(x, y)
	}
	return out
}

func (p *pen) DrawRectangle(width, height float64) string {
	var rect string
	for i := 0; i < 2; i++ {
		rect += p.Line(width, 0)
		rect += p.Line(0, height)
		width *= -1
		height *= -1
	}
	return rect
}

func (p *pen) SetX(f float64) {
	p.x = u.FloatToIntTimesTen(f)
}

func (p *pen) AddToX(f float64) {
	p.x += u.FloatToIntTimesTen(f)
}

func (p *pen) SetY(f float64) {
	p.y = u.FloatToIntTimesTen(f)
}

func (p *pen) AddToY(f float64) {
	p.y += u.FloatToIntTimesTen(f)
}

func (p pen) X() float64 {
	return u.IntSingleDecimalToFloat(p.x) * unit
}

func (p pen) Y() float64 {
	return u.IntSingleDecimalToFloat(p.y) * unit
}

func SelectPen(i int) string {
	return ConstructCommand("SP", float64(i))
}

func ConstructCommand(command string, args ...float64) string {
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

type pltFile struct {
	name, path string
	file       *os.File
	content    string
}

func NewPltFile(name, path, content string) (pltFile, error) {
	var p pltFile
	p.SetName(name)
	p.SetPath(path)
	p.SetContent(content)
	ext := ".plt"

	var err error
	p.file, err = os.Create(p.path + p.name + ext)
	if err != nil {
		return pltFile{}, err
	}
	return p, nil
}

func NewEmptyPltFile() pltFile {
	return pltFile{}
}

func (p pltFile) Close() error {
	err := p.file.Close()
	return err
}

func (p pltFile) Initialize() error {
	_, err := p.file.WriteString("IN;\nLT;\n")
	return err
}

func (p pltFile) WriteString(s string) error {
	_, err := p.file.WriteString(s)
	return err
}

func (p pltFile) WriteContent() error {
	_, err := p.file.WriteString(p.content)
	return err
}

func (p *pltFile) SetName(name string) {
	p.name = utility.TrimExtension(name)
}

func (p *pltFile) SetPath(path string) {
	p.path = path
}

func (p *pltFile) SetContent(content string) {
	p.content = content
}

func (p *pltFile) AppendContent(content string) {
	p.content += content
}

func (p *pltFile) EmptyContent() {
	p.content = ""
}

func (p pltFile) Name() string {
	return p.name
}

func (p pltFile) Path() string {
	return p.path
}

func (p pltFile) Content() string {
	return p.content
}

func (p pltFile) File() *os.File {
	return p.file
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
