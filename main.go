package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"strings"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 81920, 81920
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < (height/2 + 1); py++ {
		progress := py * 100 / height
		if (py % 10) == 0 {
			fmt.Printf("\rProgress: [%-50s] %d%%", strings.Repeat("#", progress/2), progress)
		}
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {

			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			mandelbrot := mandelbrot(z)
			img.Set(px, py, mandelbrot)
			img.Set(px, height-py, mandelbrot)
		}
	}
	f, err := os.Create("mandelbrot.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 500
	const contrast = 3855

	var v complex128
	maxc, minc, midc := uint16(65535), uint16(0), uint16(32766)
	for n := uint16(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := maxc - contrast*n
			g := midc + contrast*n
			b := minc + contrast*n
			return color.RGBA64{r, g, b, 65535}
		}
	}
	return color.Black
}
