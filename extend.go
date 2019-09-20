package dot

// Extend copies non-nil, non-default values from right to left
func Extend(to interface{}, from interface{}) error {

	keys := KeysRecursiveLeaves(from)
	for _, k := range keys {
		i, err := Get(from, k)
		if err != nil {
			return err
		}

		// if a non-nil, non-default value is encountered, allow it to overwrite

		if i == nil {
			continue
		}

		// though we check nil above, also check typed nil against pointer interface-cast nil (a Go gotcha)
		if iAsInterface, ok := i.(*interface{}); ok && iAsInterface == nil {
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

		if err := Set(to, k, i); err != nil {
			return err
		}
	}
	return nil
}
