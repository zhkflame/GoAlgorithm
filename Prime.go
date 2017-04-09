package GoAlgorithm

import (
	"fmt"
	"math"
)

func Prime(graph [][]int, point int) bool {
	lenG := len(graph)
	if lenG <= 0 || point >= lenG {
		return false
	}
	distance := make([]int, lenG) //需要的辅助空间，记录到某个节点的距离，不一定是哪个节点起始
	isVisit := make([]bool, lenG) //记录节点是否访问过
	isVisit[point] = true
	for i := 0; i < lenG; i++ {
		distance[i] = graph[point][i]
	}
	sum := 0
	count := 1                  //需要记录是否遍历了所有节点，通过记录count来表示
	for i := 0; i < lenG; i++ { //目的是为了遍历所有节点
		min := math.MaxInt32
		k := math.MaxInt32
		for j := 0; j < lenG; j++ { //目的是找到当前的最小距离，所以需要记录最小是min,和最小值的位置k
			if isVisit[j] == false && distance[j] < min {
				//fmt.Println(distance[j])
				min = distance[j]
				k = j
			}
		}
		if min == math.MaxInt32 {
			if count < lenG {
				fmt.Println("不连通")
				return false
			}
			break
		}
		isVisit[k] = true
		sum = sum + min
		count++
		for i := 0; i < lenG; i++ {
			if isVisit[i] == false && graph[k][i] < distance[i] {
				distance[i] = graph[k][i]
			}
		}
	}
	fmt.Println(sum)
	return true
}
