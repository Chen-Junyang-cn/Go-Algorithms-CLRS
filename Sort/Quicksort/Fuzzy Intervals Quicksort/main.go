/*
task : Design a randomized algorithm for fuzzy-sorting nn intervals.
definition a[][2] : The first dimension of a two-dimensional array is the left endpoint,
					and the other is the right endpoint.
 */
package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	test := [][2]int{{3,5},{1,2},{1,6},{2,7},{1,3}}
	fuzzySort(test,0,len(test)-1)
	fmt.Println(test)
}

func fuzzySort(a [][2]int, p, r int) {
	if p < r {
		low, high := findIntersection(a,p,r)
		t := partitionRight(a,low,p,r)
		q := partitionLeft(a,high,p,t)
		fuzzySort(a,p,q-1)
		fuzzySort(a,t+1,r)
	}
}

func findIntersection(a [][2]int, p,r int) (int, int) { //T(n)=theta(n)
	rand := random(p,r)
	a[rand], a[r] = a[r], a[rand]
	left := a[r][0]
	right := a[r][1]
	for i := p; i < r; i++ {
		if a[i][0] <= right && a[i][1] >= left {
			if a[i][0] > left {
				left = a[i][0]
			}
			if a[i][1] < right {
				right = a[i][1]
			}
		}
	}
	return left, right
}

func partitionRight(a [][2]int, low,p,r int) int {
	i := p - 1
	for j := p; j < r; j++ {
		if a[j][0] < low {		//left endpoint is less than low
			i++
			a[i],a[j] = a[j],a[i]
		}
	}
	a[i+1],a[r] = a[r],a[i+1]
	return i+1
}

func partitionLeft(a [][2]int, high,p,t int ) int {
	i := p - 1
	for j := p; j < t; j++ {
		if a[j][1] < high {		//right endpoint is less than high
			i++
			a[i],a[j] = a[j],a[i]
		}
	}
	a[i+1],a[t] = a[t],a[i+1]
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