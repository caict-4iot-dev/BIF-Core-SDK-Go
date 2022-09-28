package utils

import "strings"

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

func charToByte(c byte) byte {
	return byte(strings.IndexByte("0123456789ABCDEF", c))
}

func HexStringToBytes(hexString string) []byte {
	if hexString == "" || len(hexString)%2 != 0 {
		return nil
	}

	hexString = strings.ToUpper(hexString)
	length := len(hexString) / 2

	d := make([]byte, length)

	for i := 0; i < length; i++ {
		pos := i * 2
		d[i] = charToByte(hexString[pos])<<4 | charToByte(hexString[pos+1])
	}

	return d
}
