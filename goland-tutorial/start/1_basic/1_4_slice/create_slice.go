// 创建切片的各种方式
package main

import "fmt"

func main() {
	// 1. 声明一个空切片s1
	var s1 []int

	// 检查s1是否为nil
	if s1 == nil {
		fmt.Println("s1 是一个空切片")
	} else {
		fmt.Println("s1 不是一个空切片")
	}

	// 2. 使用 := 运算符声明并初始化一个空切片s2
	s2 := []int{}

	// 3. 使用 make() 函数创建一个空切片s3
	var s3 []int = make([]int, 0)

	// 打印三个切片
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)

	// 4. 初始化并赋值一个空切片s4
	var s4 []int = make([]int, 0, 0)
	fmt.Println("s4:", s4)

	// 5. 初始化并赋值一个包含元素的切片s5
	s5 := []int{1, 2, 3}
	fmt.Println("s5:", s5)

	// 6. 从数组中切片创建一个新切片s6
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int

	// 切片arr，包含索引1到索引3的元素，不包括索引3

	s6 = arr[1:4]
	fmt.Println("s6:", s6)
}
