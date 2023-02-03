package priority_queue

type Heap[T comparable] struct {
	arr     []T
	compare func(i, j T) bool
	n       int
}

func NewHeap[T comparable](less func(i, j T) bool) *Heap[T] {
	return &Heap[T]{
		arr:     []T{},
		compare: less,
		n:       0,
	}
}
func (p Heap[T]) Less(i, j int) bool {
	return p.compare(p.arr[i], p.arr[j])
}
func (p Heap[T]) Swap(i, j int) {
	p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
}
func (p Heap[T]) Len() int {
	return p.n
}
func (p *Heap[T]) addSize(x int) {
	p.n += x
}

func (p *Heap[T]) Init() {
	n := p.Len()
	for i := n/2 - 1; i >= 0; i-- {
		p.down(i, n)
	}
}

func (p *Heap[T]) Push(x ...T) {
	for _, v := range x {
		p.arr = append(p.arr, v)
		p.addSize(1)
		p.up(p.Len() - 1)
	}
}

func (p *Heap[T]) Top() T {
	return p.arr[0]
}

func (p *Heap[T]) Pop() {
	p.addSize(-1)
	p.Swap(0, p.n)
	p.down(0, p.n)
}

func (p *Heap[T]) Remove(i int) {
	n := p.Len() - 1
	if n != i {
		p.Swap(i, n)
		if !p.down(i, n) {
			p.up(i)
		}
	}
}

func (p *Heap[T]) Fix(i int) {
	if !p.down(i, p.Len()) {
		p.up(i)
	}
}

func (p *Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !p.Less(j, i) {
			break
		}
		p.Swap(i, j)
		j = i
	}
}

func (p *Heap[T]) down(i0, n int) bool {
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
