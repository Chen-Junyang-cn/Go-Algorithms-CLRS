package main

import "fmt"

type node struct {
	key int
	prev, next *node
}

func main()  {
	head := node{
		key:  0,
		prev: nil,
		next: nil,
	}
	insert(&head,1)
	fmt.Println(head,search(&head,1))
	search(&head,1)
	delete(&head,1)
	fmt.Println(head.next)
}

func search(head *node, k int) *node {
	x := head
	for x.next != nil && x.key != k {
		x = x.next
	}
	return x
}

func insert(head *node, k int)  {
	x := new(node)
	x.next, x.key = head.next, k
	if head.next != nil {
		head.next.prev = x
	}
	head.next, x.prev = x, nil
}

func delete(head *node, k int)  {
	x := search(head,k)
	if x.prev != nil {
		x.prev.next = x.next
	} else {
		head.next = x.next
	}
	if x.next != nil {
		x.next.prev = x.prev
	}
}