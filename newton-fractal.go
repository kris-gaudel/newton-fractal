package newton_fractal

import (
	"image"
	"image/color"
	"math/cmplx"
)

var Palette = []color.Color{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{85, 65, 95, 255},
	color.RGBA{100, 105, 100, 255},
	color.RGBA{215, 115, 85, 255},
	color.RGBA{80, 140, 215, 255},
	color.RGBA{100, 185, 100, 255},
	color.RGBA{230, 200, 110, 255},
	color.RGBA{220, 245, 255, 255},
}

// NewtonFunc(x) returns the step computed by Newton's method
func NewtonFunc(x complex128) complex128 {
	// The polynomial equation for the fractal is x^8 - 1 = 0
	xSeventh := x * x * x * x * x * x * x
	xEigth := xSeventh * x // d/dx(x^8 - 1) = 8x^7

	// Implementing Newton's method: xn - f(xn)/f'(xn)
	return x - (xEigth-1)/(8.0*xSeventh)
}

/*	GenerateFractal(bottomLeft, topRight) generates a 960 by 540 image of the fractal within a "space" on the complex plane
	specified by a bottom left point and a top right point
*/
func GenerateFractal(bottomLeft complex128, topRight complex128) *image.RGBA {
	realStep := (real(topRight) - real(bottomLeft)) / float64(640)
	imagStep := (imag(topRight) - imag(bottomLeft)) / float64(480)

	out := image.NewRGBA(image.Rect(0, 0, 640, 480))

	for i := 0; i < 640; i++ {
		rc := real(bottomLeft) + realStep*float64(i)

		for j := 0; j < 480; j++ {
			ic := imag(bottomLeft) + imagStep*float64(j)
			color := getColor128(complex(rc, ic))
			out.Set(i, j, color)
		}
	}

	return out
}

func getColor128(p complex128) color.Color {

	roots := [8]complex128{
		1.0 + 0.0i,
		-1.0 + 0.0i,
		0.0 + 1.0i,
		0.0 - 1.0i,
		0.70710678118 + 0.70710678118i,
		-0.70710678118 - 0.70710678118i,
		0.70710678118 - 0.70710678118i,
		-0.70710678118 + 0.70710678118i,
	}

	const epsilon = 0.00001

	pv := NewtonFunc(p)

	for iterations := uint(1); iterations < 100; iterations++ {
		switch {
		case cmplx.Abs(pv-roots[0]) < epsilon:
			return Palette[0]
		case cmplx.Abs(pv-roots[1]) < epsilon:
			return Palette[1]
		case cmplx.Abs(pv-roots[2]) < epsilon:
			return Palette[2]
		case cmplx.Abs(pv-roots[3]) < epsilon:
			return Palette[3]
		case cmplx.Abs(pv-roots[4]) < epsilon:
			return Palette[4]
		case cmplx.Abs(pv-roots[5]) < epsilon:
			return Palette[5]
		case cmplx.Abs(pv-roots[6]) < epsilon:
			return Palette[6]
		case cmplx.Abs(pv-roots[7]) < epsilon:
			return Palette[7]
		}
		pv = NewtonFunc(pv)
	}

	return color.RGBA{0, 0, 0, 255}
}
