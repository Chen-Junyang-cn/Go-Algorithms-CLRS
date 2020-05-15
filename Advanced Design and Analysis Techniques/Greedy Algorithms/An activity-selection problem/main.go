/*
Our first example is the problem of scheduling several competing activities that require exclusive
use of a common resource, with a goal of selecting a maximum-size set of mutually compatible activities.
s[i] is each activity ai start time, f[i] is the finish time.
We assume that the n input activities are already ordered by monotonically increasing finish time.
 */
package main

import (
	"fmt"
	"math"
)

func main()  {
	s, f := []int{0, 1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12, math.MaxInt32}, []int{0, 4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16, math.MaxInt32}
	res := make([]int,len(s)) // record the activity result
	recursiveActivitySelector(s,f,0,len(s)-1,res)
	fmt.Println(res)
	res2 := make([]int,len(s))
	iterativeActivitySelector(s,f,res2)
	fmt.Println(res2)
	c := dynamicMemoized(s,f,0,len(s)-1)
	fmt.Println(c)
	r, a := iterativeDP(s,f,11)
	fmt.Println(r[0])
	fmt.Println(a)
	printRes(a,0,12)
}

func printRes(r [][]int, start, end int)  {
	m := r[start][end]
	if m == -1 {
		return
	}
	printRes(r,start,m)
	fmt.Print(m," ")
	printRes(r,m,end)
}

func iterativeDP(s, f []int, n int) ([][]int,[][]int) {//11
	c, act := make([][]int,n+2), make([][]int,n+2)
	for i := 0; i < len(c); i++ {
		c[i], act[i] = make([]int,n+2), make([]int,n+2)
		for j := 0; j < n+2; j++ {
			act[i][j] = -1
		}
	}
	for l := 1; l <= n; l++ {
		for i := 0; i <= n - l; i++ {
			j := i + l + 1
			for k := i + 1; k <= j; k++ {
				if f[i] < s[k] && f[k] < s[j] {
					temp := c[i][k] + c[k][j] + 1
					if temp > c[i][j] {
						c[i][j] = temp
						act[i][j] = k
					}
				}
			}
		}
	}
	return c, act
}

func dynamicMemoized(s, f []int, k, n int) [][]int {
	c, act := make([][]int,len(s)+2), make([][]int,len(s)+2)
	for i := 0; i < len(c); i++ {
		c[i], act[i] = make([]int,len(c)), make([]int,len(act))
	}
	dynamicProgrammingSelector(s,f,k,n,c)
	return c
}

func dynamicProgrammingSelector(s, f []int, i, j int, c [][]int) int {
	if c[i][j] > 0 {
		return c[i][j]
	} else if i >= j {
		return 0
	}
	for k := i; k < j; k++ {
		q := dynamicProgrammingSelector(s,f,i,k,c) + dynamicProgrammingSelector(s,f,k+1,j,c) + 1
		if f[i] <= s[k] && f[k] <= s[j] {
			if q >= c[i][j] {
				c[i][j] = q
			}
		}
	}
	return c[i][j]
}

func recursiveActivitySelector(s, f []int, k, n int, res []int)  {
	m := k + 1
	for m <= n && s[m] < f[k] { // find the first activity in S k to finish
		m++
	}
	if m <= n {
		res[m] = 1 // record the activity in this index
		recursiveActivitySelector(s,f,m,n,res)
	}
}

func iterativeActivitySelector(s, f []int, res []int)  {
	n := len(s)-1
	k := 0 //The variable k indexes the most recent addition
	for m := k + 1; m <= n; m++ {
		if s[m] >= f[k] {
			res[m] = 1
			k = m
		}
	}
}