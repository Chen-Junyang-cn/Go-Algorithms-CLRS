package main

import (
	"fmt"
)

type queue struct {
	array [12]int
	head, tail int
}

func main()  {
	test := queue{[12]int{1},0,1}
	dequeueHead(&test)
	fmt.Println(test)
	enqueueHead(&test,2)
	fmt.Println(test)
}

func full(q *queue) bool {
	if q.head == q.tail + 1 {
		return true
	}
	return false
}

func empty(q *queue) bool {
	if q.head == q.tail {
		return true
	}
	return false
}

func dequeueHead(q *queue) int {
	if empty(q) {
		fmt.Println("underflow")
		return 0
	}
	x := q.array[q.head]
	if q.head == len(q.array) - 1 {
		q.head = 0
	} else {
		q.head++
	}
	return x
}

func dequeueTail(q *queue) int {
	if empty(q) {
		fmt.Println("underflow")
		return 0
	}
	if q.tail == 0 {
		q.tail = len(q.array) - 1
	} else {
		q.tail--
	}
	x := q.array[q.tail]
	return x
}

func enqueueHead(q *queue, x int)  {
	if full(q) {
		fmt.Println("overflow")
		return
	}
	if q.head == 0 {
		q.head = len(q.array) - 1
	} else {
		q.head--
	}
	q.array[q.head] = x
}

func enqueueTail(q *queue, x int)  {
	if full(q) {
		fmt.Println("overflow")
		return
	}
	q.array[q.tail] = x
	if q.tail == len(q.array)-1 {
		q.tail = 0
	} else {
		q.tail++
	}
}