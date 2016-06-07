package main

import (
	"github.com/yassu/gnuplot.go"
)

func main() {
	fun := gnuplot.NewFunction2d()
	fun.SetF(func(x float64) float64 {
		return x * x
	})
	fun.Configure("_xMin", []string{"-100"})
	fun.Configure("_xMax", []string{"100"})
	fun.Configure("_title", []string{"title1"})
	fun.Configure("w", []string{"dots"})

	c := gnuplot.NewCurve2d()
	c.SetC(func(t float64) [2]float64 {
		return [2]float64{t, -t * t}
	})
	c.Configures(map[string][]string{
		"_tMin": []string{"-100"},
		"_tMax": []string{"100"}})

	graph := gnuplot.NewGraph2d()
	graph.Configure("angles", []string{"degrees"})
	graph.Configure("key", []string{"false"})
	graph.AppendFunc(*fun)
	graph.AppendCurve(*c)
	graph.Run()
}
