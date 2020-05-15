/*
The ith item is worth vi dollars and weighs wi pounds, where vi and wi are integers.
The thief wants to take as valuable a load as possible, but he can carry at
most W pounds in his knapsack, for some integer W.(we use n instead W)
 he cannot take a fractional amount of an item or take an item more than once.
Define r[i][j] as optimal subset
r[i][j] = max(r[i-1][j], v[i] + r[i-1][j-w[i]])
T(n) = O(nW)
*/
package main

import "fmt"

func main()  {
	v, w := []int{0,60,100,120}, []int{0,10,20,30}
	fmt.Print(knapsack(v,w,3,50))
}

func knapsack(v, w []int, n, weight int) ([][]int,int) {
	r := make([][]int,n+1)
	for i := 0; i < len(r); i++ {
		r[i] = make([]int,weight+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= weight; j++ {
			if j - w[i] >= 0 { // item can be loaded in knapsack
				r[i][j] = max(r[i-1][j],v[i]+r[i-1][j-w[i]])
			} else {
				r[i][j] = r[i-1][j]
			}
		}
	}
	return r,r[n][weight]
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}