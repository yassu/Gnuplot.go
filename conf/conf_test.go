package conf

import (
	"github.com/yassu/gnuplot.go/utils"
	"testing"
)

// utils
func TestInStr0(t *testing.T) {
	if utils.InStr("a", []string{}) != false {
		t.Errorf("fals in TestInStr0")
	}
}

func TestInStr1(t *testing.T) {
	if utils.InStr("a", []string{"a", "b", "c"}) != true {
		t.Errorf("fals in TestInStr1")
	}
}

func TestInStr2(t *testing.T) {
	if utils.InStr("c", []string{"a", "b", "c"}) != true {
		t.Errorf("fals in TestInStr2")
	}
}

func TestInStr3(t *testing.T) {
	if utils.InStr("b", []string{"a", "b", "c"}) != true {
		t.Errorf("fals in TestInStr3")
	}
}

func TestInStr4(t *testing.T) {
	if utils.InStr("d", []string{"a", "b", "c"}) != false {
		t.Errorf("fals in TestInStr4")
	}
}

func TestIsNum1(t *testing.T) {
	if isNum("0") != true {
		t.Errorf("falis in TestIsNum1")
	}
}

func TestIsNum2(t *testing.T) {
	if isNum("+2") != true {
		t.Errorf("falis in TestIsNum2")
	}
}

func TestIsNum3(t *testing.T) {
	if isNum("+2.3") != true {
		t.Errorf("falis in TestIsNum3")
	}
}

func TestIsNum4(t *testing.T) {
	if isNum("2.3.5") != false {
		t.Errorf("falis in TestIsNum4")
	}
}

func TestIsNum5(t *testing.T) {
	if isNum("-2") != true {
		t.Errorf("falis in TestIsNum5")
	}
}

func TestIsNum6(t *testing.T) {
	if isNum("-2.8") != true {
		t.Errorf("falis in TestIsNum6")
	}
}

func TestIsNum7(t *testing.T) {
	if isNum("-2.8.3") != false {
		t.Errorf("falis in TestIsNum7")
	}
}

func TestIsSixHex(t *testing.T) {
	if isSixHex("0") != false {
		t.Errorf("fails in TestIsSixHex")
	}
}

func TestIsSixHex2(t *testing.T) {
	if isSixHex("000000") != true {
		t.Errorf("fails in TestIsSixHex2")
	}
}

func TestIsSixHex3(t *testing.T) {
	if isSixHex("00000") != false {
		t.Errorf("fails in TestIsSixHex3")
	}
}

func TestIsEightHex(t *testing.T) {
	if isEightHex("0") != false {
		t.Errorf("fails in TestIsEightHex")
	}
}

func TestIsEightHex2(t *testing.T) {
	if isEightHex("00000000") != true {
		t.Errorf("fails in TestIsEightHex2")
	}
}

func TestIsEightHex3(t *testing.T) {
	if isEightHex("0000000") != false {
		t.Errorf("fails in TestIsEightHex3")
	}
}

func TestIsSmallFloat0(t *testing.T) {
	if isSmallFloat("a") != false {
		t.Errorf("fails in TestIsSmallFloat")
	}
}

func TestIsSmallFloat(t *testing.T) {
	if isSmallFloat("0") != true {
		t.Errorf("fails in TestIsSmallFloat")
	}
}

func TestIsSmallFloat2(t *testing.T) {
	if isSmallFloat("1") != true {
		t.Errorf("fails in TestIsSmallFloat2")
	}
}

func TestIsSmallFloat3(t *testing.T) {
	if isSmallFloat("0.3") != true {
		t.Errorf("fails in TestIsSmallFloat3")
	}
}

// Configure Class
func NewConfigureTest(t *testing.T) {
	conf := PStyleWithConf()

	if conf.key != "with" {
		t.Errorf("fails in key test of NewConfigureTest")
	}

	if len(conf.aliasKeys) != 2 || conf.aliasKeys[0] != "with" || conf.aliasKeys[1] != "w" {
		t.Errorf("fails in aliasKeys test of NewConfigureTest")
	}

	if len(conf.vals) != 1 || conf.vals[0] != "lines" {
		t.Errorf("fails in vals test of NewConfigureTest")
	}

	if conf.requiredCondition([]string{}) != false {
		t.Errorf("fails in requiredCondition test of NewConfigureTest")
	}
}

func TestConfigureSetVals(t *testing.T) {
	conf := NewConfigure([]string{"key1", "key2"}, []string{"val1", "val2"}, func(vals []string) bool {
		return true
	})
	conf.SetVals([]string{"abc"})
	vals := conf.vals
	if len(vals) != 1 || vals[0] != "abc" {
		t.Errorf("fails in TestConfigureSetVals")
	}
}

func TestConfigureGetKey(t *testing.T) {
	conf := NewConfigure([]string{"key1", "key2"}, []string{"val1", "val2"}, func(vals []string) bool {
		return true
	})
	if conf.GetKey() != "key1" {
		t.Errorf("fails in TestConfigureGetKey")
	}
}

// Validation of Configurations

// for Plot Element
func TestPStyleWithConfValidation(t *testing.T) {
	conf := PStyleWithConf()
	if conf.requiredCondition([]string{"dots"}) != true {
		t.Errorf("fails in TestPStyleWithConfValidation")
	}
}

func TestPStyleWithConfValidation2(t *testing.T) {
	conf := PStyleWithConf()
	if conf.requiredCondition([]string{"dot"}) != false {
		t.Errorf("fails in TestPStyleWithConfValidation2")
	}
}

func TestPStyleWithConfValidation3(t *testing.T) {
	conf := PStyleWithConf()
	if conf.requiredCondition([]string{""}) != false {
		t.Errorf("fails in TestPStyleWithConfValidation3")
	}
}

func TestPStyleWithConfValidation4(t *testing.T) {
	conf := PStyleWithConf()
	if conf.requiredCondition([]string{"dot", "lines"}) != false {
		t.Errorf("fails in TestPStyleWithConfValidation4")
	}
}

func TestPStyleLineColorConfValidation(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"blue"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation")
	}
}

func TestPStyleLineColorConfValidation2(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"dummy"}) != false {
		t.Errorf("fails in TestPStyleLineColorConfValidation2")
	}
}

func TestPStyleLineColorConfValidation3(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"rgbcolor", "0x000000"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation3")
	}
}

func TestPStyleLineColorConfValidation4(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"rgbcolor", "0x00000000"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation4")
	}
}

func TestPStyleLineColorConfValidation5(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"rgbcolor", "#000000"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation5")
	}
}

func TestPStyleLineColorConfValidation6(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"rgbcolor", "#00000000"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation6")
	}
}

func TestPStyleLineColorConfValidation7(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"rgbcolor", "0"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation7")
	}
}

func TestPStyleLineColorConfValidation8(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"rgbcolor", ""}) != false {
		t.Errorf("fails in TestPStyleLineColorConfValidation8")
	}
}

func TestPStyleLineColorConfValidation9(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"rgbcolor", "variable"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation9")
	}
}

func TestPStyleLineColorConfValidation10(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"rgbcolor", "variable", ""}) != false {
		t.Errorf("fails in TestPStyleLineColorConfValidation10")
	}
}

func TestPStyleLineColorConfValidation11(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"palette", "frac", "0.0"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation11")
	}
}

func TestPStyleLineColorConfValidation12(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"palette", "frac", "1.0"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation12")
	}
}

func TestPStyleLineColorConfValidation13(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"palette", "frac", "3.0"}) != false {
		t.Errorf("fails in TestPStyleLineColorConfValidation13")
	}
}

func TestPStyleLineColorConfValidation14(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"palette", "cb", "-5.3"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation14")
	}
}

func TestPStyleLineColorConfValidation15(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"palette", "cb", "-5.3", ""}) != false {
		t.Errorf("fails in TestPStyleLineColorConfValidation15")
	}
}

func TestPStyleLineColorConfValidation16(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"palette", "z"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation16")
	}
}

func TestPStyleLineColorConfValidation17(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"palette", "z", ""}) != false {
		t.Errorf("fails in TestPStyleLineColorConfValidation17")
	}
}

func TestPStyleLineColorConfValidation18(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"variable"}) != true {
		t.Errorf("fails in TestPStyleLineColorConfValidation18")
	}
}

func TestPStyleLineColorConfValidation19(t *testing.T) {
	conf := PStyleLineColorConf()
	if conf.requiredCondition([]string{"variable", "z"}) != false {
		t.Errorf("fails in TestPStyleLineColorConfValidation19")
	}
}

func TestPStylePointTypeConf(t *testing.T) {
	conf := PStylePointTypeConf()
	if conf.requiredCondition([]string{"1"}) != true {
		t.Errorf("fails in TestPStylePointTypeConf")
	}
}

func TestPStyleLineTypeConf2(t *testing.T) {
	conf := PStylePointTypeConf()
	if conf.requiredCondition([]string{"1", ""}) != false {
		t.Errorf("fails in TestPLinePointTypeConf2")
	}
}

func TestPStyleLineTypeConf(t *testing.T) {
	conf := PStyleLineTypeConf()
	if conf.requiredCondition([]string{"1"}) != true {
		t.Errorf("fails in TestPStyleLineTypeConf")
	}
}

func TestPStylePointTypeConf2(t *testing.T) {
	conf := PStyleLineTypeConf()
	if conf.requiredCondition([]string{"1", ""}) != false {
		t.Errorf("fails in TestPStylePointTypeConf2")
	}
}

func TestPStyleLineWidthConf(t *testing.T) {
	conf := PStyleLineWidthConf()
	if conf.requiredCondition([]string{"2.3"}) != true {
		t.Errorf("fails in TestPStyleLineWidthConf")
	}
}

func TestPStylePointWidthConf2(t *testing.T) {
	conf := PStyleLineWidthConf()
	if conf.requiredCondition([]string{"2.3", ""}) != false {
		t.Errorf("fails in TestPStylePointWidthConf2")
	}
}

func TestPStylePointSizeConf(t *testing.T) {
	conf := PStylePointSizeConf()
	if conf.requiredCondition([]string{"2.3"}) != true {
		t.Errorf("fails in TestPStylePointSizeConf")
	}
}

func TestPStylePointSizeConf2(t *testing.T) {
	conf := PStylePointSizeConf()
	if conf.requiredCondition([]string{"2.3", ""}) != false {
		t.Errorf("fails in TestPStylePointSizeConf2")
	}
}

func TestPStyleNoHidden3dConf(t *testing.T) {
	conf := PStyleNoHidden3dConf()
	if conf.requiredCondition([]string{"true"}) != true {
		t.Errorf("fails in TestPStyleNoHidden3dConf")
	}
}

func TestPStyleNoHidden3dConf2(t *testing.T) {
	conf := PStyleNoHidden3dConf()
	if conf.requiredCondition([]string{"abc"}) != false {
		t.Errorf("fails in TestPStyleNoHidden3dConf2")
	}
}

func TestPStyleNoHidden3dConf3(t *testing.T) {
	conf := PStyleNoHidden3dConf()
	if conf.requiredCondition([]string{"true", "false"}) != false {
		t.Errorf("fails in TestPStyleNoHidden3dConf3")
	}
}

func TestPStyleNoContoursConf(t *testing.T) {
	conf := PStyleNoContoursConf()
	if conf.requiredCondition([]string{"true"}) != true {
		t.Errorf("fails in TestPStyleNoContoursConf")
	}
}

func TestPStyleNoContoursConf2(t *testing.T) {
	conf := PStyleNoContoursConf()
	if conf.requiredCondition([]string{"abc"}) != false {
		t.Errorf("fails in TestPStyleNoContoursConf2")
	}
}

func TestPStyleNoContoursConf3(t *testing.T) {
	conf := PStyleNoContoursConf()
	if conf.requiredCondition([]string{"true", "true"}) != false {
		t.Errorf("fails in TestPStyleNoContoursConf3")
	}
}

func TestPStyleNoSurfaceConf(t *testing.T) {
	conf := PStyleNoSurfaceConf()
	if conf.requiredCondition([]string{"true"}) != true {
		t.Errorf("fails in TestPStyleNoSurfaceConf")
	}
}

func TestPStyleNoSurfaceConf2(t *testing.T) {
	conf := PStyleNoSurfaceConf()
	if conf.requiredCondition([]string{"abc"}) != false {
		t.Errorf("fails in TestPStyleNoSurfaceConf2")
	}
}

func TestPStyleNoSurfaceConf3(t *testing.T) {
	conf := PStyleNoSurfaceConf()
	if conf.requiredCondition([]string{"true", "true"}) != false {
		t.Errorf("fails in TestPStyleNoSurfaceConf3")
	}
}

func TestPStylePaletteConf(t *testing.T) {
	conf := PStylePaletteConf()
	if conf.requiredCondition([]string{"true"}) != true {
		t.Errorf("fails in TestPStylePaletteConf")
	}
}

func TestPStylePaletteConf2(t *testing.T) {
	conf := PStylePaletteConf()
	if conf.requiredCondition([]string{"abc"}) != false {
		t.Errorf("fails in TestPStyleNoSurfaceConf2")
	}
}

func TestPStylePaletteConf3(t *testing.T) {
	conf := PStylePaletteConf()
	if conf.requiredCondition([]string{"true", "true"}) != false {
		t.Errorf("fails in TestPStylePaletteConf3")
	}
}

func TestPStyleGoXMinConf(t *testing.T) {
	conf := PStyleGoXMaxConf()
	if conf.requiredCondition([]string{"-3.2"}) != true {
		t.Errorf("fails in PStyleGoXMaxConf")
	}
}

func TestPStyleGoXMinConf2(t *testing.T) {
	conf := PStyleGoXMaxConf()
	if conf.requiredCondition([]string{"+3.2.3"}) != false {
		t.Errorf("fails in PStyleGoXMaxConf2")
	}
}

func TestPStyleGoXMinConf3(t *testing.T) {
	conf := PStyleGoXMaxConf()
	if conf.requiredCondition([]string{"3.2", "3"}) != false {
		t.Errorf("fails in PStyleGoXMaxConf3")
	}
}

// for Graph Element
func TestGraphTermConf(t *testing.T) {
	conf := GraphTermConf()
	if conf.requiredCondition([]string{"pngs"}) != false {
		t.Errorf("fails in TestGraphTermConf")
	}
}

func TestGraphTermConf2(t *testing.T) {
	conf := GraphTermConf()
	if conf.requiredCondition([]string{"png", "eps"}) != false {
		t.Errorf("fails in TestGraphTermConf2")
	}
}

func TestGraphTermConf3(t *testing.T) {
	conf := GraphTermConf()
	if conf.requiredCondition([]string{"png"}) != true {
		t.Errorf("fails in TestGraphTermConf3")
	}
}
