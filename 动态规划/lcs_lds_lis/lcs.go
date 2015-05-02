/*
	参考 http://www.cnblogs.com/huangxincheng/archive/2012/11/11/2764625.html
*/
package main

import "fmt"

func main() {
	a := "1234586"
	b := "145869"
	fmt.Println(findLCS(a, b))
}

func findLCS(str1, str2 string) string {

	list1 := []rune(str1)
	len1 := len(list1)

	list2 := []rune(str2)
	len2 := len(list2)

	array := make([][]int, len1+1)
	for n := 0; n < len1+1; n++ {
		array[n] = make([]int, len2+1)
	}

	for m := 1; m <= len1; m++ {
		for n := 1; n <= len2; n++ {
			// 字符相等: 增加新字符后的两个字符串的最长公共子序列长度
			// 等于先前的最长公共子序列长度,加上1
			if list1[m-1] == list2[n-1] {
				array[m][n] = array[m-1][n-1] + 1
			} else if array[m-1][n] > array[m][n-1] {
				array[m][n] = array[m-1][n]
			} else {
				array[m][n] = array[m][n-1]
			}
		}
	}

	retstr := ""
	m := len1
	n := len2

	for m > 0 && n > 0 {
		if list1[m-1] == list2[n-1] {
			retstr = string(list1[m-1]) + retstr
			m--
			n--
		} else if array[m-1][n] > array[m][n-1] {
			m--
		} else {
			n--
		}
	}

	return retstr
}
