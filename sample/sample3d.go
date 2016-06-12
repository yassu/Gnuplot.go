package main

import (
	"github.com/yassu/gnup"
)

func main() {
	fun := gnuplot.NewFunction3d()
	fun.SetF(func(x, y float64) float64 {
		return x*x + y*y
	})
	fun.Configure("w", []string{"lines"})

	c := gnuplot.NewCurve3d()
	c.SetC(func(t float64) [3]float64 {
		return [3]float64{t, t * t, 0}
	})
	c.Configure("w", []string{"lines"})

	graph := gnuplot.NewGraph3d()
	graph.Configure("angles", []string{"degrees"})
	graph.Configure("key", []string{"false"})
	graph.AppendPElem(*fun)
	graph.AppendPElem(*c)
	graph.Run()
}
