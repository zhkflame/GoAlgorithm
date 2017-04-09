package Algorithm

import (
	"fmt"
)

func heapAjust(s []int, start int, end int) bool {
	if start > end {
		return false
	}
	now := start
	for now < end {
		if 2*now+1 <= end {
			max := 2*now + 1
			if 2*now+2 <= end && s[2*now+2] > s[2*now+1] {
				max = 2*now + 2
			}
			if s[now] < s[max] {
				s[now], s[max] = s[max], s[now]
				now = max
			} else {
				break
			}
		} else {
			break
		}
	}
	return true
}

func HeapSort(s []int) bool {
	lenS := len(s)
	if lenS <= 0 {
		return false
	}
	for i := (lenS - 2) / 2; i >= 0; i-- { //从最后一个有孩子的几点开始调整为大顶堆
		heapAjust(s, i, lenS-1)
	}
	fmt.Println(s)
	for i := lenS - 1; i > 0; i-- {
		s[i], s[0] = s[0], s[i]
		heapAjust(s, 0, i-1)
	}
	return true

}
