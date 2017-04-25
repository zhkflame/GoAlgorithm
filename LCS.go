package GoAlgorithm

import (
	"fmt"
)

func LCS(str1 string, str2 string) int {
	len1 := len(str1)
	len2 := len(str2)
	matrix := make([]int, len2+1)
	leftup := 0
	for i := 1; i <= len1; i++ {
		leftup = 0
		matrix[0] = 0
		for j := 1; j <= len2; j++ {
			if str1[i-1] == str2[j-1] {
				matrix[j] = leftup + 1
			} else {
				matrix[j] = getMax(matrix[j-1], matrix[j])
			}
			leftup = matrix[j]
		}
		fmt.Println(matrix)
	}
	return matrix[len2]

}

func getMax(t1 int, t2 int) int {
	temp1 := t1
	if t2 > t1 {
		temp1 = t2
	}
	return temp1
}
