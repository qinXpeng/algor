package hep

type Value interface {
}
type LessFunc func(i, j Value) bool
type Heap struct {
	arr     []Value
	compare LessFunc
	n       int
}

func NewHeap(less LessFunc) *Heap {
	return &Heap{
		arr:     []Value{},
		compare: less,
		n:       0,
	}
}
func (p Heap) Less(i, j int) bool {
	return p.compare(p.arr[i], p.arr[j])
}
func (p Heap) Swap(i, j int) {
	p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
}
func (p Heap) Len() int {
	return p.n
}
func (p *Heap) addSize(x int) {
	p.n += x
}

func (p *Heap) Init() {
	n := p.Len()
	for i := n/2 - 1; i >= 0; i-- {
		p.down(i, n)
	}
}

func (p *Heap) Push(x ...Value) {
	for _, v := range x {
		p.arr = append(p.arr, v)
		p.addSize(1)
		p.up(p.Len() - 1)
	}
}

func (p *Heap) Top() Value {
	return p.arr[0]
}

func (p *Heap) Pop() {
	p.addSize(-1)
	p.Swap(0, p.n)
	p.down(0, p.n)
}

func (p *Heap) Remove(i int) {
	n := p.Len() - 1
	if n != i {
		p.Swap(i, n)
		if !p.down(i, n) {
			p.up(i)
		}
	}
}

func (p *Heap) Fix(i int) {
	if !p.down(i, p.Len()) {
		p.up(i)
	}
}

func (p *Heap) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !p.Less(j, i) {
			break
		}
		p.Swap(i, j)
		j = i
	}
}

func (p *Heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && p.Less(j2, j1) {
			j = j2
		}
		if !p.Less(j, i) {
			break
		}
		p.Swap(i, j)
		i = j
	}
	return i > i0
}
