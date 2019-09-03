package dot

import "testing"

func TestExtend(t *testing.T) {
	x := map[string]interface{} {
		"A": 1,
		"C": map[string]interface{} {
			"Z": 3,
		},
	}

	y := map[string]interface{} {
		"B": 8,
		"C": map[string]interface{} {
			"X": 6,
		},
		"D": map[string]interface{} {
			"Y": 2,
		},
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
}
