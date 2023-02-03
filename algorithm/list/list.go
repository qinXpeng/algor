package qlist


type ListNode[T comparable] struct {
	Prev *ListNode[T]
	Next *ListNode[T]
	Val  T
}

type Link[T comparable] struct {
	Root *ListNode[T]
	n    int
}

func (list *Link[T] ) Len() int {
	return list.n
}
func NewList[T comparable]() *Link[T] {
	var rt = NewNode[T]()
	rt.Next = rt
	rt.Prev = rt
	return &Link[T]{
		Root: rt,
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
	lhs.Next = rhs
	rhs.Prev = lhs
}

// insert node before rt
func (list *Link[T]) InsertBefore(rt, node *ListNode[T]) {
	if rt == list.Front_pointer() {
		list.push_front(node)
	} else {
		list.linkNode(rt.Prev, node)
		list.linkNode(node, rt)
	}
}

// insert node after rt
func (list *Link[T]) InsertAfter(rt, node *ListNode[T]) {
	if rt == list.Back_pointer() {
		list.push_back(node)
	} else {
		list.linkNode(node, rt.Next)
		list.linkNode(rt, node)
	}
}
func (list *Link[T]) push_back(rt *ListNode[T]) {
	list.linkNode(list.Back_pointer(), rt)
	list.linkNode(rt, list.Root)
	list.addSize(1)
}
func (list *Link[T]) push_front(rt *ListNode[T]) {
	list.linkNode(rt, list.Front_pointer())
	list.linkNode(list.Root, rt)
	list.addSize(1)
}
func (list *Link[T]) Front_value() T {
	return list.Root.Next.Val
}
func (list *Link[T]) Back_value() T {
	return list.Root.Prev.Val
}
func (list *Link[T]) Front_pointer() *ListNode[T] {
	return list.Root.Next
}

func (list *Link[T]) Back_pointer() *ListNode[T] {
	return list.Root.Prev
}
func (list *Link[T]) remove(rt *ListNode[T]) {
	list.linkNode(rt.Prev, rt.Next)
	list.addSize(-1)
	rt.Next = nil
	rt.Prev = nil
}
func (list *Link[T]) Erase(rt *ListNode[T]) {
	list.remove(rt)
}
func (list *Link[T]) Pop_back() {
	if list != nil{
		list.remove(list.Root.Prev)
	}
}
func (list *Link[T]) Pop_front() {
	if list !=nil{
		list.remove(list.Root.Next)
	}
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
	return rt.Next
}
func (list *Link[T]) Prev(rt *ListNode[T]) *ListNode[T] {
	if rt == list.Front_pointer() {
		return nil
	}
	return rt.Prev
}
