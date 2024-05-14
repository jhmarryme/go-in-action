// iota
// iota是go语言的常量计数器，只能在常量的表达式中使用
// iota在const关键字出现时将被重置为0
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
// 使用iota能简化定义，在定义枚举时很有用

package main

func main() {
	const (
		n1 = iota
		n2
		n3
		n4
	)

	// 使用_跳过某些值
	const (
		x1 = iota
		x2
		_
		x4
	)

	// iota声明中间插队
	const (
		y1 = iota
		y2 = 100
		y3 = iota
		y4
	)
	const y5 = iota

	// 定义数量级
	// 这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位
	// 也就是由1变成了10000000000，也就是十进制的1024。
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	// 多个iota定义在一行
	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
}
