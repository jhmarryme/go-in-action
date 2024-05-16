package main

import "fmt"

// 取变量指针的语法
// ptr := &v
// 		v:代表被取地址的变量，类型为T
// 		ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。

func main() {
	// 取变量指针
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc000012100
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc000012100 type:*int
	fmt.Println(&b)                    // 0xc00005a020
}
