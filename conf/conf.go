package conf

import (
	"fmt"
	"github.com/yassu/gnuplot.go/utils"
	"regexp"
	"strconv"
)

func isNum(s string) bool {
	r := regexp.MustCompile(`^[+-]?[0-9]*[\.]?[0-9]+$`)
	return r.MatchString(s)
}

func isIntStr(s string) bool {
	r := regexp.MustCompile(`^[+-]?[0-9]+$`)
	return r.MatchString(s)
}

func isSixHex(s string) bool {
	r := regexp.MustCompile(`^[0-9a-f]{6}$`)
	return r.MatchString(s)
}

func isEightHex(s string) bool {
	r := regexp.MustCompile(`^[0-9a-f]{8}$`)
	return r.MatchString(s)
}

// either s is float of 0 ~ 1 or not
func isSmallFloat(s string) bool {
	if !isNum(s) {
		return false
	}
	f, _ := strconv.ParseFloat(s, 32)
	return 0 <= f && f <= 1
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

var COLOR_NAMES = []string{
	"white", "black", "dark-grey", "red", "web-green", "web-blue",
	"dark-magenta", "dark-cyan", "dark-orange", "dark-yellow", "royalblue",
	"goldenrod", "dark-spring-green", "purple", "steelblue", "dark-red",
	"dark-chartreuse", "orchild", "aquamarine", "brown", "yellow",
	"turquoise",
	"grey", "grey0", "grey10", "grey20", "grey30", "grey40", "grey50",
	"grey60", "grey70", "grey80", "grey90", "grey100",
	"light-red", "light-green", "light-blue", "light-magenta", "light-cyan",
	"light-goldenrod", "light-pink", "light-turquoise", "gold", "green",
	"dark-green", "sprint-green", "forest-green", "sea-green", "blue",
	"dark-blue", "midnight-blue", "navy", "medium-blue", "skyblue",
	"cyan", "magenta", "dark-turquoise", "dark-pink", "coral", "light-coral",
	"orange-red", "salmon", "dark-salmon", "khaki", "dark-khaki",
	"dark-goldenrod", "beige", "olive", "orange", "violet", "dark-violet",
	"plum", "dark-plum", "dark-olivegreen", "orangered4", "brown4", "sienna4",
	"orchid4", "mediumpurple3", "slateblue1", "yellow4", "sienna1", "tan1",
	"standybrown", "light-salmon", "pink", "khaki1", "lemonchiffon", "bisque",
	"honeydew", "slategrey", "seagreen", "antiquewhite", "chartreuse",
	"greenyellow", "gray", "light-gray", "light-grey",
	"dark-gray", "slategray",
	"gray0", "gray10", "gray20", "gray30", "gray40", "gray50", "gray60",
	"gray70", "gray80", "gray90", "gray100"}

func LineColorConf() *Configure {
	return NewConfigure([]string{"linecolor", "lc"}, []string{"1"}, func(vals []string) bool {
		if len(vals) == 0 {
			return false
		}
		// in case of linecolor "colorname"
		val := vals[0]
		if utils.InStr(val, COLOR_NAMES) {
			return true
		}

		// in case of linecolor <n>
		if isIntStr(val) {
			return true
		}

		// in case of linecolor <colorspec> and len(vals) == 1
		if utils.InStr(val, []string{"variable", "bgnd", "black"}) {
			return true
		}

		if len(vals) == 2 && vals[0] == "rgbcolor" && utils.InStr(vals[1], COLOR_NAMES) {
			return true
		}
		if len(vals) == 2 && vals[0] == "rgbcolor" && vals[1][0:2] == "0x" && isSixHex(vals[1][2:]) {
			return true
		}
		if len(vals) == 2 && vals[0] == "rgbcolor" && vals[1][0:2] == "0x" && isEightHex(vals[1][2:]) {
			return true
		}
		if len(vals) == 2 && vals[0] == "rgbcolor" && vals[1][0] == '#' && isSixHex(vals[1][1:]) {
			return true
		}
		if len(vals) == 2 && vals[0] == "rgbcolor" && vals[1][0] == '#' && isEightHex(vals[1][1:]) {
			return true
		}
		if len(vals) == 2 && vals[0] == "rgbcolor" && isIntStr(vals[1]) {
			return true
		}
		return false
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
