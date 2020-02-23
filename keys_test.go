package dot

import "testing"

func TestKeys(t *testing.T) {
	if len(Keys(nil)) != 0 {
		t.Fail()
	}

	if len(Keys(4)) != 0 {
		t.Fail()
	}

	mapTest := map[string]interface{}{
		"one": 1,
		"two": 3,
		"3":   4,
		"a":   "h",
		"d": map[string]interface{}{
			"one": 1,
			"two": 3,
			"3":   4,
			"a":   "h",
		},
		"j": nil,
	}

	// try root-level keys without a parent path
	keysFromMap := Keys(mapTest)

	if len(keysFromMap) != 6 {
		t.Fail()
	}

	if !contains(keysFromMap, "one") || !contains(keysFromMap, "two") || !contains(keysFromMap, "3") || !contains(keysFromMap, "a") || !contains(keysFromMap, "d") {
		t.Fail()
	}

	// now, try root-level keys again with a parent path
	keysFromMap = Keys(mapTest, "root.data")

	if len(keysFromMap) != 6 {
		t.Fail()
	}

	if !contains(keysFromMap, "root.data.one") || !contains(keysFromMap, "root.data.two") ||
		!contains(keysFromMap, "root.data.3") || !contains(keysFromMap, "root.data.a") ||
		!contains(keysFromMap, "root.data.d") {

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

func TestKeysRecursive(t *testing.T) {
	mapTest := map[string]interface{}{
		"one": 1,
		"two": 3,
		"3":   4,
		"a":   "h",
		"d": map[string]interface{}{
			"one": 1,
			"two": 3,
			"3":   4,
			"a":   "h",
		},
		"j": nil,
	}
	// try recursive keys without a parent path
	deepKeysFromMap := KeysRecursive(mapTest)
	if len(deepKeysFromMap) != 10 {
		t.Fail()
	}

	if !contains(deepKeysFromMap, "one") || !contains(deepKeysFromMap, "two") || !contains(deepKeysFromMap, "3") ||
		!contains(deepKeysFromMap, "a") || !contains(deepKeysFromMap, "d") || !contains(deepKeysFromMap, "d.one") ||
		!contains(deepKeysFromMap, "d.two") || !contains(deepKeysFromMap, "d.3") || !contains(deepKeysFromMap, "d.a") ||
		!contains(deepKeysFromMap, "j") {

		t.Fail()
	}

	// try recursive keys with a parent path
	deepKeysFromMap = KeysRecursive(mapTest, "top.root.node")
	if len(deepKeysFromMap) != 10 {
		t.Fail()
	}

	if !contains(deepKeysFromMap, "top.root.node.one") || !contains(deepKeysFromMap, "top.root.node.two") || !contains(deepKeysFromMap, "top.root.node.3") ||
		!contains(deepKeysFromMap, "top.root.node.a") || !contains(deepKeysFromMap, "top.root.node.d") || !contains(deepKeysFromMap, "top.root.node.d.one") ||
		!contains(deepKeysFromMap, "top.root.node.d.two") || !contains(deepKeysFromMap, "top.root.node.d.3") || !contains(deepKeysFromMap, "top.root.node.d.a") ||
		!contains(deepKeysFromMap, "top.root.node.j") {

		t.Fail()
	}

	type TestInnerStruct struct {
		F bool
	}

	type TestStruct struct {
		A bool
		B map[string]interface{}
		C string
		D int64
		E []TestInnerStruct
	}

	testStruct := TestStruct{
		A: false,
		B: map[string]interface{}{
			"A": 1,
		},
		C: "4",
		D: 7,
		E: []TestInnerStruct{
			{
				F: true,
			},
		},
	}

	keysFromStruct := KeysRecursive(testStruct)

	if len(keysFromStruct) != 6 {
		t.Fail()
	}

	if !contains(keysFromStruct, "A") || !contains(keysFromStruct, "B") || !contains(keysFromStruct, "C") || !contains(keysFromStruct, "D") {
		t.Fail()
	}

	// keys from nested arrays should not be included
	if contains(keysFromStruct, "E.F") {
		t.Fail()
	}

	// now, try again with a parent path
	keysFromStruct = KeysRecursive(testStruct, "data.raw")

	if len(keysFromStruct) != 6 {
		t.Fail()
	}

	if !contains(keysFromStruct, "data.raw.A") || !contains(keysFromStruct, "data.raw.B") || !contains(keysFromStruct, "data.raw.C") || !contains(keysFromStruct, "data.raw.D") {
		t.Fail()
	}
}

func TestKeysRecursiveLeaves(t *testing.T) {
	mapTest := map[string]interface{}{
		"one": 1,
		"two": 3,
		"3":   4,
		"a":   "h",
		"d": map[string]interface{}{
			"6": 1,
			"7": 3,
			"8": 4,
			"z": "h",
		},
		"j": nil,
	}
	// try recursive keys without a parent path
	deepKeysFromMap := KeysRecursiveLeaves(mapTest)
	if len(deepKeysFromMap) != 9 {
		t.Fail()
	}

	if !contains(deepKeysFromMap, "one") || !contains(deepKeysFromMap, "two") || !contains(deepKeysFromMap, "3") ||
		!contains(deepKeysFromMap, "a") || !contains(deepKeysFromMap, "d.6") ||
		!contains(deepKeysFromMap, "d.7") || !contains(deepKeysFromMap, "d.8") || !contains(deepKeysFromMap, "d.z") ||
		!contains(deepKeysFromMap, "j") {

		t.Fail()
	}

	// try recursive keys with a parent path
	deepKeysFromMap = KeysRecursiveLeaves(mapTest, "top.root.node")
	if len(deepKeysFromMap) != 9 {
		t.Fail()
	}

	if !contains(deepKeysFromMap, "top.root.node.one") || !contains(deepKeysFromMap, "top.root.node.two") || !contains(deepKeysFromMap, "top.root.node.3") ||
		!contains(deepKeysFromMap, "top.root.node.a") || !contains(deepKeysFromMap, "top.root.node.d.6") ||
		!contains(deepKeysFromMap, "top.root.node.d.7") || !contains(deepKeysFromMap, "top.root.node.d.8") || !contains(deepKeysFromMap, "top.root.node.d.z") ||
		!contains(deepKeysFromMap, "top.root.node.j") {

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
