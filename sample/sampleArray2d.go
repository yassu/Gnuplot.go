package main

import (
	"github.com/yassu/gnup"
)

func main() {
	array := gnuplot.NewArray2d()
	array.Append([2]float64{0, 2})
	array.Append([2]float64{0, 0})
	array.Append([2]float64{2, 0})
	array.Append([2]float64{2, 2})
	array.Append([2]float64{0, 2})
	array.Append([2]float64{1, 4})
	array.Append([2]float64{2, 2})
	array.Configure("_title", []string{"title1"})
	array.Configure("w", []string{"lines"})

	graph := gnuplot.NewGraph2d()
	graph.Configure("key", []string{"false"})
	graph.Configure("xrange", []string{"[-1:3]"})
	graph.Configure("yrange", []string{"[-1:5]"})
	graph.AppendPElem(*array)
	graph.Run()
}
