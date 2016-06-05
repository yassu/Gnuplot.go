package conf

import (
	"github.com/yassu/gnuplot.go/utils"
	"testing"
)

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

func NewConfigureTest(t *testing.T) {
	conf := WithConf()

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
