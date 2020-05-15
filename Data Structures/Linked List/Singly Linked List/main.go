package main

import "fmt"

type node struct {
	key int
	next *node 		//pointer filed : The pointer points to the next node
}

func main()  {
	head := node{
		key:  1,
		next: nil,
	}
	fmt.Println(search(&head,1))
	insert(&head,2)
	fmt.Println(head.next)
	delete(&head,2)
	fmt.Println(head.next)
}

func empty(head *node) bool {
	if head.next == nil {
		return true
	}
	return false
}

func insert(head *node, k int)  {	//adds an element in the beginning of the list.likes the stack PUSH
	x := new(node)
	x.key = k
	x.next = head.next
	head.next = x
}

func search(head *node, k int) *node {
	x := head
	for x.next != nil && x.key != k {
		x = x.next
	}
	return x
}

func delete(head *node, k int)  {
	x := head
	for x.next != nil && x.next.key != k { //寻找下一个节点为k的节点
		x = x.next
	}
	x.next = x.next.next
}

func pop(head *node) *node {	//removes the first element from the list.
	if empty(head) {
		fmt.Println("underflow")
	} else {
		x := head.next
		head.next = head.next.next
		return x
	}
	return head
}