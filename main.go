package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/lidded"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	var dimensions []float64
	for _, v := range args[1:] {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatal(err)
		}
		dimensions = append(dimensions, f)
		log.Println(dimensions)
	}

	name := args[0]
	var size utility.Triad
	size.SetValues(dimensions[0], dimensions[1], dimensions[2])
	product, err := lidded.NewProduct(
		name,
		size,
	)
	if err != nil {
		log.Fatal(err)
	}

	rName := product.Name()
	rX, rY, rZ, _ := product.Size()
	fmt.Println("name:", rName)
	fmt.Println("size:", rX, rY, rZ)
}
