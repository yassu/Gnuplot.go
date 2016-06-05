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

var POSITIONS = []string{"x", "y", "first", "second", "graph", "screen", "character"}

// Function2d or Curve2d options
func WithConf() *Configure {
	return NewConfigure([]string{"with", "w"}, []string{"lines"}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{
			"lines", "dots", "steps", "errorbars", "xerrorbar",
			"xyerrorlines", "points", "impulses", "fsteps", "errorlines", "xerrorlines",
			"yerrorlines", "surface", "vectors", "parallelaxes"})
	})
}

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
//     angles            arrow             autoscale         bars
//     bmargin           border            boxwidth          cbdata
//     cbdtics           cblabel           cbmtics           cbrange
//     cbtics            clabel            clip              cntrlabel
//     cntrparam         color             colorbox          colorsequence
//     contour           dashtype          data              datafile
//     date_specifiers   decimalsign       dgrid3d           dummy
//     encoding          fit               fontpath          format
//     function          grid              hidden3d          history
//     historysize       isosamples        key               label
//     linetype          link              lmargin           loadpath
//     locale            log               logscale          macros
//     mapping           margin            margins           missing
//     monochrome        mouse             multiplot         mx2tics
//     mxtics            my2tics           mytics            mztics
//     object            offsets           origin            output
//     palette           parametric        paxis             pm3d
//     pointintervalbox  pointsize         polar             print
//     psdir             raxis             rmargin           rrange
//     rtics             samples           size              style
//     surface           table             term              terminal
//     termoption        tics              ticscale          ticslevel
//     time_specifiers   timefmt           timestamp         title
//     tmargin           trange            urange            view
//     vrange            x2data            x2dtics           x2label
//     x2mtics           x2range           x2tics            x2zeroaxis
//     xdata             xdtics            xlabel            xmtics
//     xrange            xtics             xyplane           xzeroaxis
//     y2data            y2dtics           y2label           y2mtics
//     y2range           y2tics            y2zeroaxis        ydata
//     ydtics            ylabel            ymtics            yrange
func Graph2dAnglesConf() *Configure {
	return NewConfigure([]string{"angles"}, []string{"radians"}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{"degrees", "radians", "true", "false"})
	})
}

func Graph2dArrowConf() *Configure {
	return NewConfigure([]string{"arrow"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dAutoScaleConf() *Configure {
	return NewConfigure([]string{"autoscale"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dBarsConf() *Configure {
	return NewConfigure([]string{"bars"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dBmarginConf() *Configure {
	return NewConfigure([]string{"bmargin"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dBorderConf() *Configure {
	return NewConfigure([]string{"border"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dBoxwidthConf() *Configure {
	return NewConfigure([]string{"boxwidth"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dCbdataConf() *Configure {
	return NewConfigure([]string{"cbdata"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dCbdticsConf() *Configure {
	return NewConfigure([]string{"cbdtics"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dCblabelConf() *Configure {
	return NewConfigure([]string{"cblabel"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dCbmticsConf() *Configure {
	return NewConfigure([]string{"cbmtics"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dCbrangeConf() *Configure {
	return NewConfigure([]string{"cbrange"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dCbticsConf() *Configure {
	return NewConfigure([]string{"cbtics"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dClabelConf() *Configure {
	return NewConfigure([]string{"clabel"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dClipConf() *Configure {
	return NewConfigure([]string{"clip"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dCntrlabelConf() *Configure {
	return NewConfigure([]string{"cntrlabel"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dCntrparamConf() *Configure {
	return NewConfigure([]string{"cntrparam"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dColorConf() *Configure {
	return NewConfigure([]string{"color"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dColorboxConf() *Configure {
	return NewConfigure([]string{"colorbox"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dColorsequenceConf() *Configure {
	return NewConfigure([]string{"colorsequence"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dContourConf() *Configure {
	return NewConfigure([]string{"contour"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dDashtypeConf() *Configure {
	return NewConfigure([]string{"dashtype"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dDataConf() *Configure {
	return NewConfigure([]string{"data"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dDatafileConf() *Configure {
	return NewConfigure([]string{"datafile"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dDateSpecifiersConf() *Configure {
	return NewConfigure([]string{"date_specifiers"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dDecimalsignConf() *Configure {
	return NewConfigure([]string{"decimalsign"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dDgrid3dConf() *Configure {
	return NewConfigure([]string{"dgrid3d"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dDummyConf() *Configure {
	return NewConfigure([]string{"dummy"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dEncodingConf() *Configure {
	return NewConfigure([]string{"encoding"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dFitConf() *Configure {
	return NewConfigure([]string{"fit"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dFontPathConf() *Configure {
	return NewConfigure([]string{"fontpath"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dFormatConf() *Configure {
	return NewConfigure([]string{"format"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dFunctionConf() *Configure {
	return NewConfigure([]string{"function"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dGridConf() *Configure {
	return NewConfigure([]string{"grid"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dHidden3dConf() *Configure {
	return NewConfigure([]string{"hidden3d"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dHistoryConf() *Configure {
	return NewConfigure([]string{"history"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dHistorysizeConf() *Configure {
	return NewConfigure([]string{"historysize"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dIsosamplesConf() *Configure {
	return NewConfigure([]string{"isosamples"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dKeyConf() *Configure {
	return NewConfigure([]string{"key"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dLabelConf() *Configure {
	return NewConfigure([]string{"label"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dLinetypeConf() *Configure {
	return NewConfigure([]string{"label"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dLinkConf() *Configure {
	return NewConfigure([]string{"link"}, []string{""}, func(vals []string) bool {
		return true
	})
}

func Graph2dConfs() []*Configure {
	return []*Configure{AnglesConf()}
}
