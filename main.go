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

const THK = 6

func main() {
	p := box.Product{}
	p.ProcessUserInput()

	outputFile, err := plotter.NewPltFile(p.Name+"_0", "", "")
	if err != nil {
		log.Fatal(err)
	}
	outputFile.Initialize()

	var draft box.Drafter
	draft = lidded.Box{
		Content:        p,
		Margin:         utility.Triad{},
		BoardThickness: THK,
	}

	for _, s := range draft.Draw() {
		outputFile.WriteString(s)
	}

	o, err := filepath.Abs(outputFile.Pointer.Name())
	if err != nil {
		log.Fatal(err)
	}

	err = outputFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	outputFile2, err := plotter.NewPltFile(p.Name+"_1", "", "")
	if err != nil {
		log.Fatal(err)
	}

	p.Size = utility.Triad{
		X: p.Size.X + 2*(THK+1),
		Y: p.Size.Y + 2*(THK+1),
		Z: p.Size.Z,
	}
	draft = lidded.Box{
		Content:        p,
		Margin:         utility.Triad{},
		BoardThickness: THK,
	}
	for _, s := range draft.Draw() {
		outputFile2.WriteString(s)
	}
	o2, err := filepath.Abs(outputFile2.Pointer.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plik zapisano pod ścieżką:")
	fmt.Println(o)
	fmt.Println(o2)
}
