package gnuplot

import (
	"fmt"
	"github.com/yassu/gnuplot.go/conf"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Plotter
type Plotter struct {
	configures []*conf.Configure
}

func NewPlotter() *Plotter {
	plotter := new(Plotter)
	return plotter
}

func (p *Plotter) Configure(conf *conf.Configure) {
	for j := range p.configures {
		if p.configures[j].GetKey() == conf.GetKey() {
			p.configures[j].SetVal(conf.GetVal())
			return
		}
	}
	p.configures = append(p.configures, conf)
}

func (p *Plotter) GetC(key string) string {
	for j := range p.configures {
		if p.configures[j].GetKey() == key {
			return p.configures[j].GetVal()
		}
	}
	return ""
}

// Function2d
const DefaultFunction2dSplitNum int = 1000

type Function2d struct {
	plotter  Plotter
	splitNum int
	f        func(float64) float64
}

func NewFunction2d() *Function2d {
	fun := new(Function2d)
	fun.splitNum = DefaultFunction2dSplitNum
	fun.setConfigure()
	return fun
}

func (fun *Function2d) setConfigure() {
	for _, conf := range conf.Function2dConfs() {
		fun.plotter.Configure(conf)
	}
}

func (fun *Function2d) Configure(key, val string) {
	for j, conf := range fun.plotter.configures {
		if conf.GetKey() == key {
			fun.plotter.configures[j].SetVal(val)
			return
		}
	}
	panic(fmt.Sprintf("%v is not a key.", key))
}

func (fun *Function2d) Configures(sconf map[string]string) {
	for key, val := range sconf {
		fun.Configure(key, val)
	}
}

func (fun *Function2d) GetData() [][2]float64 { // TODO: テスト書く
	xMin, _ := strconv.ParseFloat(fun.plotter.GetC("_xMin"), 32)
	xMax, _ := strconv.ParseFloat(fun.plotter.GetC("_xMax"), 32)
	var sep = float64(xMax-xMin) / float64(fun.splitNum-1)

	var a [][2]float64
	for j := 0; j < fun.splitNum; j++ {
		t := xMin + float64(j)*sep
		y := fun.f(t)
		a = append(a, [2]float64{t, y})
	}
	return a
}

func (fun *Function2d) getGnuData() string {
	var s string
	for _, xs := range fun.GetData() {
		s += fmt.Sprintf("%f %f\n", xs[0], xs[1])
	}
	return s
}

func (fun *Function2d) SetF(_f func(float64) float64) {
	fun.f = _f
}

func (fun Function2d) gnuplot(filename string) string {
	var s = fmt.Sprintf("\"%v\"", filename)
	for _, conf := range fun.plotter.configures {
		if !strings.HasPrefix(conf.GetKey(), "_") {
			s += fmt.Sprintf(" %v %v", conf.GetKey(), conf.GetVal())
		}
	}
	return s
}

func (fun *Function2d) writeIntoGnufile(f os.File) {
	f.WriteString(fun.getGnuData())
}

// Curve2d
const DefaultCurve2dSplitNum int = 1000

type Curve2d struct {
	plotter  Plotter
	splitNum int
	c        func(float64) [2]float64
}

func NewCurve2d() *Curve2d {
	c := new(Curve2d)
	c.splitNum = DefaultCurve2dSplitNum
	c.setConfigure()
	// c.setConfigure()
	// c.plotter.configures = map[string]string{
	// 	"_tMin": "-10.0",
	// 	"_tMax": "10.0"}
	return c
}

func (c *Curve2d) setConfigure() {
	for _, conf := range conf.Curve2dConfs() {
		c.plotter.Configure(conf)
	}
}

func (c *Curve2d) Configure(key, val string) {
	for j, conf := range c.plotter.configures {
		if conf.GetKey() == key {
			c.plotter.configures[j].SetVal(val)
			return
		}
	}
	panic(fmt.Sprintf("%v is not a key.", key))
}

func (c *Curve2d) Configures(sconf map[string]string) {
	for key, val := range sconf {
		c.Configure(key, val)
	}
}

func (c *Curve2d) GetData() [][2]float64 { // TODO: test
	tMin, _ := strconv.ParseFloat(c.plotter.GetC("_tMin"), 32)
	tMax, _ := strconv.ParseFloat(c.plotter.GetC("_tMax"), 32)
	var sep = float64(tMax-tMin) / float64(c.splitNum-1)

	var a [][2]float64
	for j := 0; j < c.splitNum; j++ {
		cs := c.c(tMin + float64(j)*sep)
		a = append(a, [2]float64{cs[0], cs[1]})
	}
	return a
}

func (c *Curve2d) getGnuData() string {
	var s string
	for _, xs := range c.GetData() {
		s += fmt.Sprintf("%f %f\n", xs[0], xs[1])
	}
	return s
}

func (c *Curve2d) SetC(_c func(float64) [2]float64) {
	c.c = _c
}

func (c Curve2d) gnuplot(fileName string) string {
	var s = fmt.Sprintf("\"%v\" ", fileName)
	for _, conf := range c.plotter.configures {
		if !strings.HasPrefix(conf.GetKey(), "_") {
			s += fmt.Sprintf(" %v %v", conf.GetKey(), conf.GetVal())
		}
	}
	return s
}

// Graph
type Graph2d struct {
	plotter   Plotter
	functions []Function2d
	curves    []Curve2d
}

func NewGraph2d() *Graph2d {
	g := new(Graph2d)
	g.setConfigure()
	return g
}

func (g *Graph2d) setConfigure() {
	for _, conf := range conf.Graph2dConfs() {
		g.plotter.Configure(conf)
	}
}

func (g *Graph2d) Configure(key, val string) {
	for j, conf := range g.plotter.configures {
		if conf.GetKey() == key {
			g.plotter.configures[j].SetVal(val)
			return
		}
	}
	panic(fmt.Sprintf("%v is not a key.", key))
}

func (g *Graph2d) Configures(sconf map[string]string) {
	for key, val := range sconf {
		g.Configure(key, val)
	}
}

func (g *Graph2d) AppendFunc(f Function2d) {
	g.functions = append(g.functions, f)
}

func (g *Graph2d) AppendCurve(c Curve2d) {
	g.curves = append(g.curves, c)
}

func (g Graph2d) writeIntoFile(data string, f *os.File) {
	f.WriteString(data)
}

func (g Graph2d) gnuplot(funcFilenames []string, curveFilenames []string) string {
	var s string

	for _, conf := range g.plotter.configures {
		if !strings.HasPrefix(conf.GetKey(), "_") {
			if conf.GetVal() == "true" {
				s += fmt.Sprintf("set %v;\n", conf.GetKey())
			} else if conf.GetVal() == "false" {
				s += fmt.Sprintf("set no%v;\n", conf.GetKey())
			} else {
				s += fmt.Sprintf("set %v %v;\n", conf.GetKey(), conf.GetVal())
			}
		}
	}

	s += "plot "
	for j, _ := range g.functions {
		s += g.functions[j].gnuplot(funcFilenames[j]) + ", "
	}
	for j, _ := range g.curves {
		s += g.curves[j].gnuplot(curveFilenames[j])
		if j != len(g.curves)-1 {
			s += ", "
		}
	}
	s += ";\n"
	s += "pause -1;\n"
	return s
}

func (g *Graph2d) Run() {
	tmpDir := os.TempDir() + "/gnuplot.go/"
	// TODO: tmpDirがなければ作る
	// execFilename := tmpDir + "exec.gnu"
	execFilename := "exec.gnu"

	// それぞれのfunctionのdataをtempファイルに書き込む
	// また, それらのファイルの名前を func_filenames []string に格納する
	var funcFilenames []string
	for _, fun := range g.functions {
		file, err := ioutil.TempFile(tmpDir, "")
		defer func() {
			file.Close()
		}()
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		} else {
			g.writeIntoFile(fun.getGnuData(), file)
			funcFilenames = append(funcFilenames, file.Name())
		}
	}

	// それぞれのcurveのdataをtempファイルに書き込む
	// また, それらのファイルの名前を curve_filenames []stringに格納する
	var curveFilenames []string
	for _, c := range g.curves {
		file, _ := ioutil.TempFile(tmpDir, "")
		defer func() {
			file.Close()
		}()
		g.writeIntoFile(c.getGnuData(), file)
		curveFilenames = append(curveFilenames, file.Name())
	}

	// 実行するgnuplotの実行ファイルをtempファイルに書き込む
	os.Remove(execFilename)
	execFile, _ := os.OpenFile(execFilename, os.O_CREATE|os.O_WRONLY, 0666)
	defer func() {
		execFile.Close()
	}()
	execFile.WriteString(g.gnuplot(funcFilenames, curveFilenames))
}
