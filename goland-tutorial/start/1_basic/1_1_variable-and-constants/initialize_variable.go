// 变量的初始化
// Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用
// 变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号
// 整型和浮点型变量的默认值为0。
// 字符串变量的默认值为空字符串。
// 布尔型变量默认为false。
// 切片、函数、指针变量的默认为nil。
package main

import "fmt"

func main() {
	// var 变量名 类型 = 表达式
	var name string = "test.com"
	var sex int = 1
	fmt.Println(name, sex)

	// 批量变量声明
	var (
		n1 string
		n2 int
		n3 bool
		n4 float32
	)
	fmt.Println(n1, n2, n3, n4)

	// 一次初始化多个变量
	var x, y = "test.com", 1
	fmt.Println(x, y)

	// 类型推导
	// 编译器会根据等号右边的值来推导变量的类型完成初始化
	var a = "test.com"
	var b = 1
	fmt.Println(a, b)
}
