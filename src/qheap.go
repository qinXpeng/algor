package algor

type Heap[T comparable] struct {
	queue   []T
	compare func(i, j T) bool
}

func NewHeap[T comparable](less func(i, j T) bool) *Heap[T] {
	return &Heap[T]{
		queue:   []T{},
		compare: less,
	}
}
func (p Heap[T]) Less(i, j int) bool {
	return p.compare(p.queue[i], p.queue[j])
}
func (p Heap[T]) Swap(i, j int) {
	p.queue[i], p.queue[j] = p.queue[j], p.queue[i]
}
func (p Heap[T]) Len() int {
	return len(p.queue)
}

func (p Heap[T]) Empty() bool {
	return p.Len() == 0
}

func (p *Heap[T]) Init() {
	n := p.Len()
	for i := n/2 - 1; i >= 0; i-- {
		p.down(i, n)
	}
}

func (p *Heap[T]) Push(x ...T) {
	for _, v := range x {
		p.queue = append(p.queue, v)
		p.up(p.Len() - 1)
	}
}

func (p *Heap[T]) Top() T {
	return p.queue[0]
}

func (p *Heap[T]) Pop() {
	n := p.Len()
	p.Swap(0, n-1)
	p.down(0, n-1)
	p.queue = p.queue[:n-1]
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
