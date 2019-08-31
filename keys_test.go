package dot

import "testing"

func TestKeys(t *testing.T) {
	mapTest := map[string]interface{} {
		"one": 1,
		"two": 3,
		"3": 4,
		"a": "h",
		"d": map[string]interface{} {
			"one": 1,
			"two": 3,
			"3": 4,
			"a": "h",
		},
	}

	// try without a parent path
	keysFromMap := Keys(mapTest)

	if len(keysFromMap) != 5 {
		t.Fail()
	}

	if !contains(keysFromMap, "one") || !contains(keysFromMap, "two") || !contains(keysFromMap, "3") || !contains(keysFromMap, "a") || !contains(keysFromMap, "d") {
		t.Fail()
	}

	// now, try again with a parent path
	keysFromMap = Keys(mapTest, "root.data")

	if len(keysFromMap) != 5 {
		t.Fail()
	}

	if !contains(keysFromMap, "root.data.one") || !contains(keysFromMap, "root.data.two") || !contains(keysFromMap, "root.data.3") || !contains(keysFromMap, "root.data.a") || !contains(keysFromMap, "root.data.d") {
		t.Fail()
	}

	type TestStruct struct {
		A bool
		B map[string]interface{}
		C string
		D int64
	}

	testStruct := TestStruct{
		A: false,
		B: map[string]interface{}{
			"A": 1,
		},
		C: "4",
		D: 7,
	}

	keysFromStruct := Keys(testStruct)

	if len(keysFromStruct) != 4 {
		t.Fail()
	}

	if !contains(keysFromStruct, "A") || !contains(keysFromStruct, "B") || !contains(keysFromStruct, "C") || !contains(keysFromStruct, "D") {
		t.Fail()
	}

	// now, try again with a parent path
	keysFromStruct = Keys(testStruct, "data.raw")

	if len(keysFromStruct) != 4 {
		t.Fail()
	}

	if !contains(keysFromStruct, "data.raw.A") || !contains(keysFromStruct, "data.raw.B") || !contains(keysFromStruct, "data.raw.C") || !contains(keysFromStruct, "data.raw.D") {
		t.Fail()
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}