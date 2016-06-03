package conf

import (
	"testing"
)

func TestInStr0(t *testing.T) {
	if inStr("a", []string{}) != false {
		t.Errorf("fals in TestInStr0")
	}
}

func TestInStr1(t *testing.T) {
	if inStr("a", []string{"a", "b", "c"}) != true {
		t.Errorf("fals in TestInStr1")
	}
}

func TestInStr2(t *testing.T) {
	if inStr("c", []string{"a", "b", "c"}) != true {
		t.Errorf("fals in TestInStr2")
	}
}

func TestInStr3(t *testing.T) {
	if inStr("b", []string{"a", "b", "c"}) != true {
		t.Errorf("fals in TestInStr3")
	}
}

func TestInStr4(t *testing.T) {
	if inStr("d", []string{"a", "b", "c"}) != false {
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
