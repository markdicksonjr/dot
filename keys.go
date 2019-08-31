package dot

import "encoding/json"

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
