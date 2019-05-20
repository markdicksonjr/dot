package dot

import (
	"errors"
	"github.com/oleiade/reflections"
	"reflect"
	"strings"
)

// Get will return the value in obj at the "location" given by dot notation property candidates.
// The candidates are processed in the order given, and the first non-nil result is returned.
func Get(obj interface{}, props ...string) (interface{}, error) {
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

// Loop through this to get properties via dot notation
func getProperty(obj interface{}, prop string) (interface{}, error) {

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

func Set(obj interface{}, prop string, value interface{}) error {
	// Get the array access
	arr := strings.Split(prop, ".")

	// fmt.Println(arr)
	var err error
	var key string
	var effectiveObj = obj
	last, arr := arr[len(arr)-1], arr[:len(arr)-1]
	for _, key = range arr {
		effectiveObj, err = getProperty(effectiveObj, key)
		if err != nil {
			return err
		}
	}

	if effectiveObj == nil {
		propPath := strings.Split(prop, ".")
		currentPath := ""
		for i, prop := range propPath {
			if currentPath != "" {
				currentPath += "."
			}
			currentPath += prop

			testVal, err := getProperty(obj, currentPath)
			if err != nil {
				return err
			}
			if testVal == nil && i < len(propPath)-1 {
				if err := setProperty(obj, currentPath, make(map[string]interface{})); err != nil {
					return err
				}
			}
		}

		return Set(obj, prop, value)
	}

	return setProperty(effectiveObj, last, value)
}

func setProperty(obj interface{}, prop string, val interface{}) error {
	if reflect.TypeOf(obj).Kind() == reflect.Map {

		value := reflect.ValueOf(obj)
		value.SetMapIndex(reflect.ValueOf(prop), reflect.ValueOf(val))
		return nil
	}

	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("object must be a pointer to a struct")
	}
	prop = strings.Title(prop)

	return reflections.SetField(obj, prop, val)
}
