package gnuplot

import (
	"github.com/yassu/gnup/conf"
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

// Function2d
func TestFunction2dSplitNum(t *testing.T) {
	fun := NewFunction2d()
	if fun.splitNum != 1000 {
		t.Errorf("fails in TestFunction2dSplitNum")
	}
}

func TestGetGnuData(t *testing.T) {
	fun := NewFunction2d()

	f := func(x float64) float64 {
		return x * x
	}
	fun.SetF(f)
	fun.getGnuData()
}

// Curve2d
func TestNewCurve2d(t *testing.T) {
	NewCurve2d()
}

func TestCurve2dSplitNum(t *testing.T) {
	c := NewCurve2d()
	if c.splitNum != 1000 {
		t.Errorf("fails in TestFunction2dSplitNum")
	}
}
