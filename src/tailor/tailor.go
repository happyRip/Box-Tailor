package tailor

import (
	"bufio"
	"errors"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

type File struct {
	Name, Directory, Extention string
}

func (f *File) Initialize(path string) {
	f.Extention = filepath.Ext(path)
	f.Name = filepath.Base(path)
	f.Name = f.Name[0 : len(f.Name)-len(f.Extention)] // remove extension from file name
	f.Directory = filepath.Dir(path)
	if f.Directory == "." || f.Directory == "/" {
		f.Directory = "./"
	}
}

type dimensions struct {
	X, Y, Z float32
}

type Product struct {
	Source File
	Size   dimensions
}

func (p Product) GetDimensions(sizeZ float32) error {
	if p.Source.Name == "" {
		return errors.New("product source path not specified")
	} else if p.Source.Extention != ".plt" {
		return errors.New("file extension is incorrect")
	}

	file, err := os.Open(p.Source.Directory)
	if err != nil {
		file.Close()
		return err
	}

	ext := extremes{
		min: point{x: math.MaxFloat32, y: math.MaxFloat32},
		max: point{x: -math.MaxFloat32, y: -math.MaxFloat32},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == 'P' {
			if line[1] == 'D' {
				for i, v := range getNumbers(line) {
					f, err := strconv.ParseFloat(v, 32)
					if err != nil {
						file.Close()
						return err
					}

					if i%2 == 0 {
						ext.min.x = min(ext.min.x, float32(f))
						ext.max.x = max(ext.max.x, float32(f))
					} else {
						ext.min.y = min(ext.min.y, float32(f))
						ext.max.y = max(ext.max.y, float32(f))
					}
				}
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return err
	}
	p.Size.X, p.Size.Y, p.Size.Z = (ext.max.x-ext.min.x)/unit, (ext.max.y-ext.min.y)/unit, sizeZ

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
