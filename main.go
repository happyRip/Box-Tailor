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

	outputFile, err := plotter.NewPltFile(p.Name, "", "")
	if err != nil {
		log.Fatal(err)
	}
	outputFile.Initialize()

	var draft box.Drafter
	draft = lidded.Lid{
		Content:        p,
		Margin:         utility.Triad{},
		BoardThickness: 5,
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
	fmt.Println("Plik zapisano pod ścieżką:")
	fmt.Println(o)
}
