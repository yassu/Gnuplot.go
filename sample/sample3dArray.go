package main

import (
	"github.com/yassu/gnup"
)

func main() {
	array := gnuplot.NewArray3d()
	array.Configure("w", []string{"lines"})
	array.Append([3]float64{0, 0, 2})
	array.Append([3]float64{0, 0, 0})
	array.Append([3]float64{0, 2, 0})
	array.Append([3]float64{0, 2, 2})
	array.Append([3]float64{0, 0, 2})
	array.Append([3]float64{0, 1, 4})
	array.Append([3]float64{0, 2, 2})

	graph := gnuplot.NewGraph3d()
	graph.Configure("key", []string{"false"})
	graph.AppendPElem(*array)
	graph.Run()
}
