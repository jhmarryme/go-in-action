// 空指针的判断
// 当一个指针被定义后没有分配到任何变量时，它的值为 nil
package main

import "fmt"

func main() {
	var p *string                // 声明一个指向字符串的指针变量 p，但未初始化，其默认值为 nil
	fmt.Println("p 的值:", p)      // 打印指针 p 的值，由于 p 是空指针，输出结果为 nil
	fmt.Printf("p 的地址是：%p\n", p) // 打印指针 p 的地址，%p 用于格式化输出指针的地址
	if p != nil {                // 判断指针 p 是否为 nil
		fmt.Println("p 指向的内存地址非空")
	} else {
		fmt.Println("p 是空指针") // 如果指针 p 是空指针，输出 "p 是空指针"，否则输出 "p 指向的内存地址非空"
	}
}
