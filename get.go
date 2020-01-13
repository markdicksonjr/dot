package dot

import (
	"errors"
	"github.com/oleiade/reflections"
	"reflect"
	"strconv"
	"strings"
)

// Get will return the value in obj at the "location" given by dot notation property candidates.
// The candidates are processed in the order given, and the first non-nil result is returned.
// If a property
func Get(obj interface{}, props ...string) (interface{}, error) {
	if obj == nil {
		return nil, nil
	}

	// allow fallback to other properties if props earlier in the list
	// have errors (probably because they don't exist)
	var lastError error
	var err error

	// loop through each property option
	for _, prop := range props {

		// initialize a cursor for the current descendent of obj
		objCursor := obj

		// continue to follow the dot-path, using the cursor
		for _, key := range strings.Split(prop, ".") {

			// get the value one level down from the objCursor
			if objCursor, err = getProperty(objCursor, key); err != nil {

				// if we can't follow the path, mark it as the most recent error and move to the next property option
				lastError = err
				break
			}
		}

		// if we ended up picking a non-nil leaf, return it (don't process more options)
		// note that lastError is no longer applicable, as we found a valid fallback
		if objCursor != nil {
			return objCursor, nil
		}
	}

	return nil, lastError
}

// GetString does what Get does, except it continues through props until
// it not only gets a non-nil value, but also gets something that can be
// cast or coerced to a string that isn't the empty string.  Will return
// "" if the property doesn't exist or could not be coerced.
// It can't be implemented by attempting to coerce Get once complete,
// because of the fallback logic for when more than one prop is passed.
func GetString(obj interface{}, props ...string) string {
	if obj == nil {
		return ""
	}

	var err error

	for _, prop := range props {
		objCursor := obj

		for _, key := range strings.Split(prop, ".") {
			objCursor, err = getProperty(objCursor, key)
			if err != nil {
				break
			}
		}

		if objCursor != nil {
			asString, ok := CoerceString(objCursor)
			if ok {
				return asString
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

		objCursor := obj

		var err error
		for _, key := range arr {
			objCursor, err = getProperty(objCursor, key)
			if err != nil {
				break
			}
		}

		if objCursor != nil {
			as64, ok := CoerceInt64(objCursor)
			if ok {
				return as64
			}
		}
	}
	return 0
}

// CoerceInt64 will make a best-effort to convert the provided argument to an int64.  It supports int64, int32, int,
// float64, float32 as acceptable inputs, but expect this list to expand further with new releases
func CoerceInt64(obj interface{}) (int64, bool) {
	as64, ok := obj.(int64)
	if ok {
		return as64, true
	}

	as32, ok := obj.(int32)
	if ok {
		return int64(as32), true
	}

	asInt, ok := obj.(int)
	if ok {
		return int64(asInt), true
	}

	fromFloat64, ok := obj.(float64)
	if ok {
		return int64(fromFloat64), true
	}

	fromFloat32, ok := obj.(float32)
	if ok {
		return int64(fromFloat32), true
	}

	return 0, false
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

		objCursor := obj

		var err error
		for _, key := range arr {
			objCursor, err = getProperty(objCursor, key)
			if err != nil {
				break
			}
		}

		if objCursor != nil {
			as64, ok := CoerceFloat64(objCursor)
			if ok {
				return as64
			}
		}
	}
	return 0
}

// CoerceFloat64 will make a best-effort to convert the provided argument to an float64.  It supports int64, int32, int,
// float64, float32 as acceptable inputs, but expect this list to expand further with new releases
func CoerceFloat64(obj interface{}) (float64, bool) {
	as64, ok := obj.(float64)
	if ok {
		return as64, true
	}

	as32, ok := obj.(float32)
	if ok {
		return float64(as32), true
	}

	asInt, ok := obj.(int)
	if ok {
		return float64(asInt), true
	}

	asInt64, ok := obj.(int64)
	if ok {
		return float64(asInt64), true
	}

	asInt32, ok := obj.(int32)
	if ok {
		return float64(asInt32), true
	}

	asString, ok := obj.(string)
	if ok {
		floatVal, err := strconv.ParseFloat(asString, 64)
		if err == nil {
			return floatVal, true
		}
	}

	asBytes, ok := obj.([]byte)
	if ok {
		floatVal, err := strconv.ParseFloat(string(asBytes), 64)
		if err == nil {
			return floatVal, true
		}
	}

	return 0, false
}

// CoerceString will make a best-effort to convert the provided argument to a string.  It supports string as well as
// anything supported by CoerceFloat64 and CoerceInt64.  Expect the list of supported argument types to expand.
func CoerceString(objCursor interface{}) (string, bool) {
	asString, ok := objCursor.(string)
	if ok && asString != "" {
		return asString, true
	}

	asFloat64, ok := CoerceFloat64(objCursor)
	if ok {
		return strconv.FormatFloat(asFloat64, 'f', -1, 64), true
	}

	asInt64, ok := CoerceInt64(objCursor)
	if ok {
		return strconv.Itoa(int(asInt64)), true
	}

	asBool, ok := objCursor.(bool)
	if ok {
		return strconv.FormatBool(asBool), true
	}

	return "", false
}

// Loop through this to get properties via dot notation
func getProperty(obj interface{}, prop string) (interface{}, error) {
	if obj == nil {
		return nil, nil
	}

	// try to get the value without further use of reflections (only works if obj is castable to map[string]interface{})
	// while the reflections version works for map[string](ANY)
	asMap, ok := obj.(map[string]interface{})
	if ok {
		return asMap[prop], nil
	}

	kind := reflect.TypeOf(obj).Kind()

	if kind == reflect.Slice {
		return obj, nil // TODO: this kind of seems funny - probably should be nil, nil
	} else if kind == reflect.Map {

		// the inbound object is a map, but not map[string]interface{}, use reflections to get the value
		val := reflect.ValueOf(obj)
		valueOf := val.MapIndex(reflect.ValueOf(prop))

		// check for nil (zero of value type)
		if valueOf == reflect.Zero(reflect.ValueOf(prop).Type()) {
			return nil, nil
		}

		// index into the map to get the property's value
		idx := val.MapIndex(reflect.ValueOf(prop))
		if !idx.IsValid() {
			return nil, errors.New("property " + prop + " not found")
		}
		return idx.Interface(), nil
	}

	return reflections.GetField(obj, strings.Title(prop))
}
