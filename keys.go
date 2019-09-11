package dot

import "encoding/json"

// Keys will get the list of keys for an arbitrary structure (non-recursively).  In the result below, the result will
// be ["A", "B"], though it's best to not assume the elements are ordered.
func Keys(obj interface{}, parentPath ...string) []string {
	if obj == nil {
		return []string{}
	}

	strParentPath := ""
	if len(parentPath) > 0 {
		strParentPath = parentPath[0]
	}

	asMap, ok := obj.(map[string]interface{})
	if ok {
		var keys []string
		for k := range asMap {
			adjustedKey := k
			if len(strParentPath) > 0 {
				adjustedKey = strParentPath + "." + adjustedKey
			}
			keys = append(keys, adjustedKey)
		}
		return keys
	} else {
		// TODO: find faster way

		// json to str, str to map
		strJson, err := json.Marshal(obj)
		if err != nil {
			return []string{}
		}

		var asMap map[string]interface{}
		if err := json.Unmarshal(strJson, &asMap); err != nil {
			return []string{}
		}

		var keys []string
		for k := range asMap {
			adjustedKey := k
			if len(strParentPath) > 0 {
				adjustedKey = strParentPath + "." + adjustedKey
			}
			keys = append(keys, adjustedKey)
		}
		return keys
	}
}

// KeysRecursive is just like Keys, only recursive.  The ordering of elements in the resulting slice is not to be
// assumed at any time
func KeysRecursive(obj interface{}, parentPath ...string) []string {
	strParentPath := ""
	if len(parentPath) > 0 {
		strParentPath = parentPath[0]
	}

	var allKeys []string
	keys := Keys(obj)
	for _, k := range keys {
		adjustedChildPath := k
		if len(strParentPath) > 0 {
			adjustedChildPath = strParentPath + "." + k
		}
		allKeys = append(allKeys, adjustedChildPath)

		v, _ := Get(obj, k)
		if v != nil {
			allKeys = append(allKeys, KeysRecursive(v, adjustedChildPath)...)
		}
	}

	return allKeys
}

// KeysRecursive is like KeysRecursive, except it returns only items with no "children"
func KeysRecursiveLeaves(obj interface{}, parentPath ...string) []string {
	strParentPath := ""
	if len(parentPath) > 0 {
		strParentPath = parentPath[0]
	}

	var allKeys []string
	keys := Keys(obj)
	for _, k := range keys {
		adjustedChildPath := k
		if len(strParentPath) > 0 {
			adjustedChildPath = strParentPath + "." + k
		}

		v, _ := Get(obj, k)
		if v != nil {
			leaves := KeysRecursiveLeaves(v, adjustedChildPath)
			if len(leaves) == 0 {
				allKeys = append(allKeys, adjustedChildPath)
			} else {
				allKeys = append(allKeys, leaves...)
			}
		}
	}

	return allKeys
}
