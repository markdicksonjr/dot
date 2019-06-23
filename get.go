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
				return nil, nil
			}
		}
	}
	return obj, nil
}

// GetString does what Get does, except it continues through props until
// it not only gets a non-nil value, but also gets something that can be
// cast to a string that isn't the empty string.  Will return "" if the
// property doesn't exist or could not be coerced.
// It can't be implemented by attempting to coerce Get once complete,
// because of the fallback logic for when more than one prop is passed.
func GetString(obj interface{}, props ...string) string {
	if obj == nil {
		return ""
	}

	for _, prop := range props {

		// Get the array access
		arr := strings.Split(prop, ".")

		var err error
		for _, key := range arr {
			obj, err = getProperty(obj, key)
			if err != nil {
				return ""
			}
			if obj == nil {
				continue
			}
		}

		if obj != nil {
			asString, ok := obj.(string)
			if ok && asString != "" {
				return asString
			} else {
				obj = nil
			}
		}
	}
	return ""
}

// GetInt64 does what Get does, except it continues through props until
// it not only gets a non-nil value, but also gets something that can be
// cast/coerced to an int64 value.  Will return 0 if the property doesn't
// exist or could not be coerced.
// It can't be implemented by attempting to coerce Get once complete,
// because of the fallback logic for when more than one prop is passed.
func GetInt64(obj interface{}, props ...string) int64 {
	if obj == nil {
		return 0
	}

	for _, prop := range props {

		// Get the array access
		arr := strings.Split(prop, ".")

		var err error
		for _, key := range arr {
			obj, err = getProperty(obj, key)
			if err != nil {
				return 0
			}
			if obj == nil {
				continue
			}
		}

		if obj != nil {
			as64, ok := obj.(int64)
			if ok {
				return as64
			}

			as32, ok := obj.(int32)
			if ok {
				return int64(as32)
			}

			asInt, ok := obj.(int)
			if ok {
				return int64(asInt)
			}

			obj = nil
		}
	}
	return 0
}

// GetFloat64 does what Get does, except it continues through props until
// it not only gets a non-nil value, but also gets something that can be
// cast/coerced to a float64 value.  Will return 0 if the property doesn't
// exist or could not be coerced.
// It can't be implemented by attempting to coerce Get once complete,
// because of the fallback logic for when more than one prop is passed.
func GetFloat64(obj interface{}, props ...string) float64 {
	if obj == nil {
		return 0
	}

	for _, prop := range props {

		// Get the array access
		arr := strings.Split(prop, ".")

		var err error
		for _, key := range arr {
			obj, err = getProperty(obj, key)
			if err != nil {
				return 0
			}
			if obj == nil {
				continue
			}
		}

		if obj != nil {
			as64, ok := obj.(float64)
			if ok {
				return as64
			}

			as32, ok := obj.(float32)
			if ok {
				return float64(as32)
			}


			asInt, ok := obj.(int)
			if ok {
				return float64(asInt)
			}

			obj = nil
		}
	}
	return 0
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
