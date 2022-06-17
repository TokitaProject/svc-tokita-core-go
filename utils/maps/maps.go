package maps

func Merge(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func MergeUnique(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	mapUnique := make(map[string]bool)
	for _, m := range maps {
		for k, v := range m {
			if _, ok := mapUnique[k]; !ok {
				mapUnique[k] = true
				result[k] = v
			}
		}
	}
	return result
}
