package main

import (
	"fmt"
	"github.com/yassu/gnuplot.go"
	"github.com/yassu/gnuplot.go/conf"
)

func main() {
	fun := gnuplot.NewFunction2d()
	fun.SetF(func(x float64) float64 {
		return x * x
	})
	fun.SetConfigures([]*conf.Configure{
		conf.GoXMinConf(),
		conf.GoXMaxConf(),
		conf.WithConf()})
	fun.Configure("_xMin", "-100")
	fun.Configure("_xMax", "100")
	fmt.Println(fun)

	graph := gnuplot.NewGraph2d()
	graph.AppendFunc(*fun)
	graph.Run()
}
