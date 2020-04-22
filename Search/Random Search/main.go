package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	test, target := []int{1,2,3,4}, 2
	fmt.Println(randomSearch(test,target))

}

func randomSearch(a []int, target int) int {
	view, n := make([]int,len(a)), len(a)
	for count :=1; count != n ;  {
		i := rand.Intn(n)
		if i == target {
			return i
		} else if view[i] == 0 {	//某元素已经被查看过，则被标记为1，否则为0
			count++					//记录数组中以已查看的元素
			view[i] = 1
		}
	}
	return 0
}