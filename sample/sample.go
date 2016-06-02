package main

import (
	// "fmt"
	"github.com/yassu/gnuplot.go"
	// "github.com/yassu/gnuplot.go/conf"
)

func main() {
	fun := gnuplot.NewFunction2d()
	fun.SetF(func(x float64) float64 {
		return x * x
	})
	fun.Configure("_xMin", "-100")
	fun.Configure("_xMax", "100")
	fun.Configure("with", "dots")

	c := gnuplot.NewCurve2d()
	c.SetC(func(t float64) [2]float64 {
		return [2]float64{t, -t * t}
	})
	c.Configure("_tMin", "-100")
	c.Configure("_tMax", "100")

	graph := gnuplot.NewGraph2d()
	graph.Configure("angles", "degrees")
	graph.AppendFunc(*fun)
	graph.AppendCurve(*c)
	graph.Run()
}
