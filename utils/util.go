package utils

// RemoveDuplicateElement 数组去重
func RemoveDuplicateElement(array []string) []string {
	result := make([]string, 0, len(array))
	temp := map[string]struct{}{}
	for _, item := range array {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
