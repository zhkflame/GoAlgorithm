package Algorithm

import (
	"fmt"
)

func GetActive(act [][]int) (int, bool) {
	lenA := len(act)
	if lenA <= 0 {
		return 0, false
	}
	quickSortActive(act, 0, lenA-1)
	count := 0
	start := 0
	for i := 0; i < lenA; i++ {
		if act[i][0] > start {
			count++
			start = act[i][1]
			fmt.Println(i, start)
		}
	}
	fmt.Println(count)
	return count, true

}

func quickSortActive(act [][]int, left int, right int) bool {
	//lenA := len(act)
	if left >= right {
		return false
	}
	l := left
	r := right
	temp := act[l]
	for l < r {
		for l < r && temp[1] <= act[r][1] {
			r--
		}
		act[l] = act[r]
		for l < r && temp[1] > act[l][1] {
			l++
		}
		act[r] = act[l]
	}
	act[l] = temp
	quickSortActive(act, left, l-1)
	quickSortActive(act, l+1, right)
	return true
}
