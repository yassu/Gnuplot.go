package gnuplot

import (
  "testing"
  // "fmt"
  // "math"
)

// Plotter
func TestConfigure(t *testing.T) {
    p := new(Plotter)
    p.Init()
    p.Configure("xMin", "3.0")
    if p.GetC("xMin") != "3.0" {
        t.Errorf("fails in TestConfigure")
    }
}

// Function2d
func TestFunction2dSplitNum(t *testing.T) {
    fun := new(Function2d)
    fun.Init()
    if fun.splitNum != 1000 {
        t.Errorf("fails in TestFunction2dSplitNum")
    }
}

// Curve2d
func TestCurve2dSplitNum(t *testing.T) {
    c := new(Curve2d)
    c.Init()
    if c.splitNum != 100 {
        t.Errorf("fails in TestFunction2dSplitNum")
    }
}
