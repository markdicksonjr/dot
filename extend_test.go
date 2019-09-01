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

	z, err := Extend(x, y)
	if err != nil {
		t.Fail()
	}

	zMap, ok := z.(map[string]interface{})
	if !ok {
		t.Fail()
	}

	if GetString(zMap, "A") != "1" {
		t.Fail()
	}

	if GetString(zMap, "B") != "8" {
		t.Fail()
	}

	cMap, ok := zMap["C"].(map[string]interface{})
	if GetString(cMap, "Z") != "3" {
		t.Fail()
	}
	if GetString(cMap, "X") != "6" {
		t.Fail()
	}
}
