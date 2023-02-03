package rb_tree

type Value interface {
}
type Color int

const (
	red Color = iota
	black
)

var NIL = &TreeNode{color: black, son: 0}

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	color  Color
	val    Value
	son    int
}

func NewTreeNode(val Value) *TreeNode {
	return &TreeNode{
		left:   NIL,
		right:  NIL,
		parent: nil,
		color:  red,
		val:    val,
		son:    1,
	}
}
func (t *TreeNode) father() *TreeNode {
	return t.parent
}
func (t *TreeNode) leftSon() *TreeNode {
	return t.left
}
func (t *TreeNode) rightSon() *TreeNode {
	return t.right
}
func (t *TreeNode) grandpa() *TreeNode {
	if t.parent == nil {
		return nil
	}
	return t.parent.parent
}
func (t *TreeNode) sibing() *TreeNode {
	if t.parent.left == t {
		return t.parent.right
	}
	return t.parent.left
}
func (t *TreeNode) uncle() *TreeNode {
	if t.grandpa() == nil {
		return nil
	}
	if t.grandpa().left == t.parent {
		return t.grandpa().right
	}
	return t.grandpa().left
}
func (t *TreeNode) setBlack() {
	t.color = black
}
func (t *TreeNode) setRed() {
	t.color = red
}
func (t *TreeNode) numSon() int {
	return t.son
}
func (t *TreeNode) upSon() {
	t.son = t.leftSon().numSon() + t.rightSon().numSon() + 1
}
func (t *TreeNode) isLeft() bool {
	return t.father().leftSon() == t
}
func (t *TreeNode) isRight() bool {
	return t.father().rightSon() == t
}
func (t *TreeNode) isBlack() bool {
	return t.color == black
}
func (t *TreeNode) isRed() bool {
	return t.color == red
}
func (t *TreeNode) setSon(p *TreeNode) {
	p.parent = t
}

func DFS(rt *TreeNode) []Value {
	if rt == NIL {
		return nil
	}
	a := DFS(rt.leftSon())
	a = append(a, rt.val)
	b := DFS(rt.rightSon())
	a = append(a, b...)
	return a
}
