package algor


type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Len() int {
	a := *s
	return len(a)
}
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Top() any {
	a := *s
	return a[s.Len()-1]
}
func (s *Stack[T]) Push(val ...T) {
	*s=append(*s,val...)
}
func (s *Stack[T]) Pop() {
	a := *s
	*s = a[:s.Len()-1]
}
