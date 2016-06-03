package gnuplot

import (
	"github.com/yassu/gnuplot.go/conf"
	"testing"
	// "fmt"
	// "math"
)

// Plotter
func TestNewPlotter(t *testing.T) {
	NewPlotter()
}

func TestConfigure2(t *testing.T) {
	p := NewPlotter()
	p.Configure(conf.WithConf())
	if len(p.configures) != 1 {
		t.Errorf("fails in TestConfigure2")
	}
}

// // Function2d
// func TestFunction2dSplitNum(t *testing.T) {
// 	fun := new(Function2d)
// 	fun.Init()
// 	if fun.splitNum != 1000 {
// 		t.Errorf("fails in TestFunction2dSplitNum")
// 	}
// }
//
// func TestFunction2dUpdatePlotter(t *testing.T) {
// 	fun := new(Function2d)
// 	fun.Init()
//
// 	p := new(Plotter)
// 	p.Init()
// 	p.Configure("xMin", "12345")
// 	fun.UpdatePlotter(p)
// 	if fun.plotter.GetC("xMin") != "12345" {
// 		t.Errorf("fails in TestFunction2dUpdatePlotter")
// 	}
// }
//
// // func TestGetGnuData(t *testing.T){
// //     // without panic
// //     fun := new(Function2d)
// //     fun.Init()
// //
// //     f := func(x float64) float64 {
// //         return x * x
// //     }
// //     fun.SetF(f)
// //     fmt.Println(fun.getGnuData())
// // }
//
// // Curve2d
// func TestCurve2dSplitNum(t *testing.T) {
// 	c := new(Curve2d)
// 	c.Init()
// 	if c.splitNum != 100 {
// 		t.Errorf("fails in TestFunction2dSplitNum")
// 	}
// }
//
// func TestCurve2dUpdatePlotter(t *testing.T) {
// 	c := new(Curve2d)
// 	c.Init()
//
// 	p := new(Plotter)
// 	p.Init()
// 	p.Configure("tMin", "12345")
// 	c.UpdatePlotter(p)
// 	if c.plotter.GetC("tMin") != "12345" {
// 		t.Errorf("fails in TestCurve2dUpdatePlotter")
// 	}
// }
