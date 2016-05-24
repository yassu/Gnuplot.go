package gnuplot

import (
  // "fmt"
  "strconv"
)

type Plotter struct {
    Configures map[string] string
}

func (p *Plotter) init() {
    p.Configures = map[string] string{}
}

func (p *Plotter) configure(key, val string) {
    p.Configures[key] = val
}

func (p *Plotter) getC(key string) string {
    return p.Configures[key]
}

var DefaultFunction2dSplitNum int = 1000
type Function2d struct {
    plotter Plotter
    splitNum int
    f func(float64) float64
}

func (fun *Function2d) init(){
    fun.splitNum = DefaultFunction2dSplitNum
    fun.plotter.Configures = map[string] string {
        "xMin": "-10.0",
        "xMax": "10.0",
        "yMin": "-10.0",
        "yMax": "10.0"}
}

func (fun *Function2d) getData() [][2]float64 { // TODO: テスト書く
    xMin, _ := strconv.ParseFloat(fun.plotter.Configures["xMin"], 32)
    xMax, _ := strconv.ParseFloat(fun.plotter.Configures["xMax"], 32)
    yMin, _ := strconv.ParseFloat(fun.plotter.Configures["yMin"], 32)
    yMax, _ := strconv.ParseFloat(fun.plotter.Configures["yMax"], 32)
    var sep = float64(xMax - xMin) / float64(fun.splitNum - 1)

    var a [][2]float64
    for j := 0; j < fun.splitNum; j++ {
        var t float64 = xMin + float64(j) * sep
        y := fun.f(xMin + t * float64(j))
        if yMin <= y && y <= yMax {
            a = append(a, [2]float64{t, fun.f(t)})
        }
    }
    return a
}
