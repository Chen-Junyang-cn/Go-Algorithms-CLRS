/*
go语言结构体作为函数参数，采用的是值传递。所以对于大型结构体传参，考虑到值传递的性能损耗，最好能采用指针传递。
指针传递可以改变外层值，而值传递不行
 */
package main

import (
	"fmt"
)

type stack struct {
	array [7]int
	top int
}

func main()  {
	test := stack{[7]int{15,6,2,9}, 3}
	fmt.Println(test)
	push(&test,17)
	fmt.Println(test)
	pop(&test)
	fmt.Println(test)
}

func push(s *stack, x int)  {
	s.top++
	s.array[s.top] = x
}

func pop(s *stack) int {
	if s.top == 0 {
		fmt.Println("underflow")
	} else {
		s.top--
		return s.array[s.top+1]
	}
	return 0
}

func stackEmpty(s *stack) bool {
	if s.top == 0 {
		return true
	}
	return false
}
