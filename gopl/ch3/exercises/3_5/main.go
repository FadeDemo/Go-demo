// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	err := png.Encode(os.Stdout, img)
	if err != nil {
		return
	} // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	//const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colorizeYCbCr(n)
		}
	}
	return color.Black
}

func colorizeYCbCr(n uint8) color.Color {
	// 根据迭代次数生成 YCbCr 颜色
	y := 255 * n / 200                      // Y 分量（亮度）
	cb := uint8(128 + 127*int(n%20)/20)     // Cb 分量（蓝色色度）
	cr := uint8(128 + 127*int(200-n%20)/20) // Cr 分量（红色色度）

	return color.YCbCr{Y: y, Cb: cb, Cr: cr}
}
