package Algorithm

import (
	"fmt"
)

func LIS(s []int) {
	l := len(s)
	matrix := make([]int, l, 2*l)
	for i := 0; i < l; i++ {
		matrix[i] = 0
	}
	//max := 1
	matrix[0] = s[0]
	lm := 1 //表示matrix的长度
	for i := 1; i < l; i++ {
		left := 0
		right := lm - 1
		for left < right {
			mid := (left + right) / 2
			//fmt.Println(mid)
			if matrix[mid] > s[i] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		fmt.Println(matrix, "     ", s[i])
		if matrix[left] <= s[i] {
			left = left + 1
			lm = lm + 1
		}
		matrix[left] = s[i]
	}
	fmt.Println(matrix[lm-1])
	fmt.Println(matrix)
}
