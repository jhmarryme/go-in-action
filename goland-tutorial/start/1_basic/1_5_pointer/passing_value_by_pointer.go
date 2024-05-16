// 指针传值

package main

import "fmt"

// modify1 通过值传递方式修改变量值，不影响原始变量
func modify1(x int) {
	x = 100
}

// modify2 通过指针传递方式修改变量值，影响原始变量
func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)     // 通过值传递方式修改变量 a，但不影响原始变量
	fmt.Println(a) // 输出 10，原始变量值未改变
	modify2(&a)    // 通过指针传递方式修改变量 a，影响原始变量
	fmt.Println(a) // 输出 100，原始变量值已被修改
}
