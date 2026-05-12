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

func main() {
	p := plot.New()

	p.Title.Text = "Tensão de Entrada"
	p.X.Label.Text = "rad"
	p.Y.Label.Text = "V"

	axis := plotter.NewFunction(func(x float64) float64 { return 0.0 })
	axis.LineStyle = plotter.DefaultLineStyle
	axis.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}

	sin := plotter.NewFunction(func(x float64) float64 { return 50 * math.Sin(x) })
	sin.LineStyle = plotter.DefaultLineStyle
	sin.Color = color.RGBA{R: 255, A: 255}

	p.Add(axis, sin)
	p.Legend.Add("Vin", sin)
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	p.X.Min = 0
	p.X.Max = 10

	p.Y.Min = -100
	p.Y.Max = 100

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "resultado.png"); err != nil {
		panic(err)
	}
}
