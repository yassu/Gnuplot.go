package conf

import (
	"fmt"
	"regexp"
)

func InStr(elem string, array []string) bool {
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

func (conf *Configure) SetVal(val string) {
	if conf.requiredCondition(val) {
		conf.val = val
	} else {
		panic(fmt.Sprintf("%v is illegal value of %v.", val, conf.key))
	}
}

func WithConf() *Configure {
	return NewConfigure("with", "lines", func(val string) bool {
		return InStr(val, []string{
			"lines", "dots", "steps", "errorbars", "xerrorbar",
			"xyerrorlines", "points", "impulses", "fsteps", "errorlines", "xerrorlines",
			"yerrorlines", "surface", "vectors", "parallelaxes"})
	})
}

// for gnuplot.go options
func GoXMinConf() *Configure {
	return NewConfigure("_xMin", "-10.0", isNum)
}

func GoXMaxConf() *Configure {
	return NewConfigure("_xMax", "10.0", isNum)
}
