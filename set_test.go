package dot

import "testing"

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

	// now, test using a struct
	type SampleStruct struct {
		A float64
		B int
		C string
	}
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
