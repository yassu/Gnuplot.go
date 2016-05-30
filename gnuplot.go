package gnuplot

import (
	"fmt"
	"io/ioutil"
	"os"
	// "os/exec"
	"strconv"
	"strings"
)

// Plotter
type Plotter struct {
	configures map[string]string
}

func NewPlotter() *Plotter {
	plotter := new(Plotter)
	plotter.configures = map[string]string{}
	return plotter
}

func (p *Plotter) Configure(key, val string) {
	p.configures[key] = val
}

func (p *Plotter) GetC(key string) string {
	return p.configures[key]
}

// Configure
type Configure struct {
	key               string
	val               string
	requiredCondition func(val string) bool
}

func NewConfigure(key, defaultVal string, requiredCondition func(val string) bool) *Configure {
	conf := new(Configure)
	conf.key = key
	conf.val = defaultVal
	conf.requiredCondition = requiredCondition
	return conf
}

var WITH_CONF = NewConfigure("with", "line", func(val string) bool {
	return val == "line" || val == "dots"
})

func (conf *Configure) SetVal(val string) {
	if conf.requiredCondition(val) {
		conf.val = val
	} else {
		panic(fmt.Sprintf("%v is illegal value of %v.", val, conf.key))
	}
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
	fun.plotter.configures = map[string]string{
		"_xMin": "-10.0",
		"_xMax": "10.0"}
	return fun
}

func (fun *Function2d) Configure(key, val string) {
	fun.plotter.Configure(key, val)
}

func (fun *Function2d) Configures(m map[string]string) {
	for key, val := range m {
		fun.plotter.Configure(key, val)
	}
}

func (fun *Function2d) UpdatePlotter(plotter *Plotter) {
	for key, val := range plotter.configures {
		fun.plotter.configures[key] = val
	}
}

func (fun *Function2d) GetData() [][2]float64 { // TODO: テスト書く
	xMin, _ := strconv.ParseFloat(fun.plotter.configures["_xMin"], 32)
	xMax, _ := strconv.ParseFloat(fun.plotter.configures["_xMax"], 32)
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
	for key, val := range fun.plotter.configures {
		if !strings.HasPrefix(key, "_") {
			s += fmt.Sprintf(" %v %v", key, val)
		}
	}
	return s
}

func (fun *Function2d) writeIntoGnufile(f os.File) {
	f.WriteString(fun.getGnuData())
}

const DefaultCurve2dSplitNum int = 100

type Curve2d struct {
	plotter  Plotter
	splitNum int
	c        func(float64) [2]float64
}

func NewCurve2d() *Curve2d {
	c := new(Curve2d)
	c.splitNum = DefaultCurve2dSplitNum
	c.plotter.configures = map[string]string{
		"_tMin": "-10.0",
		"_tMax": "10.0"}
	return c
}

func (c *Curve2d) Configure(key, val string) {
	c.plotter.Configure(key, val)
}

func (c *Curve2d) Configures(m map[string]string) {
	for key, val := range m {
		c.plotter.Configure(key, val)
	}
}

func (c *Curve2d) UpdatePlotter(plotter *Plotter) {
	for key, val := range plotter.configures {
		c.plotter.Configure(key, val)
	}
}

func (c *Curve2d) GetData() [][2]float64 { // TODO: test
	tMin, _ := strconv.ParseFloat(c.plotter.configures["_tMin"], 32)
	tMax, _ := strconv.ParseFloat(c.plotter.configures["_tMax"], 32)
	var sep = float64(tMax-tMin) / float64(c.splitNum-1)

	var a [][2]float64
	for j := 0; j < c.splitNum; j++ {
		var t float64 = tMin + float64(j)*sep
		cs := c.c(tMin + t*float64(j))
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
	for key, val := range c.plotter.configures {
		if !strings.HasPrefix(key, "_") {
			s += fmt.Sprintf(" %v %v", key, val)
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
	g.plotter.configures = map[string]string{}
	return g
}

func (g *Graph2d) Configure(key, val string) {
	g.plotter.Configure(key, val)
}

func (g *Graph2d) Configures(m map[string]string) {
	for key, val := range m {
		g.plotter.Configure(key, val)
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

func (g *Graph2d) UpdatePlotter(plotter *Plotter) {
	for key, val := range plotter.configures {
		g.plotter.Configure(key, val)
	}
}

func (g Graph2d) gnuplot(funcFilenames []string, curveFilenames []string) string {
	var s string

	for key, val := range g.plotter.configures {
		if !strings.HasPrefix(key, "_") {
			if val == "true" {
				s += fmt.Sprintf("set %v;\n", key)
			} else if val == "false" {
				s += fmt.Sprintf("set no%v;\n", key)
			} else {
				s += fmt.Sprintf("set %v %v;\n", key, val)
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
