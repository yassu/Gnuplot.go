package conf

import (
	"fmt"
	"regexp"
)

func inStr(elem string, array []string) bool {
	for _, a := range array {
		if elem == a {
			return true
		}
	}
	return false
}

func isNum(s string) bool {
	r := regexp.MustCompile(`^[+-]?[0-9]*[\.]?[0-9]+$`)
	return r.MatchString(s)
}

// Configures
type Configure struct {
	key               string
	vals              []string
	requiredCondition func(vals []string) bool
}

func NewConfigure(key string, defaultVals []string, requiredCondition func(vals []string) bool) *Configure {
	conf := new(Configure)
	conf.key = key
	conf.vals = defaultVals
	conf.requiredCondition = requiredCondition
	return conf
}

func (conf *Configure) SetVal(vals []string) {
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

// Function2d or Curve2d options
func WithConf() *Configure {
	return NewConfigure("with", []string{"lines"}, func(vals []string) bool {
		return len(vals) == 1 && inStr(vals[0], []string{
			"lines", "dots", "steps", "errorbars", "xerrorbar",
			"xyerrorlines", "points", "impulses", "fsteps", "errorlines", "xerrorlines",
			"yerrorlines", "surface", "vectors", "parallelaxes"})
	})
}

func GoXMinConf() *Configure {
	return NewConfigure("_xMin", []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func GoXMaxConf() *Configure {
	return NewConfigure("_xMax", []string{"10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func GoTMinConf() *Configure {
	return NewConfigure("_tMin", []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func GoTMaxConf() *Configure {
	return NewConfigure("_tMax", []string{"10.0"}, func(vals []string) bool {
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
	return NewConfigure("angles", []string{"radians"}, func(vals []string) bool {
		return len(vals) == 1 && inStr(vals[0], []string{"degrees", "radians", "true", "false"})
	})
}

func Graph2dConfs() []*Configure {
	return []*Configure{AnglesConf()}
}
