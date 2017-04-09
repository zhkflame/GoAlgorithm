package GoAlgorithm

import (
	"fmt"
	//"strconv"
)

func MinMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	} else {
		var mq MyQueue
		mq.Init(len(bank))
		mq.Push(start)
		step := 0
		lenL := map[string]int{}
		lenL[start] = 0
		for !mq.IsEmpty() { //遍历队列
			temp := mq.Pop() //取队首元素
			if temp == end {
				return lenL[end]
			} else { //如果队首元素不是要找的
				step = lenL[temp] + 1
				for i := 0; i < len(temp); i++ { //修改temp的每一个字符，
					for j := 65; j < 91; j++ { //
						if temp[i] != byte(j) {
							tempbyte := []byte(temp)
							tempbyte[i] = byte(j)
							temp2 := string(tempbyte)
							for k := 0; k < len(bank); k++ {
								if temp2 == bank[k] {
									if _, ok := lenL[temp2]; !ok {
										mq.Push(temp2)
										lenL[temp2] = step
									}
								}
							}
						}
					}

				}
			}

		}
		return -1

	}

}
