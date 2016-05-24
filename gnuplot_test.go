package gnuplot

import (
  "testing"
  // "fmt"
)

func TestConfigure(t *testing.T) {
    p := new(Plotter)
    p.init()
    p.configure("xMin", "3.0")
    if p.getC("xMin") != "3.0" {
        t.Errorf("fails in TestConfigure")
    }
}

func TestConfigure2(t *testing.T) {
    p := new(Plotter)
    p.init()
    if p.getC("xMin") != "-10.0" {
        t.Errorf("fails in TestConfigure2")
    }
}
