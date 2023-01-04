package main

import (
	"image/png"
	"log"
	"os"

	newton_fractal "github.com/kris-gaudel/newton-fractal"
)

func main() {
	const upper = -1.0 + 1.0i
	const lower = 1.0 - 1.0i

	img := newton_fractal.GenerateFractal(upper, lower)
	ofile, err := os.Create("fractal.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(ofile, img)
}
