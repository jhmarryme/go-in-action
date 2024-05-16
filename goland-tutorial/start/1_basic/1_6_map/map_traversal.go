// map的遍历
package main

import (
	"fmt"
	"sort"
)

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	// 使用for range遍历map
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 只想遍历key时
	for k := range scoreMap {
		fmt.Println(k)
	}
	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 3)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
