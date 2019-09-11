package dot

import (
	"math"
	"testing"
)

func TestGet_Struct(t *testing.T) {
	type SampleStructA struct {
		Text string
	}

	type SampleStruct struct {
		A     SampleStructA
		B     int
		C     string
		Slice []int
	}

	sample := SampleStruct{
		A: SampleStructA{
			Text: "sf",
		},
		B:     8,
		C:     "something",
		Slice: []int{5, 6},
	}

	// nested get, make "a" lowercase to demonstrate case-insensitivity
	aText, err := Get(sample, "a.Text")
	if err != nil {
		t.Fail()
	}
	if aText == nil {
		t.Fatal("dot-prop get failed")
	}
	if aText.(string) != "sf" {
		t.Fatal("failed to dot-prop get string")
	}

	// root get of base type
	bInt, err := Get(sample, "B")
	if err != nil {
		t.Fail()
	}
	if bInt == nil {
		t.Fatal("dot-prop get root int failed")
	}
	if bInt.(int) != 8 {
		t.Fatal("failed to dot-prop get root int")
	}

	// root get of string
	cText, err := Get(sample, "C")
	if err != nil {
		t.Fail()
	}
	if cText == nil {
		t.Fatal("dot-prop get failed")
	}
	if cText.(string) != "something" {
		t.Fatal("failed to dot-prop get root string")
	}

	// root get of missing prop
	_, err = Get(sample, "D")
	if err == nil {
		t.Fail()
	}

	// root get of missing sub-prop
	_, err = Get(sample, "A.x")
	if err == nil {
		t.Fail()
	}

	// get on a nil object
	rootNil, err := Get(nil, "e")
	if err != nil {
		t.Fail()
	}
	if rootNil != nil {
		t.Fatal("got something that should from root have been nil")
	}

	// test fallback where first property is found, make A upper-case to demonstrate
	// case insensitivity, along with the assertions above on "a"
	firstProp, err := Get(sample, "A.Text", "B")
	if err != nil {
		t.Fail()
	}
	if firstProp == nil {
		t.Fatal("dot-prop fallback get failed")
	}
	if firstProp.(string) != "sf" {
		t.Fatal("failed to dot-prop get string with fallback")
	}

	// test fallback where first property is nil
	secondProp, err := Get(sample, "A.bogus", "B", "C")
	if err != nil {
		t.Fail()
	}
	if secondProp == nil {
		t.Fatal("dot-prop fallback get failed")
	}
	if secondProp.(int) != 8 {
		t.Fatal("failed to dot-prop get string with fallback")
	}

	// test fallback where no property is available
	_, err = Get(sample, "A.bogus", "bogus", "foo")
	if err == nil {
		t.Fail()
	}

	// test getting a slice from the object
	_, err = Get(sample, "Slice")
	if err != nil {
		t.Fail()
	}
}

func BenchmarkGet_Map(b *testing.B) {
	data := map[string]interface{} {
		"A": map[string]interface{} {
			"Z": 1,
			"AA": 0,
		},
	}

	aGet, _ := Get(data, "A")
	if aGet == nil {
		b.Fatal("got nil 'A'")
	}

	azGet, _ := Get(data, "A.Z")
	if azGetInt, _ := CoerceInt64(azGet); azGetInt != 1 {
		b.Fatal("A.Z was not 1")
	}
}

func TestGetInt64(t *testing.T) {
	data := make(map[string]interface{})
	a := make(map[string]interface{})
	a["b"] = 5
	a["c"] = 6.5
	a["d"] = int32(7)
	a["e"] = "should be 0"
	a["f"] = int64(9)
	a["g"] = false
	data["a"] = a

	// test coercion of int
	if GetInt64(data, "a.b") != 5 {
		t.Error("a.b was not 5")
	}

	// test simple get of float flooring
	if GetInt64(data, "a.c") != 6 {
		t.Error("a.c was not 6")
	}

	// test coercion of float32
	if GetInt64(data, "a.d") != 7 { // will be some rounding error in float32
		t.Error("a.d was not 7")
	}

	// test that incompatible coercion results in 0
	if GetInt64(data, "a.e") != 0 {
		t.Error("a.e was not 0")
	}

	// test that coercion on nil results in 0
	if GetInt64(nil, "a.e") != 0 {
		t.Error("GetInt64 on nil was not 0")
	}

	// test simple get of target type
	if GetInt64(data, "a.g") != 0 {
		t.Error("a.c was not 0")
	}

	// test that coercion on explicit int64
	if GetInt64(data, "a.f") != 9 {
		t.Error("GetInt64 on explicitly int64 not correct")
	}

	// test fallback
	if GetInt64(data, "x", "a.b") != 5 {
		t.Error("fallback did not fall back to the correct value")
	}
}

func TestGetFloat64(t *testing.T) {
	data := make(map[string]interface{})
	a := make(map[string]interface{})
	a["b"] = 5
	a["c"] = 6.5
	a["d"] = float32(7.8)
	a["e"] = "should be 0"
	data["a"] = a

	// test coercion of int
	if GetFloat64(data, "a.b") != 5 {
		t.Error("a.b was not 5")
	}

	// test simple get of target type
	if GetFloat64(data, "a.c") != 6.5 {
		t.Error("a.c was not 6.5")
	}

	// test coercion of float32
	if math.Floor(GetFloat64(data, "a.d")*10)/10 != 7.8 { // will be some rounding error in float32
		t.Error("a.d was not 7.8")
	}

	// test that incompatible coercion results in 0
	if GetFloat64(data, "a.e") != 0 { // will be some rounding error in float32
		t.Error("a.e was not 0")
	}

	// test that coercion on nil results in 0
	if GetFloat64(nil, "a.e") != 0 {
		t.Error("GetFloat64 on nil was not 0")
	}

	// test fallback
	if GetFloat64(data, "x", "a.c") != 6.5 {
		t.Error("fallback did not fall back to the correct value")
	}
}

func TestGetString(t *testing.T) {
	data := make(map[string]interface{})
	data["biz"] = "tammy"
	res := GetString(data, "biz")
	if res != "tammy" {
		t.Error("result did not equal tammy")
	}

	res = GetString(nil, "p")
	if res != "" {
		t.Error("GetString on nil was not the empty string")
	}

	data["obj"] = make(map[string]interface{})
	res = GetString(data, "obj")
	if res != "" {
		t.Error("result was non-empty when it should not be")
	}
}

func TestGetStringFromFloatCoercion(t *testing.T) {
	data := make(map[string]interface{})
	data["rough"] = 89.12
	res := GetString(data, "rough")
	if res != "89.12" {
		t.Error("result did not equal 89.12 as a string")
	}
}

func TestGetStringFromIntCoercion(t *testing.T) {
	data := make(map[string]interface{})
	data["rough"] = 801
	res := GetString(data, "rough")
	if res != "801" {
		t.Error("result did not equal 801 as a string")
	}
}

func TestGetStringFallback(t *testing.T) {
	data := make(map[string]interface{})
	data["b"] = "todd"
	res := GetString(data, "a", "b")
	if res != "todd" {
		t.Error("result did not equal todd - it should have fallen back")
	}
}

func TestGetStringFallbackMiss(t *testing.T) {
	data := make(map[string]interface{})
	data["b"] = "todd"
	res := GetString(data, "a", "c")
	if res != "" {
		t.Error("result did not equal '' when no match was found during fallback")
	}
}
