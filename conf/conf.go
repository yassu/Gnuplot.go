package conf

import (
	"fmt"
	"github.com/yassu/gnup/utils"
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

func isPosNum(s string) bool {
	r := regexp.MustCompile(`^[+]?[0-9]+[\.]?[0-9]+$`)
	return r.MatchString(s)
}

func isNaturalStr(s string) bool {
	r := regexp.MustCompile(`^[+]?[0-9]+$`)
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

func Function2dConfs() []*Configure {
	return []*Configure{
		PStyleTitleConf(),
		PStyleWithConf(),
		PStyleLineColorConf(),
		PStyleLineTypeConf(),
		PStyleLineWidthConf(),
		PStylePointTypeConf(),
		PStylePointSizeConf(),
		PStyleFillConf(),
		PStyleNoHidden3dConf(),
		PStyleNoContoursConf(),
		PStyleNoSurfaceConf(),
		PStylePaletteConf(),
		PStyleSplitNumConf(),
		PStyleGoXMinConf(),
		PStyleGoXMaxConf()}
}

func Function3dConfs() []*Configure {
	return []*Configure{
		PStyleTitleConf(),
		PStyleWithConf(),
		PStyleLineColorConf(),
		PStyleLineTypeConf(),
		PStyleLineWidthConf(),
		PStylePointTypeConf(),
		PStylePointSizeConf(),
		PStyleFillConf(),
		PStyleNoHidden3dConf(),
		PStyleNoContoursConf(),
		PStyleNoSurfaceConf(),
		PStylePaletteConf(),
		PStyleSplitNumConf(),
		PStyleGoXMinConf(),
		PStyleGoXMaxConf(),
		PStyleGoYMinConf(),
		PStyleGoYMaxConf()}
}

func Curve2dConfs() []*Configure {
	return []*Configure{
		PStyleWithConf(),
		PStyleLineColorConf(),
		PStylePointTypeConf(),
		PStyleLineTypeConf(),
		PStyleLineWidthConf(),
		PStylePointSizeConf(),
		PStyleFillConf(),
		PStyleNoHidden3dConf(),
		PStyleNoContoursConf(),
		PStyleNoSurfaceConf(),
		PStylePaletteConf(),
		PStyleSplitNumConf(),
		PStyleGoTMinConf(),
		PStyleGoTMaxConf()}
}

func Curve3dConfs() []*Configure {
	return []*Configure{
		PStyleWithConf(),
		PStyleLineColorConf(),
		PStylePointTypeConf(),
		PStyleLineTypeConf(),
		PStyleLineWidthConf(),
		PStylePointSizeConf(),
		PStyleFillConf(),
		PStyleNoHidden3dConf(),
		PStyleNoContoursConf(),
		PStyleNoSurfaceConf(),
		PStylePaletteConf(),
		PStyleSplitNumConf(),
		PStyleGoTMinConf(),
		PStyleGoTMaxConf()}
}

func Surface3dConfs() []*Configure {
	return []*Configure{
		PStyleWithConf(),
		PStyleLineColorConf(),
		PStylePointTypeConf(),
		PStyleLineTypeConf(),
		PStyleLineWidthConf(),
		PStylePointSizeConf(),
		PStyleFillConf(),
		PStyleNoHidden3dConf(),
		PStyleNoContoursConf(),
		PStyleNoSurfaceConf(),
		PStylePaletteConf(),
		PStyleSplitNumConf(),
		PStyleGoUMinConf(),
		PStyleGoUMaxConf(),
		PStyleGoVMinConf(),
		PStyleGoVMaxConf()}
}

func GraphConfs() []*Configure {
	return []*Configure{
		GraphAnglesConf(),
		GraphAutoScaleConf(),
		GraphBarsConf(),
		GraphBmarginConf(),
		GraphBorderConf(),
		GraphBoxwidthConf(),
		GraphCbdataConf(),
		GraphCbdticsConf(),
		GraphCblabelConf(),
		GraphCbmticsConf(),
		GraphCbrangeConf(),
		GraphCbticsConf(),
		GraphClabelConf(),
		GraphClipConf(),
		GraphCntrlabelConf(),
		GraphCntrparamConf(),
		GraphColorboxConf(),
		GraphColorsequenceConf(),
		GraphContourConf(),
		GraphDashtypeConf(),
		GraphDataConf(),
		GraphDatafileConf(),
		GraphDateSpecifiersConf(),
		GraphDecimalsignConf(),
		GraphDgrid3dConf(),
		GraphDummyConf(),
		GraphEncodingConf(),
		GraphFitConf(),
		GraphFontPathConf(),
		GraphFormatConf(),
		GraphFunctionConf(),
		GraphGridConf(),
		GraphHidden3dConf(),
		GraphHistoryConf(),
		GraphHistorysizeConf(),
		GraphIsosamplesConf(),
		GraphKeyConf(),
		GraphLabelConf(),
		GraphLinetypeConf(),
		GraphLinkConf(),
		GraphLmarginConf(),
		GraphLoadpathConf(),
		GraphLocaleConf(),
		GraphLogConf(),
		GraphLogscaleConf(),
		GraphMacrosConf(),
		GraphMappingConf(),
		GraphMarginConf(),
		GraphMarginsConf(),
		GraphMissingConf(),
		GraphMonochromeConf(),
		GraphMouseConf(),
		GraphMultiplotConf(),
		GraphMx2ticsConf(),
		GraphMxticsConf(),
		GraphMy2ticsConf(),
		GraphMyticsConf(),
		GraphMzticsConf(),
		GraphObjectConf(),
		GraphOffsetsConf(),
		GraphOriginConf(),
		GraphOutputConf(),
		GraphPaletteConf(),
		GraphParametricConf(),
		GraphPaxisConf(),
		GraphPm3dConf(),
		GraphPointintervalboxConf(),
		GraphPointsizeConf(),
		GraphPolarConf(),
		GraphPolarConf(),
		GraphPsdirConf(),
		GraphRaxisConf(),
		GraphRmarginConf(),
		GraphRrangeConf(),
		GraphRticsConf(),
		GraphSamplesConf(),
		GraphSizeConf(),
		GraphStyleConf(),
		GraphSurfaceConf(),
		GraphTableConf(),
		GraphTermConf(),
		GraphTerminalConf(),
		GraphTermoptionConf(),
		GraphTicsConf(),
		GraphTicscaleConf(),
		GraphTicslevelcaleConf(),
		GraphTimeSpecifiersConf(),
		GraphTimefmtConf(),
		GraphTimestampConf(),
		GraphTitleConf(),
		GraphTmarginConf(),
		GraphTrangeConf(),
		GraphUrangeConf(),
		GraphViewConf(),
		GraphVrangeConf(),
		GraphX2dataConf(),
		GraphX2dticsConf(),
		GraphX2labelConf(),
		GraphX2mticsConf(),
		GraphX2rangeConf(),
		GraphX2ticsConf(),
		GraphX2zeroaxisConf(),
		GraphXdataConf(),
		GraphXdticsConf(),
		GraphXlabelConf(),
		GraphXmticsConf(),
		GraphXrangeConf(),
		GraphXticsConf(),
		GraphXyplaneConf(),
		GraphXzeroaxisConf(),
		GraphY2dataConf(),
		GraphY2dticsConf(),
		GraphY2dlabelConf(),
		GraphY2mticsConf(),
		GraphY2rangeConf(),
		GraphY2ticsConf(),
		GraphY2zeroaxisConf(),
		GraphYdataConf(),
		GraphYdticsConf(),
		GraphYlabelConf(),
		GraphYmticsConf(),
		GraphYrangeConf()}
}

// Function2d or Curve2d options
func PStyleTitleConf() *Configure {
	return NewConfigure([]string{"_title"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 &&
			len(vals[0]) >= 1
	})
}

func PStyleWithConf() *Configure {
	return NewConfigure([]string{"with", "w"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{
			"lines", "dots", "steps", "errorbars", "xerrorbar",
			"xyerrorlines", "points", "impulses", "fsteps", "errorlines", "xerrorlines",
			"yerrorlines", "surface", "vectors", "parallelaxes"})
	})
}

func PStyleLineColorConf() *Configure {
	return NewConfigure([]string{"linecolor", "lc"}, []string{}, func(vals []string) bool {
		if len(vals) == 0 {
			return false
		}
		// in case of linecolor "colorname"
		if len(vals) == 1 && utils.InStr(vals[0], COLOR_NAMES) {
			return true
		}

		// in case of linecolor <n>
		if len(vals) == 1 && isIntStr(vals[0]) {
			return true
		}

		// in case of linecolor <colorspec> and len(vals) == 1
		if len(vals) == 1 && utils.InStr(vals[0], []string{"variable", "bgnd", "black"}) {
			return true
		}

		if len(vals) == 2 && vals[0] == "rgbcolor" && utils.InStr(vals[1], COLOR_NAMES) {
			return true
		}
		if len(vals) == 2 && len(vals[1]) >= 2 &&
			vals[0] == "rgbcolor" && vals[1][0:2] == "0x" && isSixHex(vals[1][2:]) {
			return true
		}
		if len(vals) == 2 && len(vals[1]) >= 2 &&
			vals[0] == "rgbcolor" && vals[1][0:2] == "0x" && isEightHex(vals[1][2:]) {
			return true
		}
		if len(vals) == 2 && len(vals[1]) >= 1 &&
			vals[0] == "rgbcolor" && vals[1][0] == '#' && isSixHex(vals[1][1:]) {
			return true
		}
		if len(vals) == 2 && len(vals[1]) >= 1 &&
			vals[0] == "rgbcolor" && vals[1][0] == '#' && isEightHex(vals[1][1:]) {
			return true
		}
		if len(vals) == 2 && vals[0] == "rgbcolor" && isIntStr(vals[1]) {
			return true
		}
		if len(vals) == 2 && vals[0] == "rgbcolor" && vals[1] == "variable" {
			return true
		}
		if len(vals) == 3 && vals[0] == "palette" && vals[1] == "frac" &&
			isSmallFloat(vals[2]) {
			return true
		}
		if len(vals) == 3 && vals[0] == "palette" && vals[1] == "cb" &&
			isNum(vals[2]) {
			return true
		}
		if len(vals) == 2 && vals[0] == "palette" && vals[1] == "z" {
			return true
		}
		return false
	})
}

func PStylePointTypeConf() *Configure {
	return NewConfigure([]string{"pointtype", "pt"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && isIntStr(vals[0])
	})
}

func PStyleLineTypeConf() *Configure {
	return NewConfigure([]string{"linetype", "lt"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleLineWidthConf() *Configure {
	return NewConfigure([]string{"linewidth", "lw"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStylePointSizeConf() *Configure {
	return NewConfigure([]string{"pointsize", "ps"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleFillConf() *Configure {
	return NewConfigure([]string{"fill", "fs"}, []string{}, func(vals []string) bool {
		return true
	})
}

func PStyleNoHidden3dConf() *Configure {
	return NewConfigure([]string{"nohidden3d"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{"true"})
	})
}

func PStyleNoContoursConf() *Configure {
	return NewConfigure([]string{"nocontours"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && vals[0] == "true"
	})
}

func PStyleNoTitleConf() *Configure {
	return NewConfigure([]string{"notitle"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && vals[0] == "true"
	})
}

func PStyleNoSurfaceConf() *Configure {
	return NewConfigure([]string{"nosurface"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && vals[0] == "true"
	})
}

func PStylePaletteConf() *Configure {
	return NewConfigure([]string{"palette"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && vals[0] == "true"
	})
}

func PStyleSplitNumConf() *Configure {
	return NewConfigure([]string{"_splitNum"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && isNaturalStr(vals[0])
	})
}

func PStyleGoXMinConf() *Configure {
	return NewConfigure([]string{"_xMin"}, []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoXMaxConf() *Configure {
	return NewConfigure([]string{"_xMax"}, []string{"10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoYMinConf() *Configure {
	return NewConfigure([]string{"_yMin"}, []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoYMaxConf() *Configure {
	return NewConfigure([]string{"_yMax"}, []string{"10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoTMinConf() *Configure {
	return NewConfigure([]string{"_tMin"}, []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoTMaxConf() *Configure {
	return NewConfigure([]string{"_tMax"}, []string{"10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoUMinConf() *Configure {
	return NewConfigure([]string{"_uMin"}, []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoUMaxConf() *Configure {
	return NewConfigure([]string{"_uMax"}, []string{"10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoVMinConf() *Configure {
	return NewConfigure([]string{"_vMin"}, []string{"-10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
}

func PStyleGoVMaxConf() *Configure {
	return NewConfigure([]string{"_vMax"}, []string{"10.0"}, func(vals []string) bool {
		return len(vals) == 1 && isNum(vals[0])
	})
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
func GraphAnglesConf() *Configure {
	return NewConfigure([]string{"angles"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{"degrees", "radians", "true", "false"})
	})
}

func GraphAutoScaleConf() *Configure {
	return NewConfigure([]string{"autoscale"}, []string{}, func(vals []string) bool {
		if len(vals) != 1 {
			return false
		}
		val := vals[0]

		axes := []string{"x", "y", "z", "cb", "x2", "y2", "xy"}
		for _, axe := range axes {
			if axe+"min" == val {
				return true
			}
			if axe+"max" == val {
				return true
			}
			if axe+"fixmin" == val {
				return true
			}
			if axe+"fixmax" == val {
				return true
			}
			if axe+"fix" == val {
				return true
			}
		}

		if utils.InStr(vals[0], []string{"fix", "keepfix"}) {
			return true
		}
		if vals[0] == "noextend" {
			return true
		}
		return false
	})
}

func GraphBarsConf() *Configure {
	return NewConfigure([]string{"bars"}, []string{}, func(vals []string) bool {
		if len(vals) == 1 {
			val := vals[0]
			return utils.InStr(val, []string{"small", "large", "fullwidth", "front", "back"}) ||
				isPosNum(val)
		} else if len(vals) == 2 {
			return (utils.InStr(vals[0], []string{"small", "large", "fullwidth"}) || isPosNum(vals[0])) &&
				utils.InStr(vals[1], []string{"front", "back"})
		} else {
			return false
		}
	})
}

func GraphBmarginConf() *Configure {
	return NewConfigure([]string{"bmargin"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && isPosNum(vals[0])
	})
}

func GraphBorderConf() *Configure {
	return NewConfigure([]string{"border"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphBoxwidthConf() *Configure {
	return NewConfigure([]string{"boxwidth"}, []string{}, func(vals []string) bool {
		if len(vals) == 1 {
			return utils.InStr(vals[0], []string{"absolute", "relative"}) ||
				isPosNum(vals[0])
		} else if len(vals) == 2 {
			return isPosNum(vals[0]) && utils.InStr(vals[1], []string{"absolute", "relative"})
		} else {
			return false
		}
	})
}

func GraphCbdataConf() *Configure {
	return NewConfigure([]string{"cbdata"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphCbdticsConf() *Configure {
	return NewConfigure([]string{"cbdtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphCblabelConf() *Configure {
	return NewConfigure([]string{"cblabel"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphCbmticsConf() *Configure {
	return NewConfigure([]string{"cbmtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphCbrangeConf() *Configure {
	return NewConfigure([]string{"cbrange"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphCbticsConf() *Configure {
	return NewConfigure([]string{"cbtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphClabelConf() *Configure {
	return NewConfigure([]string{"clabel"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphClipConf() *Configure {
	return NewConfigure([]string{"clip"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{"points", "one", "two"})
	})
}

func GraphCntrlabelConf() *Configure {
	return NewConfigure([]string{"cntrlabel"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphCntrparamConf() *Configure {
	return NewConfigure([]string{"cntrparam"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphColorboxConf() *Configure {
	return NewConfigure([]string{"colorbox"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphColorsequenceConf() *Configure {
	return NewConfigure([]string{"colorsequence", "colors"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 && utils.InStr(vals[0], []string{"default", "classic", "podo"})
	})
}

func GraphContourConf() *Configure {
	return NewConfigure([]string{"contour"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 &&
			utils.InStr(vals[0], []string{"true", "base", "surface", "both"})
	})
}

func GraphDashtypeConf() *Configure {
	return NewConfigure([]string{"dashtype"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphDataConf() *Configure {
	return NewConfigure([]string{"data"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphDatafileConf() *Configure {
	return NewConfigure([]string{"datafile"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphDateSpecifiersConf() *Configure {
	return NewConfigure([]string{"date_specifiers"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphDecimalsignConf() *Configure {
	return NewConfigure([]string{"decimalsign"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphDgrid3dConf() *Configure {
	return NewConfigure([]string{"dgrid3d"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphDummyConf() *Configure {
	return NewConfigure([]string{"dummy"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphEncodingConf() *Configure {
	return NewConfigure([]string{"encoding"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphFitConf() *Configure {
	return NewConfigure([]string{"fit"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphFontPathConf() *Configure {
	return NewConfigure([]string{"fontpath"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphFormatConf() *Configure {
	return NewConfigure([]string{"format"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphFunctionConf() *Configure {
	return NewConfigure([]string{"function"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphGridConf() *Configure {
	return NewConfigure([]string{"grid"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphHidden3dConf() *Configure {
	return NewConfigure([]string{"hidden3d"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphHistoryConf() *Configure {
	return NewConfigure([]string{"history"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphHistorysizeConf() *Configure {
	return NewConfigure([]string{"historysize"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphIsosamplesConf() *Configure {
	return NewConfigure([]string{"isosamples"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphKeyConf() *Configure {
	return NewConfigure([]string{"key"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphLabelConf() *Configure {
	return NewConfigure([]string{"label"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphLinetypeConf() *Configure {
	return NewConfigure([]string{"label"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphLinkConf() *Configure {
	return NewConfigure([]string{"link"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphLmarginConf() *Configure {
	return NewConfigure([]string{"lmargin"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphLoadpathConf() *Configure {
	return NewConfigure([]string{"loadpath"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphLocaleConf() *Configure {
	return NewConfigure([]string{"locale"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphLogConf() *Configure {
	return NewConfigure([]string{"log"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphLogscaleConf() *Configure {
	return NewConfigure([]string{"logscale"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMacrosConf() *Configure {
	return NewConfigure([]string{"macros"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMappingConf() *Configure {
	return NewConfigure([]string{"mapping"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMarginConf() *Configure {
	return NewConfigure([]string{"margin"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMarginsConf() *Configure {
	return NewConfigure([]string{"margins"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMissingConf() *Configure {
	return NewConfigure([]string{"missing"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMonochromeConf() *Configure {
	return NewConfigure([]string{"monochrome"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMouseConf() *Configure {
	return NewConfigure([]string{"mouse"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMultiplotConf() *Configure {
	return NewConfigure([]string{"multiplot"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMx2ticsConf() *Configure {
	return NewConfigure([]string{"mx2tics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMxticsConf() *Configure {
	return NewConfigure([]string{"mxtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMy2ticsConf() *Configure {
	return NewConfigure([]string{"my2tics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMyticsConf() *Configure {
	return NewConfigure([]string{"mytics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphMzticsConf() *Configure {
	return NewConfigure([]string{"mztics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphObjectConf() *Configure {
	return NewConfigure([]string{"object"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphOffsetsConf() *Configure {
	return NewConfigure([]string{"offsets"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphOriginConf() *Configure {
	return NewConfigure([]string{"origin"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphOutputConf() *Configure {
	return NewConfigure([]string{"output"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphPaletteConf() *Configure {
	return NewConfigure([]string{"palette"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphParametricConf() *Configure {
	return NewConfigure([]string{"parametric"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphPaxisConf() *Configure {
	return NewConfigure([]string{"paxis"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphPm3dConf() *Configure {
	return NewConfigure([]string{"pm3d"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphPointintervalboxConf() *Configure {
	return NewConfigure([]string{"pointintervalbox"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphPointsizeConf() *Configure {
	return NewConfigure([]string{"pointsize"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphPolarConf() *Configure {
	return NewConfigure([]string{"polar"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphPrintConf() *Configure {
	return NewConfigure([]string{"print"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphPsdirConf() *Configure {
	return NewConfigure([]string{"psdir"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphRaxisConf() *Configure {
	return NewConfigure([]string{"raxis"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphRmarginConf() *Configure {
	return NewConfigure([]string{"rmargin"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphRrangeConf() *Configure {
	return NewConfigure([]string{"rrangle"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphRticsConf() *Configure {
	return NewConfigure([]string{"rtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphSamplesConf() *Configure {
	return NewConfigure([]string{"samples"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphSizeConf() *Configure {
	return NewConfigure([]string{"size"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphStyleConf() *Configure {
	return NewConfigure([]string{"style"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphSurfaceConf() *Configure {
	return NewConfigure([]string{"surface"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTableConf() *Configure {
	return NewConfigure([]string{"table"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTermConf() *Configure {
	return NewConfigure([]string{"term"}, []string{}, func(vals []string) bool {
		return len(vals) == 1 &&
			utils.InStr(vals[0], []string{
				"cairolatex", "canvas", "cgm", "context",
				"corel", "dump", "dxf", "eepic",
				"emf", "emtex", "epscairo", "epslatex",
				"fig", "gif", "hpgl", "jpeg",
				"latex", "lua", "mf", "mp",
				"pcl5", "pdfcairo", "png", "pngcairo",
				"pop", "postscript", "pslatex", "pstex",
				"pstricks", "push", "qms", "qt",
				"svg", "tek40xx", "tek410x", "texdraw",
				"tgif", "tikz", "tkcanvas", "tpic",
				"vttek", "wxt", "xt11", "xlib",
				"xterm"})
	})
	// TODO: in case of len(vals) >= 2
}

func GraphTerminalConf() *Configure {
	return NewConfigure([]string{"terminal"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTermoptionConf() *Configure {
	return NewConfigure([]string{"termoption"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTicsConf() *Configure {
	return NewConfigure([]string{"tics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTicscaleConf() *Configure {
	return NewConfigure([]string{"ticscale"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTicslevelcaleConf() *Configure {
	return NewConfigure([]string{"ticslevel"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTimeSpecifiersConf() *Configure {
	return NewConfigure([]string{"time_specifiers"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTimefmtConf() *Configure {
	return NewConfigure([]string{"timefmt"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTimestampConf() *Configure {
	return NewConfigure([]string{"timestamp"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTitleConf() *Configure {
	return NewConfigure([]string{"title"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTmarginConf() *Configure {
	return NewConfigure([]string{"tmargin"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphTrangeConf() *Configure {
	return NewConfigure([]string{"trange"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphUrangeConf() *Configure {
	return NewConfigure([]string{"urange"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphViewConf() *Configure {
	return NewConfigure([]string{"view"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphVrangeConf() *Configure {
	return NewConfigure([]string{"vrange"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphX2dataConf() *Configure {
	return NewConfigure([]string{"x2data"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphX2dticsConf() *Configure {
	return NewConfigure([]string{"x2dtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphX2labelConf() *Configure {
	return NewConfigure([]string{"x2label"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphX2mticsConf() *Configure {
	return NewConfigure([]string{"x2mtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphX2rangeConf() *Configure {
	return NewConfigure([]string{"x2range"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphX2ticsConf() *Configure {
	return NewConfigure([]string{"x2tics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphX2zeroaxisConf() *Configure {
	return NewConfigure([]string{"x2zeroaxis"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphXdataConf() *Configure {
	return NewConfigure([]string{"xdata"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphXdticsConf() *Configure {
	return NewConfigure([]string{"xdtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphXlabelConf() *Configure {
	return NewConfigure([]string{"label"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphXmticsConf() *Configure {
	return NewConfigure([]string{"xmtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphXrangeConf() *Configure {
	return NewConfigure([]string{"xrange"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphXticsConf() *Configure {
	return NewConfigure([]string{"xtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphXyplaneConf() *Configure {
	return NewConfigure([]string{"xyplane"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphXzeroaxisConf() *Configure {
	return NewConfigure([]string{"xzeroaxis"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphY2dataConf() *Configure {
	return NewConfigure([]string{"y2data"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphY2dticsConf() *Configure {
	return NewConfigure([]string{"y2dtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphY2dlabelConf() *Configure {
	return NewConfigure([]string{"y2label"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphY2mticsConf() *Configure {
	return NewConfigure([]string{"y2mtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphY2rangeConf() *Configure {
	return NewConfigure([]string{"y2range"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphY2ticsConf() *Configure {
	return NewConfigure([]string{"y2tics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphY2zeroaxisConf() *Configure {
	return NewConfigure([]string{"y2zeroaxis"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphYdataConf() *Configure {
	return NewConfigure([]string{"ydata"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphYdticsConf() *Configure {
	return NewConfigure([]string{"ydtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphYlabelConf() *Configure {
	return NewConfigure([]string{"ylabel"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphYmticsConf() *Configure {
	return NewConfigure([]string{"ymtics"}, []string{}, func(vals []string) bool {
		return true
	})
}

func GraphYrangeConf() *Configure {
	return NewConfigure([]string{"yrange"}, []string{}, func(vals []string) bool {
		return true
	})
}
