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
        "_xMin": "-10.0",
        "_xMax": "10.0",
        "_yMin": "-10.0",
        "_yMax": "10.0"}
}

func (fun *Function2d) getData() [][2]float64 { // TODO: テスト書く
    xMin, _ := strconv.ParseFloat(fun.plotter.Configures["_xMin"], 32)
    xMax, _ := strconv.ParseFloat(fun.plotter.Configures["_xMax"], 32)
    yMin, _ := strconv.ParseFloat(fun.plotter.Configures["_yMin"], 32)
    yMax, _ := strconv.ParseFloat(fun.plotter.Configures["_yMax"], 32)
    var sep = float64(xMax - xMin) / float64(fun.splitNum - 1)

    var a [][2]float64
    for j := 0; j < fun.splitNum; j++ {
        t := xMin + float64(j) * sep
        y := fun.f(t)
        if yMin <= y && y <= yMax {
            a = append(a, [2]float64{t, y})
        }
    }
    return a
}

var DefaultCurve2dSplitNum int = 100
type Curve2d struct {
    plotter Plotter
    splitNum int
    c func(float64) [2]float64
}

func (c *Curve2d) init(){
    c.splitNum = DefaultCurve2dSplitNum
    c.plotter.Configures = map[string] string {
        "_tMin": "-10.0",
        "_tMax": "10.0"}
}

func (c *Curve2d) getData() [][2]float64 { // TODO: test
    tMin, _ := strconv.ParseFloat(c.plotter.Configures["_tMin"], 32)
    tMax, _ := strconv.ParseFloat(c.plotter.Configures["_tMax"], 32)
    var sep = float64(tMax - tMin) / float64(c.splitNum - 1)

    var a [][2]float64
    for j := 0; j < c.splitNum; j++ {
        var t float64 = tMin + float64(j) * sep
        cs := c.c(tMin + t * float64(j))
        x0 := cs[0]
        y0 := cs[1]
        a = append(a, [2]float64{x0, y0})
    }
    return a
}

// Graph
type Graph2d struct {
    plotter Plotter
    functions []Function2d
    curves []Curve2d
}

func (g *Graph2d)run() {
    // それぞれのfunctionのdataをtempファイルに書き込む
    // また, それらのファイルの名前を func_filenames []string に格納する

    // それぞれのcurveのdataをtempファイルに書き込む
    // また, それらのファイルの名前を curve_filenames []stringに格納する

    // 実行するgnuplotの実行ファイルをtempファイルに書き込む
    // gnuplotのプログラムを実行する
}
