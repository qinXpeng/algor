package rb_tree

type RbTree struct {
	root *TreeNode
	less LessFunc
}
type LessFunc func(i, j Value) bool

func NewRbTree(less LessFunc, val ...Value) *RbTree {
	t := &RbTree{
		less: less,
	}
	t.root = nil
	t.Insert(val...)
	return t
}
func (r *RbTree) Empty() bool {
	return r.Len() == 0
}
func (r *RbTree) Len() int {
	if r.root == nil {
		return 0
	}
	return r.root.numSon()
}
func (r *RbTree) Less(a, b *TreeNode) bool {
	return r.less(a.val, b.val)
}
func (r *RbTree) Insert(val ...Value) {
	for _, key := range val {
		r.insert(key)
	}
}
func (r *RbTree) insert(key Value) {
	if r.root == nil {
		r.root = NewTreeNode(key)
		r.root.setBlack()
		return
	}
	r.push(r.root, NewTreeNode(key))
}
func (r *RbTree) push(rt, item *TreeNode) {
	if r.Less(item, rt) {
		if rt.leftSon() != NIL {
			r.push(rt.leftSon(), item)
		} else {
			rt.left = item
			rt.setSon(item)
			r.upInsert(item)
		}
	} else if !r.Less(item, rt) {
		if rt.rightSon() != NIL {
			r.push(rt.rightSon(), item)
		} else {
			rt.right = item
			rt.setSon(item)
			r.upInsert(item)
		}
	}
	rt.upSon()
}
func (r *RbTree) upInsert(p *TreeNode) {
	if p.father() == nil {
		r.root = p
		p.setBlack()
		return
	}
	if p.father().isBlack() {
		return
	}
	if p.uncle().isRed() {
		p.father().setBlack()
		p.uncle().setBlack()
		p.grandpa().setRed()
		r.upInsert(p.grandpa())
		return
	}
	if p.isLeft() && p.father().isRight() {
		r.rotate_right(p)
		p = p.rightSon()
	} else if p.isRight() && p.father().isLeft() {
		r.rotate_left(p)
		p = p.leftSon()
	}
	if p.isLeft() && p.father().isLeft() {
		p.father().setBlack()
		p.grandpa().setRed()
		r.rotate_right(p.father())
	} else if p.isRight() && p.father().isRight() {
		p.father().setBlack()
		p.grandpa().setRed()
		r.rotate_left(p.father())
	}
}

func (r *RbTree) rotate_left(p *TreeNode) {
	grandpa := p.grandpa()
	parent := p.father()
	left := p.leftSon()
	parent.left = left
	if left != NIL {
		left.parent = parent
	}
	p.left = parent
	parent.parent = p
	p.parent = grandpa
	if grandpa == nil {
		r.root = p
	} else {
		if grandpa.leftSon() == parent {
			grandpa.left = p
		} else {
			grandpa.right = p
		}
	}
	parent.upSon()
	p.upSon()
}

func (r *RbTree) rotate_right(p *TreeNode) {
	grandpa := p.grandpa()
	parent := p.father()
	right := p.rightSon()
	parent.left = right
	if right != NIL {
		right.parent = parent
	}
	p.right = parent
	parent.parent = p
	p.parent = grandpa
	if grandpa == nil {
		r.root = p
	} else {
		if grandpa.leftSon() == parent {
			grandpa.left = p
		} else {
			grandpa.right = p
		}
	}
	parent.upSon()
	p.upSon()
}
func (r *RbTree) GetRoot() *TreeNode {
	return r.root
}
