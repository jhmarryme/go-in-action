// byte和rune类型
// 组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来
// Go 语言的字符有以下两种：
// - uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
// - rune类型，代表一个 UTF-8字符。
// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型
package main

import "fmt"

func main() {
	traversalString()
}

// 遍历字符串
func traversalString() {
	s := "test.com博客"
	// byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	// rune
	for _, r := range s {
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
