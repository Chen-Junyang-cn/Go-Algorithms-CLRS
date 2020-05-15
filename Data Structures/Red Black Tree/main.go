/*
RB tree are balanced that basic dynamic-set operations take O(lgn) time in the worse case.

 */
package main

import "fmt"

type node struct {
	key int
	left, right, parent *node
	color string
}

type tree struct {
	root *node
}

func main()  {
	x := node{
	key:    41,
	left:   nil,
	right:  nil,
	parent: nil,
	color:  "black",
	}
	t := tree{root:nil}
	y := node{
		key:    38,
		left:   nil,
		right:  nil,
		parent: nil,
		color:  "red",
	}
	z,a := node{
		key:    31,
		left:   nil,
		right:  nil,
		parent: nil,
		color:  "red",
	},node{
		key:    12,
		left:   nil,
		right:  nil,
		parent: nil,
		color:  "red",
	}
	RBInsert(&t,&x)
	RBInsert(&t,&y)
	RBInsert(&t,&z)
	RBInsert(&t,&a)
	fmt.Println(t.root,t.root.left,t.root.right,t.root.left.left)
	RBDelete(&t,&a)
	fmt.Println(t.root,t.root.left,t.root.right,t.root.left.left)
}

func RBDelete(t *tree, z *node)  {
	x := new(node) 				// keep track of the node x that moves into node y's original position
	y := z
	yOriginalColor := y.color // node y's color might change. save it in order to test it at the end.
	if z.left == nil {
		x = z.right 			// set x point to either y's only child
		RBTransplant(t,z,z.right) // x moves into node y's original position
	} else if z.right ==nil {
		x = z.left
		RBTransplant(t,z,z.left)
	} else {
		y = treeMinimum(z.right) // find the successor
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			 x.parent = y 		// we don't want x.p point to the y's original parent(it will be removed)
		} else {
			RBTransplant(t,y,y.right)
			y.right = z.right
			y.right.parent = y
		}
		RBTransplant(t,z,y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if yOriginalColor == "black" {
		RBDeleteFixup(t,x) 		// restore red-black tree properties
	}
}

func RBDeleteFixup(t *tree, x *node)  {
	for x != t.root && x.color == "black" {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == "red" { // case1: x's sibling w is red
				w.color = "black"
				x.parent.color = "red"
				leftRotate(t,x.parent)
				w = x.parent.right
			} // converted case 1 to 2, 3 or 4, which occurs when x and w are black
			if w.left.color == "black" && w.right.color == "black" {//case 2: both of w's children are black
				w.color = "red"  // take black off w, leaving x with black
				x = x.parent 	 // the last line will compensate for removing one black
			} else {
				if w.right.color == "black" { // case 3:
					w.left.color = "black" // switch the colors of w and its left child w.left
					w.color = "red"
					rightRotate(t,w)
					w = x.parent.right
				}
				w.color = x.parent.color // case 4:
				x.parent.color = "black"
				w.right.color = "black"
				leftRotate(t,x.parent)
				x = t.root // causes the loop to terminate
			}
		} else {
			w := x.parent.left
			if w.color == "red" { // case1: x's sibling w is red
				w.color = "black"
				x.parent.color = "red"
				rightRotate(t,x.parent)
				w = x.parent.left
			} // converted case 1 to 2, 3 or 4, which occurs when x and w are black
			if w.left.color == "black" && w.right.color == "black" {//case 2: both of w's children are black
				w.color = "red"  // take black off w, leaving x with black
				x = x.parent 	 // the last line will compensate for removing one black
			} else {
				if w.left.color == "black" { // case 3:
					w.right.color = "black" // switch the colors of w and its left child w.left
					w.color = "red"
					leftRotate(t,w)
					w = x.parent.left
				}
				w.color = x.parent.color // case 4:
				x.parent.color = "black"
				w.left.color = "black"
				rightRotate(t,x.parent)
				x = t.root // causes the loop to terminate
			}
		}
	}
	x.color = "black"
}

func RBTransplant(t *tree, u, v *node)  {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil { //Accessing field 'v.parent' may lead to nil pointer dereference
		v.parent = u.parent
	}

}

func RBInsert(t *tree, z *node)  { // just like the treeInsert, but need to adjust the color
	y := t.root //y may be nil or a node
	x := t.root
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if t.root == nil { //tree is empty
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	z.color = "red"
	RBInsertFixup(t,z)
}

func RBInsertFixup(t *tree, z *node)  {
	for z.parent != nil && z.parent.color == "red" {
		if z.parent == z.parent.parent.left {
			y := new(node) // prevent y from being "nil" and causing the program to crash
			if z.parent.parent.right != nil {
				y = z.parent.parent.right
			}
			if y.color == "red" { // case1
				z.parent.color = "black"
				y.color = "black"
				z.parent.parent.color = "red"
				z = z.parent.parent
			} else {
				if z == z.parent.right { // case2
					z = z.parent
					leftRotate(t, z)
				}
				z.parent.color = "black" // case3
				z.parent.parent.color = "red"
				rightRotate(t, z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == "red" {
				z.parent.color = "black"
				y.color = "black"
				z.parent.parent.color = "red"
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					rightRotate(t,z)
				}
				z.parent.color = "black"
				z.parent.parent.color = "red"
				leftRotate(t,z.parent.parent)
			}
		}
	}
	t.root.color = "black"
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
}

func treeMinimum(root *node) *node {
	for root.left != nil {
		root = root.left
	}
	return root
}