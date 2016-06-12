gnup
==============

*Version: 0.0.1, Development Status: Beta*

Caller from golang to gnuplot

Install
---------

``` bash
$ go get -v github.com/yassu/gnup
```

Introduction
--------------

Feature of `gnup` is a wrapper of gnuplot for golang.
However, now this is supported only 1 variable functions or curves in a plane.

This project will suport 2 variable functions or surfaces in 3 dimension space.

Structure of the source of this project is very simple like as

``` go
package main

import (
	"github.com/yassu/gnup"
)

func main() {
	fun := gnup.NewFunction2d()
	fun.SetF(func(x float64) float64 {
		return x * x
	})
	fun.Configure("_xMin", []string{"-100"})
	fun.Configure("_xMax", []string{"100"})
	fun.Configure("_title", []string{"title1"})
	fun.Configure("w", []string{"dots"})

	c := gnup.NewCurve2d()
	c.SetC(func(t float64) [2]float64 {
		return [2]float64{t, -t * t}
	})
	c.Configures(map[string][]string{
		"_tMin": []string{"-100"},
		"_tMax": []string{"100"}})

	graph := gnup.NewGraph2d()
	graph.Configure("angles", []string{"degrees"})
	graph.Configure("key", []string{"false"})
	graph.AppendPElem(*fun)
	graph.AppendPElem(*c)
	graph.Run()
}
```

This code makes script for `gnuplot`, called `exec.gnu` (If there don't exist,
please make `/tmp/gnup/` directory).

Then, when you enter `gnuplot exec.gnu`, you obtain output.

Tasks
-------
- [ ] write samples by Using wiki
- [ ] write all of validations of configures of Function2d or Curve2d
- [x] support of Function3d
- [x] support of Curve3d
- [x] support of Surface3d

LICENSE
---------

MIT License

AUTHOR
----------

yassu
