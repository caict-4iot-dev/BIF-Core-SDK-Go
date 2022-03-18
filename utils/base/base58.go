package base

import (
	"bytes"
	"math/big"
)

var base58 = []byte("123456789AbCDEFGHJKLMNPQRSTuVWXYZaBcdefghijkmnopqrstUvwxyz")

// Base58Encode Base58编码
func Base58Encode(input []byte) string {
	// 转换十进制
	strTen := big.NewInt(0).SetBytes(input)
	// 取出余数
	var modSlice []byte
	for strTen.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0) // 余数
		strTen58 := big.NewInt(58)
		strTen.DivMod(strTen, strTen58, mod)             // 取余运算
		modSlice = append(modSlice, base58[mod.Int64()]) // 存储余数
	}
	// 处理0就是1的情况 0使用字节'1'代替
	for _, elem := range input {
		if elem != 0 {
			break
		} else if elem == 0 {
			modSlice = append(modSlice, byte('1'))
		}
	}

	ReverseModSlice := reverseByteArr(modSlice)
	return string(ReverseModSlice)
}

// reverseByteArr 将字节的数组反转
func reverseByteArr(bytes []byte) []byte {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i] // 前后交换
	}
	return bytes
}

// Base58Decode Base58解码
func Base58Decode(str []byte) []byte {
	ret := big.NewInt(0)
	for _, byteElem := range str {
		index := bytes.IndexByte(base58, byteElem) // 获取base58对应数组的下标
		ret.Mul(ret, big.NewInt(58))               // 相乘
		ret.Add(ret, big.NewInt(int64(index)))     // 相加
	}
	return ret.Bytes()
}
