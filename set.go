package dot

import (
	"errors"
	"github.com/oleiade/reflections"
	"reflect"
	"strings"
)

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
