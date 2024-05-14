// 字符串的常用操作
// --------------
// len(str)	求长度
// +或fmt.Sprintf	拼接字符串
// strings.Split	分割
// strings.Contains	判断是否包含
// strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
// strings.Index(),strings.LastIndex()	子串出现的位置
// strings.Join(a[]string, sep string)	join操作
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, world!"

	// 求字符串长度
	length := len(str)
	fmt.Println("Length of the string:", length)

	// 拼接字符串
	strConcat := "Hello" + ", " + "world!"
	fmt.Println("Concatenated string:", strConcat)

	// 分割字符串
	parts := strings.Split(str, ",")
	fmt.Println("String parts after split:", parts)

	// 判断是否包含某子串
	contains := strings.Contains(str, "world")
	fmt.Println("Does string contain 'world'?", contains)

	// 判断是否以指定前缀开始
	hasPrefix := strings.HasPrefix(str, "Hello")
	fmt.Println("Does string start with 'Hello'?", hasPrefix)

	// 判断是否以指定后缀结尾
	hasSuffix := strings.HasSuffix(str, "world!")
	fmt.Println("Does string end with 'world!'?", hasSuffix)

	// 获取子串第一次出现的位置
	index := strings.Index(str, "world")
	fmt.Println("Index of 'world':", index)

	// 获取子串最后一次出现的位置
	lastIndex := strings.LastIndex(str, "o")
	fmt.Println("Last index of 'o':", lastIndex)

	// join操作
	strList := []string{"Hello", "world"}
	joinedStr := strings.Join(strList, ", ")
	fmt.Println("Joined string:", joinedStr)
}
