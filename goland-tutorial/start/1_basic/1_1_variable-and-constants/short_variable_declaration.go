// 短变量声明
// 在函数内部，可以使用更简略的 := 方式声明并初始化变量
package main

import (
	"fmt"
)

func main() {
	//  :=不能使用在函数外。
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)
}
