package conf

import (
	"fmt"
	"github.com/yassu/gnuplot.go/utils"
	"regexp"
)

func isNum(s string) bool {
	r := regexp.MustCompile(`^[+-]?[0-9]*[\.]?[0-9]+$`)
	return r.MatchString(s)
}

// Configures
type Configure struct {
	key               string
	aliasKeys         []string
	vals              []string
	requiredCondition func(vals []string) bool
}

func NewConfigure(keys []string, defaultVals []string, requiredCondition func(vals []string) bool) *Configure {
	conf := new(Configure)
	conf.key = keys[0]
	conf.aliasKeys = keys
	conf.vals = defaultVals
	conf.requiredCondition = requiredCondition
	return conf
}

func (conf *Configure) SetVals(vals []string) {
	if conf.requiredCondition(vals) {
		conf.vals = vals
	} else {
		panic(fmt.Sprintf("%v is illegal values of %v.", vals, conf.key))
	}
}

func (conf *Configure) GetKey() string {
	return conf.key
}

func (conf *Configure) GetVals() []string {
	return conf.vals
}

func (conf *Configure) AliasedKeys() []string {
	return conf.aliasKeys
}

// Function2d or Curve2d options
func WithConf() *Configure {
	return NewConfigure([]string{"with", "w"}, []string{"lines"}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{
			"lines", "dots", "steps", "errorbars", "xerrorbar",
			"xyerrorlines", "points", "impulses", "fsteps", "errorlines", "xerrorlines",
			"yerrorlines", "surface", "vectors", "parallelaxes"})
	})
}

func GoXMinConf() *Configure {
	return NewConfigure([]string{"_xMin"}, []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func GoXMaxConf() *Configure {
	return NewConfigure([]string{"_xMax"}, []string{"10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func GoTMinConf() *Configure {
	return NewConfigure([]string{"_tMin"}, []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func GoTMaxConf() *Configure {
	return NewConfigure([]string{"_tMax"}, []string{"10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func Function2dConfs() []*Configure {
	return []*Configure{WithConf(), GoXMinConf(), GoXMaxConf()}
}

func Curve2dConfs() []*Configure {
	return []*Configure{WithConf(), GoTMinConf(), GoTMaxConf()}
}

// Graph options
func AnglesConf() *Configure {
	return NewConfigure([]string{"angles"}, []string{"radians"}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{"degrees", "radians", "true", "false"})
	})
}

func Graph2dConfs() []*Configure {
	return []*Configure{AnglesConf()}
}
