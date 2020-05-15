/*
An order-statistic tree T is simply a red-black tre with additional information.
We set T.nil.size to be 0, so we should notice if the child node is empty
 */
package main

import "fmt"

type node struct {
	key int
	size int
	left, right, parent *node
	color string
}

type tree struct {
	root *node
}

func main()  {
	y := node{
		key:    41,
		size:   2,
		left:   nil,
		right:  nil,
		parent: nil,
		color:  "black",
	}
	t := tree{root:&y}
	x := node{
		key:    38,
		size:   1,
		left:   nil,
		right:  nil,
		parent: &y,
		color:  "red",
	}
	y.left = &x
	rightRotate(&t,&y)
	leftRotate(&t,&x)
	fmt.Println(t.root,t.root.left,t.root.right)
	fmt.Println(OSSelect(t.root,2))
	fmt.Println(OSKeyRank(t.root,41))
	fmt.Println(OSRank(&t,&y))
}

func leftRotate(t *tree, x *node)  {
	y := x.right		// set y
	x.right = y.left	// turn y's left subtree into x's right subtree
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent // link x's parent to y, if x.p == nil ,y is the root
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
	y.size = x.size
	if x.left == nil && x.right == nil {
		x.size = 1
	} else if x.left == nil {
		x.size = x.right.size + 1
	} else if x.right == nil {
		x.size = x.left.size + 1
	} else {
		x.size = x.left.size + x.right.size + 1
	}
}

func rightRotate(t *tree,y *node)  {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x
	x.size = y.size
	if y.left == nil && y.right == nil { // notice if the child is nil
		y.size = 1
	} else if y.left == nil {
		y.size = y.right.size + 1
	} else if y.right == nil {
		y.size = y.left.size + 1
	} else {
		y.size = y.left.size + y.right.size + 1
	}
}

func OSKeyRank(root *node, k int) int { //k is a key, returns the rank of k
	if k == root.key && root.left == nil {
		return 1
	} else if k == root.key {
		return root.left.size + 1
	} else if root.key > k {
		return OSKeyRank(root.left,k)
	}
	if root.left == nil {
		return 1 + OSKeyRank(root.right,k)
	}
	return root.left.size + 1 + OSKeyRank(root.right,k)
}

func OSSelect(x *node, i int) *node { //returns a pointer to the node containing the ith smallest key in the subtree rooted at x
	r := 1
	if x.left != nil {
		r = x.left.size + 1
	}
	if i == r {
		return x
	} else if i < r {
		return  OSSelect(x.left,i)
	}
	return OSSelect(x.right,i-r)
}

func OSRank(t *tree, x *node) int { //returns the pointer of x in the linear order determined by an inorder tree walk
	r := 1
	if x.left != nil {
		r = x.left.size + 1
	}
	y := x
	for y != t.root {
		if y == y.parent.right {
			if y.parent.left != nil {
				r += y.parent.left.size + 1
			} else {
				r = 1
			}
		}
		y = y.parent
	}
	return r
}