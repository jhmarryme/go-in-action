//  1. 求数组所有元素之和
//  2. 找出数组中和为给定值的两个元素的下标，
//     例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sumArr 返回切片所有元素之和
func sumArr(nums ...int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

// findPairs 找出切片中和为给定值的两个元素的下标
// 例如数组[1,3,5,8,7]， 找出两个元素之和等于8的下标分别是（0，4）和（1，2）
func findPairs(nums [5]int, target int) {
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func main() {
	// 使用当前时间的纳秒数作为随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成随机长度的数组
	length := r.Intn(10) + 1 // 生成 1 到 10 之间的随机整数作为数组长度
	var a []int
	for i := 0; i < length; i++ {
		a = append(a, r.Intn(1000))
	}

	fmt.Println("arr:", a)
	sum := sumArr(a...) // 计算数组元素之和
	fmt.Printf("sum=%d\n", sum)

	b := [5]int{1, 3, 5, 8, 7}
	findPairs(b, 8) // 查找和为8的元素下标对
}
