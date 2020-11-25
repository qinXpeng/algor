package stack

type Value interface {
}

type Stack struct {
	n   int
	arr []Value
}

func NewSatck() *Stack {
	return &Stack{
		arr: []Value{},
		n:   0,
	}
}

func (s *Stack) Len() int {
	return s.n
}
func (s *Stack) addSize(x int) {
	s.n += x
}
func (s *Stack) Push(val Value) {
	s.arr = append(s.arr, val)
	s.addSize(1)
}
func (s *Stack) Top() Value {
	return s.arr[s.n-1]
}
func (s *Stack) Pop() {
	s.addSize(-1)
	s.arr = s.arr[:s.n-1]
}
