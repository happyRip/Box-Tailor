package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/lidded"
	"github.com/happyRip/Box-Tailor/plotter"
)

func main() {
	p := lidded.NewEmptyProduct()
	p.ProcessUserInput()

	outputFile, err := plotter.NewPltFile(p.Name(), "", "2\n1\n0\n")
	if err != nil {
		log.Fatal(err)
	}
	outputFile.Initialize()

	box := lidded.NewBox(
		p,
		utility.NewTriad(0, 0, 0),
		"",
		5,
	)

	outputFile.WriteString(box.Draw())

	o, err := filepath.Abs(outputFile.Path() + outputFile.Name() + ".plt")
	if err != nil {
		log.Fatal(err)
	}

	err = outputFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Plik zapisano pod ścieżką:")
	fmt.Println(o)
}
