package gnuplot

import (
	"fmt"
	"github.com/yassu/gnup/conf"
	"github.com/yassu/gnup/utils"
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
			p.configures[j].SetVals(conf.GetVals())
			return
		}
	}
	p.configures = append(p.configures, conf)
}

func (p *Plotter) GetC(key string) []string {
	for j := range p.configures {
		if p.configures[j].GetKey() == key {
			return p.configures[j].GetVals()
		}
	}
	return []string{}
}

type PlotElement2d interface {
	GetData() [][2]float64
	getGnuData() string
	gnuplot(filename string) string
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

func (fun *Function2d) Configure(key string, vals []string) {
	for j, conf := range fun.plotter.configures {
		if utils.InStr(key, conf.AliasedKeys()) {
			fun.plotter.configures[j].SetVals(vals)
			return
		}
	}
	panic(fmt.Sprintf("%v is not a key.", key))
}

func (fun *Function2d) Configures(sconf map[string][]string) {
	for key, vals := range sconf {
		fun.Configure(key, vals)
	}
}

func (fun Function2d) GetData() [][2]float64 { // TODO: テスト書く
	xMin, _ := strconv.ParseFloat(fun.plotter.GetC("_xMin")[0], 32)
	xMax, _ := strconv.ParseFloat(fun.plotter.GetC("_xMax")[0], 32)
	var sep = float64(xMax-xMin) / float64(fun.splitNum-1)

	var a [][2]float64
	for j := 0; j < fun.splitNum; j++ {
		t := xMin + float64(j)*sep
		y := fun.f(t)
		a = append(a, [2]float64{t, y})
	}
	return a
}

func (fun Function2d) getGnuData() string {
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
	title := fun.plotter.GetC("_title")
	var s = fmt.Sprintf("\"%v\"", filename)
	if !isDummyVal(title) {
		s += fmt.Sprintf(" title \"%v\"", title[0])
	}

	for _, conf := range fun.plotter.configures {
		if !strings.HasPrefix(conf.GetKey(), "_") && !isDummyVal(conf.GetVals()) {
			vals := conf.GetVals()
			s += fmt.Sprintf(" %v ", conf.GetKey())
			if vals[len(vals)-1] == "true" {
				vals = vals[:len(vals)-1]
			} else if vals[len(vals)-1] == "false" {
				vals = vals[:len(vals)-1]
				s += "no"
			}
			for _, val := range vals {
				s += fmt.Sprintf(" %v", val)
			}
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
	return c
}

func (c *Curve2d) setConfigure() {
	for _, conf := range conf.Curve2dConfs() {
		c.plotter.Configure(conf)
	}
}

func (c *Curve2d) Configure(key string, vals []string) {
	for j, conf := range c.plotter.configures {
		if utils.InStr(key, conf.AliasedKeys()) {
			c.plotter.configures[j].SetVals(vals)
			return
		}
	}
	panic(fmt.Sprintf("%v is not a key.", key))
}

func (c *Curve2d) Configures(sconf map[string][]string) {
	for key, vals := range sconf {
		c.Configure(key, vals)
	}
}

func (c Curve2d) GetData() [][2]float64 { // TODO: test
	tMin, _ := strconv.ParseFloat(c.plotter.GetC("_tMin")[0], 32)
	tMax, _ := strconv.ParseFloat(c.plotter.GetC("_tMax")[0], 32)
	var sep = float64(tMax-tMin) / float64(c.splitNum-1)

	var a [][2]float64
	for j := 0; j < c.splitNum; j++ {
		cs := c.c(tMin + float64(j)*sep)
		a = append(a, [2]float64{cs[0], cs[1]})
	}
	return a
}

func (c Curve2d) getGnuData() string {
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
		if !strings.HasPrefix(conf.GetKey(), "_") && !isDummyVal(conf.GetVals()) {
			vals := conf.GetVals()
			s += fmt.Sprintf(" %v ", conf.GetKey())
			if vals[len(vals)-1] == "true" {
				vals = vals[:len(vals)-1]
			} else if vals[len(vals)-1] == "false" {
				vals = vals[:len(vals)-1]
				s += "no"
			}
			for _, val := range vals {
				s += fmt.Sprintf(" %v", val)
			}
		}
	}
	return s
}

func isDummyVal(vals []string) bool {
	return len(vals) == 0
}

// Graph
type Graph2d struct {
	plotter Plotter
	pElems  []PlotElement2d
}

func NewGraph2d() *Graph2d {
	g := new(Graph2d)
	g.setConfigure()
	return g
}

func (g *Graph2d) setConfigure() {
	for _, conf := range conf.GraphConfs() {
		g.plotter.Configure(conf)
	}
}

func (g *Graph2d) Configure(key string, vals []string) {
	for j, conf := range g.plotter.configures {
		if conf.GetKey() == key {
			g.plotter.configures[j].SetVals(vals)
			return
		}
	}
	panic(fmt.Sprintf("%v is not a key.", key))
}

func (g *Graph2d) Configures(sconf map[string][]string) {
	for key, vals := range sconf {
		g.Configure(key, vals)
	}
}

func (g *Graph2d) AppendPElem(p PlotElement2d) {
	g.pElems = append(g.pElems, p)
}

func (g Graph2d) writeIntoFile(data string, f *os.File) {
	f.WriteString(data)
}

func (g Graph2d) gnuplot(elemFilenames []string) string {
	var s string

	for _, conf := range g.plotter.configures {
		if !strings.HasPrefix(conf.GetKey(), "_") && !isDummyVal(conf.GetVals()) {
			vals := conf.GetVals()
			s += "set "
			if vals[len(vals)-1] == "true" {
				vals = vals[:len(vals)-1]
			} else if vals[len(vals)-1] == "false" {
				vals = vals[:len(vals)-1]
				s += "no"
			}
			s += conf.GetKey()
			for _, val := range vals {
				s += fmt.Sprintf(" %v ", val)
			}
			s += ";\n"
		}
	}

	s += "plot "
	for j, _ := range g.pElems {
		s += g.pElems[j].gnuplot(elemFilenames[j])
		if j != len(g.pElems)-1 {
			s += ", "
		}
	}
	s += ";\n"

	s += "pause -1;\n"
	return s
}

func (g *Graph2d) Run() {
	tmpDir := os.TempDir() + "/gnup/"
	// TODO: tmpDirがなければ作る
	// execFilename := tmpDir + "exec.gnu"
	execFilename := "exec.gnu"

	// それぞれのcurveのdataをtempファイルに書き込む
	// また, それらのファイルの名前を curve_filenames []stringに格納する
	var plotElemFilenames []string
	for _, p := range g.pElems {
		file, _ := ioutil.TempFile(tmpDir, "")
		defer func() {
			file.Close()
		}()
		g.writeIntoFile(p.getGnuData(), file)
		plotElemFilenames = append(plotElemFilenames, file.Name())
	}

	// 実行するgnuplotの実行ファイルをtempファイルに書き込む
	os.Remove(execFilename)
	execFile, err := os.OpenFile(execFilename, os.O_CREATE|os.O_WRONLY, 0666)
	defer func() {
		execFile.Close()
	}()
	if err != nil {
		fmt.Println(err)
	} else {
		execFile.WriteString(g.gnuplot(plotElemFilenames))
	}
}
