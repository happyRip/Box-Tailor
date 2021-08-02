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

func main() {
	p := box.Product{}
	p.ProcessUserInput()

	outputFile, err := plotter.NewPltFile(p.Name, "", "2\n1\n0\n")
	if err != nil {
		log.Fatal(err)
	}
	outputFile.Initialize()

	box := lidded.Box{
		Content:        p,
		Margin:         utility.Triad{},
		BoardThickness: 5,
	}

	for _, s := range box.Draw() {
		outputFile.WriteString(s)
	}

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
