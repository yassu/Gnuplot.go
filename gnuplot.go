package gnuplot

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	// "os/exec"
)

type Plotter struct {
	configures map[string]string
}

func (p *Plotter) Init() {
	p.configures = map[string]string{}
}

func (p *Plotter) Configure(key, val string) {
	p.configures[key] = val
}

func (p *Plotter) GetC(key string) string {
	return p.configures[key]
}

var DefaultFunction2dSplitNum int = 1000

type Function2d struct {
	plotter  Plotter
	splitNum int
	f        func(float64) float64
}

func (fun *Function2d) Init() {
	fun.splitNum = DefaultFunction2dSplitNum
	fun.plotter.configures = map[string]string{
		"_xMin": "-10.0",
		"_xMax": "10.0"}
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
	var s = fmt.Sprintf("plot \"%v\"", filename)
	for key, val := range fun.plotter.configures {
		if !strings.HasPrefix(key, "_") {
			s += fmt.Sprintf(" %v %v", key, val)
		}
	}
	return s + ";\n"
}

func (fun *Function2d) writeIntoGnufile(f os.File) {
	f.WriteString(fun.getGnuData())
}

var DefaultCurve2dSplitNum int = 100

type Curve2d struct {
	plotter  Plotter
	splitNum int
	c        func(float64) [2]float64
}

func (c *Curve2d) Init() {
	c.splitNum = DefaultCurve2dSplitNum
	c.plotter.configures = map[string]string{
		"_tMin": "-10.0",
		"_tMax": "10.0"}
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

func (c Curve2d) gnuplot(fileName string) string {
	return fmt.Sprintf("plot %v\n;", fileName)
}

// Graph
type Graph2d struct {
	plotter   Plotter
	functions []Function2d
	curves    []Curve2d
}

func (g *Graph2d) AppendFunc(f Function2d) {
	fmt.Println("before of AppendFunc")
	g.functions = append(g.functions, f)
	fmt.Println("after of AppendFunc")
}

func (g Graph2d) writeIntoFile(data string, f *os.File) {
	f.WriteString(data)
}

func (g Graph2d) exec_gnuplot() {
	// until
}

func (g Graph2d) gnuplot(funcFilenames []string, curveFilenames []string) string {
	var s string
	for j, _ := range g.functions {
		s += g.functions[j].gnuplot(funcFilenames[j])
	}
	s += "\n"
	for j, _ := range g.curves {
		s += g.curves[j].gnuplot(curveFilenames[j])
	}
	s += "pause -1;"
	s += "\n"
	return s
}

func (g *Graph2d) Run() {
	tmpDir := os.TempDir()

	// それぞれのfunctionのdataをtempファイルに書き込む
	// また, それらのファイルの名前を func_filenames []string に格納する
	var funcFilenames []string
	for _, fun := range g.functions {
		file, _ := ioutil.TempFile(tmpDir, "")
		defer func() {
			file.Close()
		}()
		g.writeIntoFile(fun.getGnuData(), file)
		funcFilenames = append(funcFilenames, file.Name())
	}

	// それぞれのcurveのdataをtempファイルに書き込む
	// また, それらのファイルの名前を curve_filenames []stringに格納する

	// 実行するgnuplotの実行ファイルをtempファイルに書き込む
	execFile, _ := os.OpenFile("exec_gnu.gnu", os.O_CREATE|os.O_WRONLY, 0666)
	defer func() {
		execFile.Close()
	}()
	fmt.Println(funcFilenames)
	execFile.WriteString(g.gnuplot(funcFilenames, []string{}))
}
