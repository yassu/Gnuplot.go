package main

import (
	// "fmt"
	"github.com/yassu/gnuplot.go"
)

func main() {
	fun := gnuplot.NewFunction2d()
	fun.SetF(func(x float64) float64 {
		return x * x
	})
	fun.Configure("w", "l")
	fun.Configure("lt", "-1")

	c := gnuplot.NewCurve2d()
	c.SetC(func(t float64) [2]float64 {
		return [2]float64{t, -t * t}
	})
	c.Configures(map[string]string{"w": "l"})

	c_plotter := gnuplot.NewPlotter()
	c.UpdatePlotter(c_plotter)

	graph := gnuplot.NewGraph2d()
	graph.Configures(map[string]string{
		"xrange": "[-100:100]",
		"yrange": "[-100:100]"})
	graph.AppendFunc(*fun)
	graph.AppendCurve(*c)
	graph.Run()
}
