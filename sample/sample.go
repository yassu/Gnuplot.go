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
	fun.Configure("_xMin", []string{"-100"})
	fun.Configure("_xMax", []string{"100"})
	fun.Configure("w", []string{"dots"})

	c := gnuplot.NewCurve2d()
	c.SetC(func(t float64) [2]float64 {
		return [2]float64{t, -t * t}
	})
	c.Configure("_tMin", []string{"-100"})
	c.Configure("_tMax", []string{"100"})

	graph := gnuplot.NewGraph2d()
	graph.Configure("angles", []string{"degrees"})
	graph.AppendFunc(*fun)
	graph.AppendCurve(*c)
	graph.Run()
}
