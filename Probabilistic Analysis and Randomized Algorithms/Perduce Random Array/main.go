/*
T(n)=O(n)
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	test := []int{1,2,3,4}
	randomizeInPlace(test)
	fmt.Println(test)
}

func randomizeInPlace(a []int)  {
	for i := 0; i < len(a); i++ {
		temp := random(i,len(a)-1)
		a[i],a[temp] = a[temp],a[i]
	}
}

func random(a,b int) int {		//returns an integer between "a" and "b"
	if a == b {
		return a
	}
	if rand.Intn(2) == 0 {
		return random(a,(a+b)/2)
	} else {
		return random((a+b)/2+1,b)
	}
}