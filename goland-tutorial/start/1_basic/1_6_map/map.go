// map定义
// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

// Go语言中 map的定义语法如下
//
//	   map[KeyType]ValueType
//			其中，
//	   		KeyType:表示键的类型。
//	   		ValueType:表示键对应的值的类型。
//
// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
//
//	   make(map[KeyType]ValueType, [cap])
//			其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo)
}
