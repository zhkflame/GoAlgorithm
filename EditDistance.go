package GoAlgorithm

import (
	"fmt"
)

func EditDistance(str1 string, str2 string) (int, bool) {
	len1 := len(str1)
	len2 := len(str2)
	martix := make([]int, len2+1)
	for i := 0; i < len(martix); i++ {
		martix[i] = i
	}
	for i := 1; i <= len1; i++ {
		leftup := i - 1
		martix[0] = i
		for j := 1; j <= len2; j++ {
			temp := leftup
			if str1[i-1] != str2[j-1] {
				temp = leftup + 1
			}
			leftup = martix[j]
			martix[j] = getMin(temp, martix[j-1]+1, martix[j]+1)
		}
		fmt.Println(martix)
	}
	return martix[len2], true
}

func getMin(a int, b int, c int) int {
	temp1 := a
	if b < a {
		temp1 = b
	}
	temp := temp1
	if c < temp1 {
		temp = c
	}
	return temp
}
