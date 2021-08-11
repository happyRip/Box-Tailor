package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/happyRip/Box-Tailor/box"
	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/lidded"
	"github.com/happyRip/Box-Tailor/plotter"
)

const (
	BOARD_THICKNESS = 6
	BOTTOM_BOX_STR  = "_0" // denko
	TOP_BOX_STR     = "_1" // wieczko
)

func main() {
	p := box.Product{}
	p.ProcessUserInput()

	var outAbs []string
	o, err := drawBox(
		p.Name+BOTTOM_BOX_STR,
		lidded.Box{
			Content:        p,
			BoardThickness: BOARD_THICKNESS,
		})
	if err != nil {
		log.Fatal(err)
	}
	outAbs = append(outAbs, o)

	o, err = drawBox(
		p.Name+TOP_BOX_STR,
		lidded.Box{
			Content: box.Product{
				Size: utility.Triad{
					X: p.Size.X + 2*(BOARD_THICKNESS+1),
					Y: p.Size.Y + 2*(BOARD_THICKNESS+1),
					Z: p.Size.Z,
				},
			},
			BoardThickness: BOARD_THICKNESS,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	outAbs = append(outAbs, o)

	fmt.Println("Pliki zapisano pod ścieżką:")
	for _, s := range outAbs {
		fmt.Println(s)
	}
}

func drawBox(name string, draft box.Drafter) (string, error) {
	outputFile, err := plotter.NewPltFile(name, "", "")
	if err != nil {
		return "", err
	}
	outputFile.Initialize()

	for _, s := range draft.Draw() {
		outputFile.WriteString(s)
	}

	outputAbs, err := filepath.Abs(outputFile.Pointer.Name())
	if err != nil {
		return "", err
	}

	err = outputFile.Close()
	return outputAbs, err
}
