package GoAlgorithm

import (
	"fmt"
)

type MyQueue struct {
	queue    []string
	front    int
	size     int
	capacity int
}

func (mq *MyQueue) Init(c int) {
	mq.queue = make([]string, c)
	mq.front = 0
	mq.capacity = len(mq.queue)
	mq.size = 0
}

func (mq *MyQueue) IsEmpty() bool {
	if mq.size == 0 {
		return true
	} else {
		return false
	}
}

func (mq *MyQueue) IsFull() bool {
	if mq.size == mq.capacity {
		return true
	} else {
		return false
	}
}

func (mq *MyQueue) Push(s string) bool {
	if mq.IsFull() {
		fmt.Println("Full Now")
		return false
	} else {
		mq.queue[((mq.front + mq.size) % mq.capacity)] = s
		mq.size += 1
		return true
	}
}

func (mq *MyQueue) Pop() string {
	if mq.IsEmpty() {
		fmt.Println("Empty Now")
		return ""
	} else {
		temp := mq.front
		mq.front = (mq.front + 1) % mq.capacity
		mq.size -= 1
		return mq.queue[temp]
	}
}
