/*
Search, Minimum, Maximum, Successor, and Predecessor run in O(h) time of BST of height h
*/
package main

import (
	"fmt"
)

type node struct {
	key int
	left, right, parent *node
}

func main()  {
	root := node{
		key:    0,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	z := node{
		key:    1,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	test := node{
		key:    2,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	treeInsert(&root,&z)
	fmt.Println(z,treeSuccessor(&root),treeSearch(&root,0))
	treeDelete(&root)
	fmt.Println(root,root.right,root.parent,test)
}

func treeDelete(z *node)  { //z is a node to be deleted
	if z.left == nil {
		transplant(z,z.right)
	} else if z.right == nil {
		transplant(z,z.left)
	} else {
		y := treeMinimum(z.right) // find the successor of z
		if y.parent.right != z {
			transplant(y,y.right) // it's important to know that successor has no left child
			y.right = z.right 	  // let y replace the z.right
			y.right.parent = y
		}
		transplant(z,y)
		y.left = z.left
		y.left.parent = y
	}
}

func transplant(u, v *node)  { //replaces the subtree rooted at u with the subtree rooted at v
	if u.parent == nil { // u is the root. Let's u replace with v
		u.key = v.key
		u.left, u.right = v.left, v.right
	} else if u == u.parent.left { //u is either a left child or a right child of its parent
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func treeInsert(root, z *node)  {
	y := new(node)
	x := root
	for x != nil {//find an appropriate position in the tree
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {	//If we have a root node when we initialize,
		root = z 	//tree isn't empty
	} else if z.key < y.key { // insert to the tree
		y.left = z
	} else {
		y.right = z
	}
}

func treeSuccessor(root *node) *node { // the successor is the smallest key greater than root.key
	if root.right != nil {
		return treeMinimum(root.right)
	}
	y := root.parent
	//go up the tree form root until we encounter a node that is the left child of is parent
	for y != nil && root == y.right {
		root = y
		y = y.parent
	}
	return y
}

func treePredecessor(root *node) *node { // the predecessor is the biggest key less than root.key
	if root.left != nil {
		return treeMaximum(root.left)
	}
	y := root.parent
	for y != nil && root == y.left {
		root = y
		y = y.parent
	}
	return y
}

func iterativeTreeSearch(root *node, k int) *node {
	for root != nil && k != root.key {
		if k < root.key {
			root = root.left
		} else {
			root = root.right
		}
	}
	return root
}

func treeSearch(root *node, k int) *node {
	if root == nil || root.key == k {
		return root
	}
	if k < root.key {
		return treeSearch(root.left,k) // search continues in the left subtree of x
	}
	return treeSearch(root.right,k) // search right subtree
}

func inorderTreeWalk(root *node)  { // takes theta(n) time
	if root != nil {
		inorderTreeWalk(root.left)
		fmt.Println(root.key)
		inorderTreeWalk(root.right)
	}
}

func treeMinimum(root *node) *node {
	for root.left != nil {
		root = root.left
	}
	return root
}

func treeMaximum(root *node) *node {
	for root.right != nil {
		root = root.right
	}
	return root
}