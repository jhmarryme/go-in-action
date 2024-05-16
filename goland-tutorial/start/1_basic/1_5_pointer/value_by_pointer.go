// 指针取值
// 在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针取值
package main

import "fmt"

func main() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b                          // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)  // %T：表示打印变量的类型。
	fmt.Printf("value of c:%v\n", c) // %v：表示打印变量的值。
	fmt.Println("value of c:", c)    // 直接打印值
}
