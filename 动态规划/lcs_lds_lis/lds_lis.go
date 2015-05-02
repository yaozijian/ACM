package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	cnt := rand.Intn(10) + 1
	for idx := 0; idx < cnt; idx++ {

		lens := rand.Intn(15) + 2
		data := make([]int, lens)

		for x := range data {
			data[x] = rand.Intn(100)
		}

		fmt.Println("数据:", data)
		fmt.Println("最长递减子序列:", findLongestDecreaseSub(data))
		fmt.Println("最长递增子序列:", findLongestIncreaseSub(data))
		fmt.Println("\n")
	}
}

// 找最长递减子序列
func findLongestDecreaseSub(heights []int) (retval []int) {

	cnt := len(heights)
	array := make([][]int, cnt)
	maxIdx, maxLen := -1, -1

	for idx := cnt - 1; idx >= 0; idx-- {
		// 假定从这个元素开始的最长递减子序列由自身构成
		array[idx] = []int{heights[idx]}

		tmpMaxIdx, tmpMaxLen := -1, -1
		// 对于后面的每个元素
		for x := idx + 1; x < cnt; x++ {
			// 如果当前元素大于这个元素,则当前元素可以添加到它的最长递减子序列前面
			// 并且这个元素的最长递减子序列长度更长
			if heights[idx] > heights[x] && len(array[x]) > tmpMaxLen {
				tmpMaxLen = len(array[x])
				tmpMaxIdx = x
			}
		}

		if tmpMaxIdx >= 0 {
			array[idx] = append(array[idx], array[tmpMaxIdx]...)
		}

		// 全局最长递减子序列
		if len(array[idx]) > maxLen {
			maxLen = len(array[idx])
			maxIdx = idx
		}
	}

	return array[maxIdx]
}

// 找最长递增子序列
func findLongestIncreaseSub(heights []int) (retval []int) {

	cnt := len(heights)
	array := make([][]int, cnt)
	maxIdx, maxLen := -1, -1

	for idx := 0; idx < cnt; idx++ {

		array[idx] = []int{heights[idx]}

		tmpMaxIdx, tmpMaxLen := -1, -1

		for x := idx - 1; x >= 0; x-- {
			if heights[x] < heights[idx] && len(array[x]) > tmpMaxLen {
				tmpMaxIdx = x
				tmpMaxLen = len(array[x])
			}
		}

		if tmpMaxIdx >= 0 {
			array[idx] = append(array[tmpMaxIdx], heights[idx])
		}

		if len(array[idx]) > maxLen {
			maxIdx = idx
			maxLen = len(array[idx])
		}
	}

	return array[maxIdx]
}
