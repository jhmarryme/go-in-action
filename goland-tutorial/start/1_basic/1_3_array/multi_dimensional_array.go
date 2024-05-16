package main

import (
	"fmt"
)

// 全局多维数组
var (
	// 未初始化的 5x3 整数数组，所有元素默认值为 0
	arr3 [5][3]int
	// 使用初始化值的 2x3 整数数组
	arr4 = [...][3]int{{1, 2, 3}, {7, 8, 9}}
)

func main() {
	// 局部多维数组
	// 初始化 2x3 整数数组
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// 初始化 3x2 整数数组，第 2 维度不能使用 "..."
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}}

	// 1. 打印全局数组
	fmt.Println("1. Global Arrays:")
	fmt.Println("arr3:", arr3)
	fmt.Println("arr4:", arr4)

	// 2. 打印局部数组
	fmt.Println("\n2. Local Arrays:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// 3. 示例：值拷贝行为
	// 数组是值类型，函数参数传递时会复制整个数组
	fmt.Println("\n3. Value Copy Behavior:")
	fmt.Printf("Address of a before test: %p\n", &a)
	test(a) // 调用 test 函数时，传递的是 a 的副本
	fmt.Println("a after test:", a)

	// 4. 示例：内置函数 len 和 cap
	fmt.Println("\n4. len and cap functions:")
	fmt.Printf("len(a): %d, cap(a): %d\n", len(a), cap(a)) // 对于数组 len 和 cap 返回相同的值

	// 5. 示例：多维数组遍历
	fmt.Println("\n5. Multi-dimensional Array Traversal:")
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}
}

// 函数：接收一个 2x3 元素数组，打印数组地址并修改其副本的值
func test(x [2][3]int) {
	fmt.Printf("Address of x in test: %p\n", &x)
	x[1][1] = 1000 // 修改副本，不影响原数组
}
