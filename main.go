package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/happyRip/Box-Tailor/box"
	"github.com/happyRip/Box-Tailor/box/lidded"
	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/plotter"
)

const (
	BOARD_THICKNESS = 6
	// BOTTOM_BOX_STR  = "_0" // denko
	// TOP_BOX_STR     = "_1" // wieczko
	KERF = 2
)

func main() {
	p := box.Product{}
	p.ProcessUserInput()

	bottom := lidded.Box{
		Content:        p,
		BoardThickness: BOARD_THICKNESS,
		Kerf:           KERF,
		Origin: utility.Pair{
			X: 0,
			Y: -12,
		},
	}
	origin, _ := bottom.CalculateSize()
	lid := lidded.Box{
		Content: box.Product{
			Size: utility.Triad{
				X: p.Size.X + 2*(BOARD_THICKNESS+1),
				Y: p.Size.Y + 2*(BOARD_THICKNESS+1) + BOARD_THICKNESS,
				Z: p.Size.Z,
			},
		},
		Origin: utility.Pair{
			X: origin + 5,
			Y: -12,
		},
		BoardThickness: BOARD_THICKNESS,
		Kerf:           KERF,
	}

	outAbs, err := drawToSingleFile(
		p.Name,
		bottom,
		lid,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plik zapisano pod ścieżką:")
	fmt.Println(outAbs)
}

// func drawBox(name string, draft box.Drafter) (string, error) {
// 	outputFile, err := plotter.NewPltFile(name, "", "")
// 	if err != nil {
// 		return "", err
// 	}
// 	outputFile.Initialize()

// 	for _, s := range draft.Draw() {
// 		outputFile.WriteString(s)
// 	}

// 	outputAbs, err := filepath.Abs(outputFile.Pointer.Name())
// 	if err != nil {
// 		return "", err
// 	}

// 	err = outputFile.Close()
// 	return outputAbs, err
// }

func drawToSingleFile(name string, bottom box.Drafter, lid box.Drafter) (string, error) {
	outputFile, err := plotter.NewPltFile(name, "", "")
	if err != nil {
		return "", err
	}
	outputFile.Initialize()

	outputFile.WriteString(
		plotter.SelectPen(5),
		plotter.DefineTerminator(';'),
		plotter.CharacterSize(0.75, 1.5),
		plotter.Label(name),
	)

	for _, s := range bottom.Draw() {
		outputFile.WriteString(s)
	}

	x, _ := bottom.CalculateSize()
	pen := plotter.Pen{}
	outputFile.WriteString(pen.MoveAbsolute(x+1, 0))
	for _, s := range lid.Draw() {
		outputFile.WriteString(s)
	}

	outAbs, err := filepath.Abs(outputFile.Pointer.Name())
	if err != nil {
		return "", err
	}

	err = outputFile.Close()
	return outAbs, err
}
