package main

import (
	"fmt"
)

// 1. 数组：是同一种数据类型的固定长度的序列。
// 2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
// 3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
// 4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
// for i := 0; i < len(a); i++ {
// }
// for index, v := range a {
// }
// 5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
// 6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
// 7. 支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
// 8. 指针数组 [n]*T，数组指针 *[n]T。

// 全局变量
var (
	arr0 = [5]int{1, 2, 3}                       // 未初始化元素值为 0
	arr1 = [5]int{1, 2, 3, 4, 5}                 // 初始化全部元素
	arr2 = [...]int{1, 2, 3, 4, 5, 6}            // 使用初始化值确定数组长度
	str  = [5]string{3: "hello world", 4: "tom"} // 使用索引初始化元素
)

func main() {
	// 局部变量
	a := [3]int{1, 2}           // 未初始化元素值为 0
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度
	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型
		{"user2", 20}, // 别忘了最后一行的逗号
	}

	// 打印全局变量
	fmt.Println("Global Variables:")
	fmt.Println("arr0:", arr0)
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("str:", str)

	// 打印局部变量
	fmt.Println("\nLocal Variables:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	// 4. 数组的访问
	fmt.Println("\nAccessing Array Elements:")
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	for index, v := range b {
		fmt.Printf("b[%d] = %d\n", index, v)
	}

	// 5. 访问越界示例（此行代码会引发 panic，因此注释掉）
	// fmt.Println(a[3])

	// 6. 数组是值类型
	fmt.Println("\nArray as Value Type:")
	e := a
	e[0] = 10
	fmt.Println("a after modifying e:", a) // a remains unchanged
	fmt.Println("e:", e)

	// 7. 支持 "=="、"!=" 操作符
	fmt.Println("\nArray Comparison:")
	f := [3]int{1, 2, 0}
	fmt.Println("a == f:", a == f)
	fmt.Println("a != f:", a != f)

	// 8. 指针数组和数组指针
	fmt.Println("\nPointer Array and Array Pointer:")
	var ptrArr [3]*int
	for i := 0; i < len(a); i++ {
		ptrArr[i] = &a[i]
	}
	fmt.Println("Pointer array:", ptrArr)

	var arrPtr *[3]int = &a
	fmt.Println("Array pointer:", *arrPtr)
}
