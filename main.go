package main

import (
	"fmt"

	"github.com/happyRip/Box-Tailor/lidded"
)

func main() {
	p := lidded.NewEmptyProduct()
	p.ProcessUserInput()

	fmt.Println("name:", p.Name())
	fmt.Println("size:", p.Width(), p.Depth(), p.Height())
}
