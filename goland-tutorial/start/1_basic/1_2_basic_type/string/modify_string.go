// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
// 无论哪种转换，都会重新分配内存，并复制字节数组
package main

import "fmt"

func main() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}
