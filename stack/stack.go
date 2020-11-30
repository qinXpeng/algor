package stack

type Value interface {
}

type Stack []Value

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	a := *s
	return len(a)
}

func (s *Stack) Top() Value {
	a := *s
	return a[s.Len()-1]
}
func (s *Stack) Push(val ...Value) {
	for _, v := range val {
		*s = append(*s, v)
	}
}
func (s *Stack) Pop() {
	a := *s
	*s = a[:s.Len()-1]
}
