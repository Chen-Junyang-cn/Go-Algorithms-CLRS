/*
We augment red-black trees to support operations on dynamic set of intervals.

 */
package main

import "fmt"

type node struct {
	low, high int
	max int
	left, right, parent *node
	color string
}

type interval struct {
	low, high int
}

type tree struct {
	root *node
}

func main()  {
	x := node{
		low:    16,
		high:   21,
		max:    30,
		left:   nil,
		right:  nil,
		parent: nil,
		color:  "black",
	}
	y := node{
		low:    25,
		high:   30,
		max:    30,
		left:   nil,
		right:  nil,
		parent: &x,
		color:  "red",
	}
	z := node{
		low:    15,
		high:   23,
		max:    23,
		left:   nil,
		right:  nil,
		parent: &x,
		color:  "red",
	}
	x.right, x.left = &y, &z
	i := interval{
		low:  17,
		high: 19,
	}
	t := tree{root:&x}
	leftRotate(&t,&x)
	fmt.Println(t.root,t.root.left,t.root.left.left)
	fmt.Println(intervalSearch(&t,i))
}

func leftRotate(t *tree, x *node)  {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
	y.max = x.max
	if x.left == nil && x.right == nil {
		x.max = x.high
	} else if x.left == nil {
		x.max = max(x.right.max,x.high)
	} else if x.right == nil {
		x.max = max(x.left.max,x.high)
	} else {
		x.max = max3(x.high,x.left.max,x.right.max)
	}
}

func intervalSearchAll(x *node, i interval)  {
	if x.low <= i.high && x.high >= i.low { // i overlaps x.int
		fmt.Println(x)
	}
	if x.left != nil && x.left.max > x.low {
		intervalSearchAll(x.left,i)
	}
	if x.right != nil && x.low <= i.high && x.right.max >= x.low {
		intervalSearchAll(x.right,i)
	}
}

func intervalSearch(t *tree, i interval) *node { // finds a node in the tree t whose interval overlaps interval i.
	x := t.root
	for x != nil && (x.high < i.low || x.low > i.high) { // if i does not overlap x.int
		if x.left != nil && x.left.max >= i.low {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func max3(x, y, z int) int {
	return max(max(x,y),z)
}