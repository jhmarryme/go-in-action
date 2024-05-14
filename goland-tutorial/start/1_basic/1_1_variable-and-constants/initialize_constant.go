// 常量的初始化
// 相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。
// 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值
package main

func main() {
	// 常量在定义的时候必须赋值
	const pi = 3.1415
	const e = 2.7182

	// 多个常量也可以一起声明
	const (
		x = 3.1415
		y = 2.7182
	)
	// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	const (
		n1 = 100
		n2
		n3
	)
}