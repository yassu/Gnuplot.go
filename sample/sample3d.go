package main

import (
	"github.com/yassu/gnup"
)

func main() {
	fun := gnuplot.NewFunction3d()
	fun.SetF(func(x, y float64) float64 {
		return x*x + y*y
	})
	fun.Configure("_title", []string{"title1"})
	fun.Configure("w", []string{"dots"})

	graph := gnuplot.NewGraph3d()
	graph.Configure("angles", []string{"degrees"})
	graph.Configure("key", []string{"false"})
	graph.AppendPElem(*fun)
	graph.Run()
}
