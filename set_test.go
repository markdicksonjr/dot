package dot

import "testing"

func TestAbs(t *testing.T) {
	obj := make(map[string]interface{})
	err := Set(obj, "Front Sprint I.D.\n", "56")
	if err == nil {
		t.Error("Did not get an error, but should have")
	}
}