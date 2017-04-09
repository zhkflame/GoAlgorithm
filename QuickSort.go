package Algorithm

import (
//"fmt"
)

func QuickSort(s []int, l int, r int) bool {
	lenS := len(s)
	if lenS <= l || l > r {
		return false
	}
	left := l
	right := r
	temp := s[left]
	for left < right {
		for left < right && temp <= s[right] {
			right--
		}
		s[left] = s[right]
		for left < right && temp > s[left] {
			left++
		}
		s[right] = s[left]
	}
	s[left] = temp
	QuickSort(s, l, left-1)
	QuickSort(s, left+1, r)
	return true
}
