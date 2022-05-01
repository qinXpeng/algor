package list


type ListNode[T comparable] struct {
	prev *ListNode[T]
	next *ListNode[T]
	Val  T
}

type Link[T comparable] struct {
	root *ListNode[T]
	n    int
}

func (list *Link[T]) Len() int {
	return list.n
}
func NewList[T comparable]() *Link[T] {
	var rt = NewNode[T]()
	rt.next = rt
	rt.prev = rt
	return &Link[T]{
		root: rt,
		n:    0,
	}
}
func NewNode[T comparable]() *ListNode[T] {
	return &ListNode[T]{}
}
func (list *Link[T]) addSize(x int) {
	list.n += x
}
func (list *Link[T]) Push_back(val ...T) {
	for _, v := range val {
		rt := NewNode[T]()
		rt.Val = v
		list.push_back(rt)
	}
}
func (list *Link[T]) Push_front(val ...T) {
	for _, v := range val {
		rt := NewNode[T]()
		rt.Val = v
		list.push_front(rt)
	}
}

// link lhs->rhs
func (list *Link[T]) linkNode(lhs, rhs *ListNode[T]) {
	lhs.next = rhs
	rhs.prev = lhs
}

// insert node before rt
func (list *Link[T]) InsertBefore(rt, node *ListNode[T]) {
	if rt == list.Front_pointer() {
		list.push_front(node)
	} else {
		list.linkNode(rt.prev, node)
		list.linkNode(node, rt)
	}
}

// insert node after rt
func (list *Link[T]) InsertAfter(rt, node *ListNode[T]) {
	if rt == list.Back_pointer() {
		list.push_back(node)
	} else {
		list.linkNode(node, rt.next)
		list.linkNode(rt, node)
	}
}
func (list *Link[T]) push_back(rt *ListNode[T]) {
	list.linkNode(list.Back_pointer(), rt)
	list.linkNode(rt, list.root)
	list.addSize(1)
}
func (list *Link[T]) push_front(rt *ListNode[T]) {
	list.linkNode(rt, list.Front_pointer())
	list.linkNode(list.root, rt)
	list.addSize(1)
}
func (list *Link[T]) Front_value() T {
	return list.root.next.Val
}
func (list *Link[T]) Back_value() T {
	return list.root.prev.Val
}
func (list *Link[T]) Front_pointer() *ListNode[T] {
	return list.root.next
}

func (list *Link[T]) Back_pointer() *ListNode[T] {
	return list.root.prev
}
func (list *Link[T]) remove(rt *ListNode[T]) {
	list.linkNode(rt.prev, rt.next)
	list.addSize(-1)
	rt.next = nil
	rt.prev = nil
}
func (list *Link[T]) Erase(rt *ListNode[T]) {
	list.remove(rt)
}
func (list *Link[T]) Pop_back() {
	list.remove(list.root.prev)
}
func (list *Link[T]) Pop_front() {
	list.remove(list.root.next)
}
func (list *Link[T]) MoveToFront(rt *ListNode[T]) {
	list.remove(rt)
	list.push_front(rt)
}
func (list *Link[T]) MoveToBack(rt *ListNode[T]) {
	list.remove(rt)
	list.push_back(rt)
}
func (list *Link[T]) Next(rt *ListNode[T]) *ListNode[T] {
	if rt == list.Back_pointer() {
		return nil
	}
	return rt.next
}
func (list *Link[T]) Prev(rt *ListNode[T]) *ListNode[T] {
	if rt == list.Front_pointer() {
		return nil
	}
	return rt.prev
}
