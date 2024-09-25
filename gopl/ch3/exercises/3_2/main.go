// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 // canvas size in pixels
	cells         = 100      // number of grid cells: 100 * 100
	// xyrange for snow slope, f and saddle
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	// xyrange for egg box
	//xyrange = 10
	xyscale = width / 2 / xyrange // pixels per x or y unit
	// zscale for saddle and f
	zscale = height * 0.4 // pixels per z unit
	// zscale for egg box
	//zscale = height * 0.2
	// zscale for snow slope
	//zscale = height * 0.1
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, oka := corner(i+1, j)
			bx, by, okb := corner(i, j)
			cx, cy, okc := corner(i, j+1)
			dx, dy, okd := corner(i+1, j+1)
			if oka && okb && okc && okd {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	//z := f(x, y)
	// the surface is an egg box
	//z := eggBox(x, y)
	// the surface is a snow slope
	//z := snowSlope(x, y)
	// the surface is a saddle
	z := saddle(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

//func f(x, y float64) float64 {
//	r := math.Hypot(x, y) // distance from (0,0)
//	return math.Sin(r) / r
//}
//
//func eggBox(x, y float64) float64 {
//	return math.Sin(x) * math.Cos(y)
//}
//
//func snowSlope(x, y float64) float64 {
//	return -0.1 * (x + y)
//}

func saddle(x, y float64) float64 {
	return (x*x - y*y) / 10
}

//!-
