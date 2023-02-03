package stack


type Stack []any

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	a := *s
	return len(a)
}

func (s *Stack) Top() any {
	a := *s
	return a[s.Len()-1]
}
func (s *Stack) Push(val ...any) {
	*s=append(*s,val...)
}
func (s *Stack) Pop() {
	a := *s
	*s = a[:s.Len()-1]
}
