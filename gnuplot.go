package gnuplot

import (
  // "fmt"
)

type Plotter struct {
    Configures map[string] string
}

func (p *Plotter) init() {
    p.Configures = map[string] string{
      "xMin": "-10.0",
      "xMax": "10.0",
      "yMin": "-10.0",
      "yMax": "10.0"}
}

func (p *Plotter) configure(key, val string) {
    p.Configures[key] = val
}

func (p *Plotter) getC(key string) string {
    return p.Configures[key]
}
