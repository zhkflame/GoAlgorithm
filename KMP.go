package GoAlgorithm

import (
	"fmt"
)

func GetNext(s string) ([]int, bool) {
	l := len(s)
	if l <= 0 {
		return nil, false
	}
	next := make([]int, l)
	k := 0
	fmt.Println(l)
	for i := 2; i < l; { //循环不需要i++,因为不是在k!=0时，只变换k,不需要改变当前位置
		if s[i-1] == s[k] { //k表示 i-1位置需要比较的位置，即k=next[i-1]
			next[i] = k + 1
			k = next[i]
			i++
		} else {
			if k != 0 {
				k = next[k]
			} else {
				next[i] = 0
				k = next[i]
				i++
			}
		}
		//fmt.Println("l: ", l)
		//fmt.Println(next)
	}
	fmt.Println("result", next)
	return next, true
}
