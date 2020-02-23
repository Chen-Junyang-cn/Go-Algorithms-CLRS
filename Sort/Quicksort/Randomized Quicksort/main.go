package main

import (
	"fmt"
	"math/rand"
)

func main() {
	test := []int{13,19,9,5,12,8,7,4,21,2,6,11}
	randomizedQuicksort(test,0,len(test)-1)
	fmt.Println(test)
}

func randomizedQuicksort(a []int, p,r int)  {
	if p < r {
		q := randomizedPartition(a,p,r)
		randomizedQuicksort(a,p,q-1)
		randomizedQuicksort(a,q,r)
	}
}

func randomizedPartition(a []int, p,r int) int {
	i := random(p,r)
	a[r], a[i] = a[i], a[r]
	return partition(a,p,r)
}

func partition(a []int, p,r int) int { //T(n)=O(n)
	key, i := a[r], p - 1
	for j := p; j < r; j++ {
		if a[j] <= key {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i+1
}

func random(a,b int) int {
	if a == b {
		return a
	}
	if rand.Intn(2) == 0 {		//The random function generates integers in the interval [0, n)
		return random(a,(a+b)/2)
	} else {
		return random((a+b)/2+1,b)
	}
}