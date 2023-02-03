package main

import (
	"fmt"

	"github.com/qinXpeng/go-Algorithm/algorithm/priority_queue"
	"github.com/qinXpeng/go-Algorithm/algorithm/list"
)

func main() {
	q := priority_queue.NewHeap(func(i,j int) bool {return i > j})
	q.Push(12)
	q.Push(1)
	q.Push(3)
	q.Push(10)
	q.Push(341)
	fmt.Println(q.Len())
	for q.Len() > 0 {
		t := q.Top()
		q.Pop()
		fmt.Println(t)
	}
	l := qlist.NewList[int]()
	l.Push_back(213)
	l.Push_front(213123)
	for l.Len() > 0 {
		fmt.Println(l.Front_value())
		l.Pop_front()
	}
}
