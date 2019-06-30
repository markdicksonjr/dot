package dot

import (
	"math"
	"testing"
)

func TestGetInt64(t *testing.T) {
	data := make(map[string]interface{})
	a := make(map[string]interface{})
	a["b"] = 5
	a["c"] = 6.5
	a["d"] = int32(7)
	a["e"] = "should be 0"
	data["a"] = a

	// test coercion of int
	if GetInt64(data, "a.b") != 5 {
		t.Error("a.b was not 5")
	}

	// test simple get of target type
	if GetInt64(data, "a.c") != 0 {
		t.Error("a.c was not 0")
	}

	// test coercion of float32
	if GetInt64(data, "a.d") != 7 { // will be some rounding error in float32
		t.Error("a.d was not 7")
	}

	// test that incompatible coercion results in 0
	if GetInt64(data, "a.e") != 0 {
		t.Error("a.e was not 0")
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
	if math.Floor(GetFloat64(data, "a.d") * 10) / 10 != 7.8 { // will be some rounding error in float32
		t.Error("a.d was not 7.8")
	}

	// test that incompatible coercion results in 0
	if GetFloat64(data, "a.e") != 0 { // will be some rounding error in float32
		t.Error("a.e was not 0")
	}
}

func TestGetString(t *testing.T) {
	data := make(map[string]interface{})
	data["biz"] = "tammy"
	res := GetString(data, "biz")
	if res != "tammy" {
		t.Error("result did not equal tammy")
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