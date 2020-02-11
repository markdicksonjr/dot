package dot

import "testing"

type InnermostStruct struct {
	G map[string]interface{}
}

type InnerStruct struct {
	X string
	Y *string
	Z InnermostStruct
}

type SampleStruct struct {
	A float64
	B int
	C string
	D InnerStruct
}

func TestTopLevelSet(t *testing.T) {

	// first, test using a map
	obj := make(map[string]interface{})
	err := Set(obj, "X", "test34")
	if err != nil {
		t.Fatal("Got an error setting a root map prop, error = " + err.Error())
	}
	if obj["X"].(string) != "test34" {
		t.Fatal("X != test34")
	}
	err = Set(nil, "a", "t")
	if err == nil {
		t.Fatal("Did not get an error when setting on a nil object")
	}

	// now, test using structs
	s := SampleStruct{}

	// test a non-pointer struct usage
	err = Set(s, "A", 6.7)
	if err == nil {
		t.Fatal("Did not get an error when set got a non-pointer struct")
	}

	// test a simple prop set
	err = Set(&s, "A", 6.7)
	if err != nil {
		t.Fatal("Got an error when getting struct root value")
	}

	// test a simple prop set
	err = Set(&s, ".", 4)
	if err == nil {
		t.Fatal("Did not get error when one should have been returned")
	}

	err = Set(&s, "D.X", "t")
	if err != nil {
		t.Fatal("Got an error when setting non-pointer sub-struct")
	}

	f, _ := Get(s, "D.X")
	if f != "t" {
		t.Fatal("Did not get back what was set on nested struct")
	}
}


func TestSet_EscapedDotInMap(t *testing.T) {
	tMap := map[string]interface{} {}
	err := Set(tMap, "t\\.b", "r")
	if err != nil {
		t.Fatal(err)
	}
	err = Set(tMap, "t.a", "r")
	if err != nil {
		t.Fatal(err)
	}
	if tMap["t.a"] != nil {
		t.Fatal("unexpected nil result")
	}
	if tMap["t.b"] == nil {
		t.Fatal("unexpected nil result")
	}

	s, _ := tMap["t.b"].(string)
	if s != "r" {
		t.Fatal("value not correct for escape dot set")
	}
}

func TestSet_NestedCreateRequired(t *testing.T) {
	s := SampleStruct{}
	err := Set(&s, "D.Z.G.P", map[string]interface{}{
		"h": 6,
	})
	if err != nil {
		t.Fatal("Got an error when setting non-pointer sub-struct")
	}
}

func TestSet_BlankKey(t *testing.T) {

	obj := make(map[string]interface{})
	err := Set(obj, "", "foo")
	if err != nil {
		t.Fatal("Got an error = " + err.Error())
	}
	if obj[""] == nil {
		t.Fatal("could not get blank property from obj")
	}
	if obj[""] != "foo" {
		t.Fatal("failed to retrieve blank property after set")
	}
}

func TestSimpleTwoLevelSet(t *testing.T) {
	obj := make(map[string]interface{})
	err := Set(obj, "F.A", "hoo")
	if err != nil {
		t.Fatal("Got an error = " + err.Error())
	}
	f := obj["F"]
	if f == nil {
		t.Fatal("F not found")
	}
	fMap := f.(map[string]interface{})
	a := fMap["A"]
	if a.(string) != "hoo" {
		t.Fatal("F.A != hoo")
	}
}

func TestEndsWithPeriod(t *testing.T) {
	obj := make(map[string]interface{})
	err := Set(obj, "ABC.", "yes")
	if err == nil {
		t.Error("Did not get an error, but should have")
	}
}

func TestBadCharsNoPeriod(t *testing.T) {
	obj := make(map[string]interface{})
	err := Set(obj, "Front Sprint I.D.\n", "56")
	if err == nil {
		t.Error("Did not get an error, but should have")
	}
}

func TestBadCharsNoPeriodAtEnd(t *testing.T) {
	obj := make(map[string]interface{})
	err := Set(obj, "Front SPRING I.D.\n(mm)", "MLL")
	if err != nil {
		t.Fatal("Got an error = " + err.Error())
	}

	top := obj["Front SPRING I"]
	if top == nil {
		t.Fatal("Front SPRINT I not set")
	}

	topMap := top.(map[string]interface{})
	d := topMap["D"]
	if d == nil {
		t.Fatal("Front SPRINT I.D not set")
	}

	dMap := d.(map[string]interface{})
	mm := dMap["(mm)"]
	if mm == nil {
		t.Fatal("Front SPRINT I.D (mm) not set")
	}

	if mm.(string) != "MLL" {
		t.Fatal("mm != MLL")
	}
}
