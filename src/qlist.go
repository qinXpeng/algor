package algor

import "container/list"


type List[T any] struct {

	l *list.List
}

func NewList[T any]() *List[T] {
	return &List[T]{
		l:list.New(),
	}
}

