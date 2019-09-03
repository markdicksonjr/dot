package dot

import (
	"errors"
	"github.com/oleiade/reflections"
	"reflect"
	"strings"
)

func Set(obj interface{}, prop string, value interface{}) error {
	if obj == nil {
		return errors.New("obj may not be nil for dot.Set")
	}

	// trim outer spaces from property
	prop = strings.TrimSpace(prop)

	//if len(prop) == 0 {
	//	return errors.New("zero-length properties are not allowed")
	//}

	if prop[0] == '.' {
		return errors.New("dot-set property may not start with '.'")
	}

	if prop[len(prop) - 1] == '.' {
		return errors.New("dot-set property may not end in '.'")
	}

	// get the array access
	arr := strings.Split(prop, ".")

	// get each level of property, all the way down to the leaf
	var err error
	var key string
	var effectiveObj = obj
	last, arr := arr[len(arr)-1], arr[:len(arr)-1]
	for _, key = range arr {
		effectiveObj, err = getProperty(effectiveObj, key, true)
		if err != nil {
			break
		}
	}

	// if we need to allocate all the way down to the object
	if effectiveObj == nil {
		propPath := strings.Split(prop, ".")

		// if we're at the end of props
		if len(propPath) == 1 {
			if err := setProperty(obj, propPath[0], value); err != nil {
				return err
			}
			return nil
		}

		if err := setProperty(obj, propPath[0], make(map[string]interface{})); err != nil {
			return err
		}

		innerProp, err := getProperty(obj, propPath[0], true)
		if err != nil {
			return err
		}

		return Set(innerProp, strings.Join(propPath[1:], "."), value)
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
