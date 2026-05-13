package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

/* Lógica proposta pelo professor
if (math.Sin(x) > 0) {
	if (math.Sin(x) > triang) {
		saida1, saida4 = 1, 1
	} else {
		saida1, saida4 = 0, 0
	}
}
else (math.Sin(x) < 0) {
	if (math.Sin(x) < triang) {
		saida2, saida3 = 1, 1
	} else {
		saida2, saida3 = 0, 0
	}
}
*/

func sine(x float64) float64 {
	return math.Sin(x)
}

func triangular(x float64) float64 {
	return x * 0.1
}

func firstGates(x float64) float64 {
	if sine(x) > 0 {
		if sine(x) > triangular(x) {
			return 1.0
		} else {
			return 0.0
		}
	} else if sine(x) < 0 {
		if sine(x) < triangular(x) {
			return 1.0
		} else {
			return 0.0
		}
	}
	return 0.0
}

func main() {
	p := plot.New()

	p.Title.Text = "Tensão"
	p.X.Label.Text = "rad"
	p.Y.Label.Text = "V"

	axis := plotter.NewFunction(func(x float64) float64 { return 0.0 })
	axis.LineStyle = plotter.DefaultLineStyle
	axis.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}

	// Frequência de operação da senoide: 60Hz
	// fSin := 60.0
	// tSin := 0.016666667
	sin := plotter.NewFunction(func(x float64) float64 { return sine(x) })
	sin.Color = color.RGBA{R: 255, A: 255}

	// Frequência de operação da função triangular de referência: 10000Hz
	// fTriang := 10000.0
	// tTriang := 0.0001
	triang := plotter.NewFunction(func(x float64) float64 { return triangular(x) })
	triang.Color = color.RGBA{B: 255, A: 255}

	upperGates := plotter.NewFunction(func(x float64) float64 { return firstGates(x) })
	upperGates.Color = color.RGBA{G: 255, A: 255}

	p.Add(axis, triang, upperGates, sin)
	p.Legend.Add("Vin", sin)
	p.Legend.Add("1, 4", upperGates)
	p.Legend.Add("2, 3", triang)
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	p.X.Min = 0
	p.X.Max = 10

	p.Y.Min = -1
	p.Y.Max = 1

	// Save the plot to a PNG file.
	if err := p.Save(15*vg.Centimeter, 10*vg.Centimeter, "resultado.png"); err != nil {
		panic(err)
	}
}
