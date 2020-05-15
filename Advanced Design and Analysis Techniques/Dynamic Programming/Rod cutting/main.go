/*
Given a rod of length n inches and a table of prices pi for i = 1,2,...n,
determine the maximum revenue rn obtainable by cutting up the rod and selling the pieces.
r(n) = max(p(i) + r(n-i))	1 <= i <= n
*/
package main

import (
	"fmt"
	"math"
)

func main()  {
	p := []int{0,1,5,8,9,10,17,17,20,24,30} //price table for rods
	fmt.Println(memoizedCutRod(p,4))
	fmt.Println(bottomUpCutRod(p,4))
	printCutRodSolution(p,4)
}


func bottomUpCutRod(p []int, n int) int { //the basic code
	r := make([]int,n+1)
	for i := 1; i <= n; i++ {
		q := 0
		var j int
		for j = 1; j <= i; j++ {
			q = max(q,p[j]+r[i-j])
		}
		r[i] = q
	}
	return r[n]
}

func extendBottomUpCutRod(p []int, n int) ([]int, []int) {
	r, s := make([]int,n+1), make([]int,n+1)
	for j := 1; j <= n; j++ {
		q := 0
		for i := 1; i <= j; i++ {
			if q < p[i] + r[j-i] {
				q = p[i] + r[j-i]
				s[j] = i
			}
		}
		r[j] = q
	}
	return r,s
}

func printCutRodSolution(p []int, n int)  {
	var r, s []int
	r, s = extendBottomUpCutRod(p,n)
	for n > 0 {
		fmt.Println(s[n],r[n])
		n -= s[n]
	}
}

func memoizedCutRod(p []int, n int) int { //use memo
	r := make([]int,n+1) //we have already initialized its element to 0
	return memoizedCutRodAux(p,r,n)
}

func memoizedCutRodAux(p, r []int, n int) int { //auxiliary function
	if n == 0 {
		return 0
	}
	if r[n] > 0 {
		return r[n]
	}
	q := 0
	for i := 1; i <= n; i++ {
		q = max(q,p[i]+memoizedCutRodAux(p,r,n-i))
	}
	r[n] = q
	return q
}

func cutRod(p []int, n int) int { //the basic code
	if n == 0 {
		return 0
	}
	q := math.MinInt32
	for i := 1; i <= n; i++ {
		q = max(q,p[i]+cutRod(p,n-i))
	}
	return q
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}