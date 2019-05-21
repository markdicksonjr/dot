package dot

import (
	"github.com/oleiade/reflections"
	"reflect"
	"strings"
)

// Get will return the value in obj at the "location" given by dot notation property candidates.
// The candidates are processed in the order given, and the first non-nil result is returned.
func Get(obj interface{}, props ...string) (interface{}, error) {
	if obj == nil {
		return nil, nil
	}

	for _, prop := range props {

		// Get the array access
		arr := strings.Split(prop, ".")

		var err error
		for _, key := range arr {
			obj, err = getProperty(obj, key)
			if err != nil {
				return nil, err
			}
			if obj == nil {
				continue
			}
		}
	}
	return obj, nil
}

// GetString does what Get does, except it continues through props until
// it not only gets a non-nil value, but also gets something that can be
// cast to a string that isn't the empty string.  Will return "" if the
// property doesn't exist.
func GetString(obj interface{}, props ...string) (string, error) {
	if obj == nil {
		return "", nil
	}

	for _, prop := range props {

		// Get the array access
		arr := strings.Split(prop, ".")

		var err error
		for _, key := range arr {
			obj, err = getProperty(obj, key)
			if err != nil {
				return "", err
			}
			if obj == nil {
				continue
			}
		}

		if obj != nil {
			asString, ok := obj.(string)
			if ok && asString != "" {
				return asString, nil
			} else {
				obj = nil
			}
		}
	}
	return "", nil
}

// Loop through this to get properties via dot notation
func getProperty(obj interface{}, prop string) (interface{}, error) {
	if obj == nil {
		return nil, nil
	}

	if reflect.TypeOf(obj).Kind() == reflect.Map {

		val := reflect.ValueOf(obj)

		valueOf := val.MapIndex(reflect.ValueOf(prop))

		if valueOf == reflect.Zero(reflect.ValueOf(prop).Type()) {
			return nil, nil
		}

		idx := val.MapIndex(reflect.ValueOf(prop))

		if !idx.IsValid() {
			return nil, nil
		}
		return idx.Interface(), nil
	}

	kind := reflect.TypeOf(obj).Kind()
	if kind == reflect.Slice {
		return obj, nil
	}

	prop = strings.Title(prop)
	return reflections.GetField(obj, prop)
}
