package dot

import "encoding/json"

// Extend copies non-nil, non-default values from right to left
func Extend(to interface{}, from interface{}) (interface{}, error) {

	// copy "to" into "n"
	var n interface{}
	toCopyS, err := json.Marshal(to)
	if err != nil {
		return n, nil
	}
	if err := json.Unmarshal(toCopyS, &n); err != nil {
		return n, nil
	}

	keys := KeysRecursiveLeaves(from)
	for _, k := range keys {
		i, err := Get(from, k)
		if err != nil {
			return nil, err
		}

		// if a non-nil, non-default value is encountered, allow it to overwrite

		if i == nil {
			continue
		}

		iAsBool, ok := i.(bool)
		if ok && !iAsBool {
			continue
		}

		iAsString, ok := i.(string)
		if ok && iAsString == "" {
			continue
		}

		iAsCoercedInt, ok := CoerceInt64(i)
		if ok && iAsCoercedInt == 0 {
			continue
		}

		iAsCoercedFloat, ok := CoerceFloat64(i)
		if ok && iAsCoercedFloat == 0 {
			continue
		}

		if err := Set(n, k, i); err != nil {
			return nil, err
		}
	}
	return n, nil
}
