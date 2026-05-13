package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func sine(x float64) float64 {
	return math.Sin(x * frequencia)
}

func triangular(x float64) float64 {
	fTriang := 10000.0
	// tTriang := 0.0001
	return (2.0 / math.Pi) * math.Abs(math.Asin(math.Sin(2.0*math.Pi*fTriang*x)))
}

func firstGates(x float64) float64 {
	if sine(x) > 0 {
		if sine(x) > triangular(x) {
			return 1.0
		} else {
			return 0.0
		}
	}
	return 0.0
}

func secondGates(x float64) float64 {
	if sine(x) < 0 {
		if -sine(x) > triangular(x) {
			return 1.0
		} else {
			return 0.0
		}
	}
	return 0.0
}

var frequencia float64

func main() {
	p := plot.New()

	p.Title.Text = "Tensão e Chaveamento"
	p.X.Label.Text = "rad"
	p.Y.Label.Text = "V"

	fmt.Print("Digite a frequência de operação do sinal base (Hz): ")
	fmt.Scanln(&frequencia)

	// Plot do eixo X
	axis := plotter.NewFunction(func(x float64) float64 { return 0.0 })
	axis.LineStyle = plotter.DefaultLineStyle
	axis.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}

	// Plot da função seno com frequência de operação do sinal base determinada
	sin := plotter.NewFunction(func(x float64) float64 { return sine(x) })
	sin.Color = color.RGBA{R: 255, A: 255}

	// Plot das chaves 1 e 4
	upperGates := plotter.NewFunction(func(x float64) float64 { return firstGates(x) })
	upperGates.Color = color.RGBA{G: 255, A: 255}

	// Plot das chaves 2 e 3
	lowerGates := plotter.NewFunction(func(x float64) float64 { return secondGates(x) })
	lowerGates.Color = color.RGBA{B: 255, A: 255}

	// Plot da funçao triangular utilizada para fins de debug
	triang := plotter.NewFunction(func(x float64) float64 { return triangular(x) })
	triang.Color = color.RGBA{G: 200, B: 200, A: 255}

	sin.Samples, triang.Samples, upperGates.Samples, lowerGates.Samples = 10000, 10000, 10000, 10000

	p.Add(axis, upperGates, lowerGates, sin)
	// p.Add(axis, triang)

	p.Legend.Add(fmt.Sprintf("%.2fHz", frequencia), sin)
	p.Legend.Add("chaves 1, 4", upperGates)
	p.Legend.Add("chaves 2, 3", lowerGates)
	p.Legend.ThumbnailWidth = 1 * vg.Centimeter

	p.X.Min = 0.0
	p.X.Max = (2.5 * math.Pi) / frequencia

	p.Y.Min = -1.0
	p.Y.Max = 1.0

	// Salvando a imagem de 15x10cm com nome 'resultado.png'
	if err := p.Save(15*vg.Centimeter, 10*vg.Centimeter, "resultado.png"); err != nil {
		panic(err)
	}
}
