package Algorithm

import (
//"fmt"
)

func mergeAjust(s []int, p []int, l int, r int) bool {
	if l > r {
		return false
	}
	mid := (l + r) / 2
	l1, r1 := l, mid
	l2, r2 := mid+1, r
	i, j := l1, l2
	k := l
	for i <= r1 && j <= r2 {
		if s[i] <= s[j] {
			p[k] = s[i]
			i++
			k++
		} else {
			p[k] = s[j]
			j++
			k++
		}
	}
	for i <= r1 {
		p[k] = s[i]
		k++
		i++
	}
	for j <= r2 {
		p[k] = s[j]
		k++
		j++
	}
	for i := l; i <= r; i++ {
		s[i] = p[i]
	}
	return true
}

func MergeSort(s []int, p []int, l int, r int) bool {
	if l > r {
		return false
	}
	if l < r {
		mid := (l + r) / 2
		MergeSort(s, p, l, mid)
		MergeSort(s, p, mid+1, r)
		mergeAjust(s, p, l, r)
		return true
	}
	return true
}
