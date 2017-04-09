package GoAlgorithm

import (
	"fmt"
	"math"
)

func Kruskal(graph [][]int) (int, bool) {
	lenG := len(graph)
	if lenG <= 0 {
		return 0, false
	}

	masgG := make([]int, lenG)
	for i := 0; i < lenG; i++ {
		masgG[i] = i
	}
	num := lenG * (lenG - 1) / 2
	edge := make([][]int, num)
	k := 0
	for i := 0; i < lenG; i++ {
		for j := i + 1; j < lenG; j++ {
			edge[k] = make([]int, 3)
			edge[k][0] = graph[i][j]
			edge[k][1] = i
			edge[k][2] = j
			k++
		}
	}
	quickSortEdge(edge, 0, num-1)
	fmt.Println(edge)
	sum := 0
	count := 0
	fmt.Println(masgG)
	for i := 0; i < num; i++ {
		x := edge[i][1]
		for x != masgG[x] {
			x = masgG[x]
		}
		y := edge[i][2]
		for y != masgG[y] {
			y = masgG[y]
		}
		if x != y && edge[i][0] != math.MaxInt32 { //判断 如果x y不同根，并且当前的两个节点是联通的
			sum = sum + edge[i][0]
			masgG[y] = x
			count = count + 1
			fmt.Println(masgG)
		}
	}
	if count < lenG-1 {
		sum = 0
		return 0, false
	}
	return sum, true
}

func quickSortEdge(edge [][]int, l int, r int) bool {
	if l > r || r >= len(edge) || edge == nil {
		return false
	}
	left := l
	right := r
	temp := edge[left]
	for left < right {
		for left < right && temp[0] <= edge[right][0] {
			right--
		}
		edge[left] = edge[right]
		for left < right && temp[0] > edge[left][0] {
			left++
		}
		edge[right] = edge[left]
	}
	edge[left] = temp
	quickSortEdge(edge, l, left-1)
	quickSortEdge(edge, left+1, r)
	return true
}
