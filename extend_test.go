package dot

import "testing"

func TestExtend(t *testing.T) {
	x := map[string]interface{} {
		"A": 1,
		"C": map[string]interface{} {
			"Z": 3,
		},
		"F": false,
		"G": true,
		"H": "base",
		"I": "",
		"J": 0,
		"K": 9,
		"L": 0.0,
		"M": 18.9,
	}

	y := map[string]interface{} {
		"B": 8,
		"C": map[string]interface{} {
			"X": 6,
		},
		"D": map[string]interface{} {
			"Y": 2,
		},
		"F": true,
		"G": false,
		"H": "",
		"I": "over",
		"J": 12,
		"K": 0,
		"L": 9.1,
		"M": 0.0,
	}

	err := Extend(x, y)
	if err != nil {
		t.Fail()
	}

	if GetString(x, "A") != "1" {
		t.Fail()
	}

	if GetString(x, "B") != "8" {
		t.Fail()
	}

	cMap, _ := Get(x, "C")
	if GetString(cMap, "Z") != "3" {
		t.Fail()
	}
	if GetString(cMap, "X") != "6" {
		t.Fail()
	}

	f, _ := Get(x, "F")
	if f == nil {
		t.Fail()
	}
	fB := f.(bool)
	if !fB {
		t.Fail()
	}

	g, _ := Get(x, "G")
	if g == nil {
		t.Fail()
	}
	gB := f.(bool)
	if !gB {
		t.Fail()
	}

	// ensure "" doesn't clobber a meaningful value
	h := GetString(x, "H")
	if h != "base" {
		t.Fail()
	}

	// ensure a meaningful value can clobber ""
	i := GetString(x, "I")
	if i != "over" {
		t.Fail()
	}

	// ensure 0 doesn't clobber a meaningful value
	j := GetInt64(x, "J")
	if j != 12 {
		t.Fail()
	}

	// ensure a meaningful value can clobber 0
	k := GetInt64(x, "K")
	if k != 9 {
		t.Fail()
	}

	// ensure 0.0 doesn't clobber a meaningful value
	l := GetFloat64(x, "L")
	if l != 9.1 {
		t.Fail()
	}

	// ensure 0.0 doesn't clobber a meaningful value
	m := GetFloat64(x, "M")
	if m != 18.9 {
		t.Fail()
	}
}
