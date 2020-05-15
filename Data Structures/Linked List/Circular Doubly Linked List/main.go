package main

import "fmt"

type node struct {
	key int
	prev, next *node
}

func main()  {
	head := initial()
	fmt.Println(head)
	insert(head,1)
	insert(head,2)
	fmt.Println(head.next.next)
	delete(head,2)
	delete(head,1)
	enqueue(head,1)
	fmt.Println(head.next)
	enqueue(head,2)
	fmt.Println(head.next.next)
	fmt.Println(dequeque(head),dequeque(head))
}

func initial() *node {
	head := node{
		key:  0,
		prev: nil,
		next: nil,
	}
	head.prev, head.next = &head, &head
	return &head
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
	head.next.prev = x
	head.next = x
	x.prev = head
}

func delete(head *node, k int)  {
	x := search(head,k)
	x.prev.next = x.next
	x.next.prev = x.prev
}

func enqueue(head *node, k int)  { 		//inserts an element at the end of the list.
	x := new(node)
	x.key = k
	x.prev = head.prev
	x.next = head
	head.prev.next = x
	head.prev = x

}

func dequeque(head *node) *node {		//removes an element from the beginning of the list.likes the Dequeue
	x := head.next
	head.next = head.next.next
	head.next.prev = head
	return x
}