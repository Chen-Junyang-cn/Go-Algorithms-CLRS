/*
For given set of probabilities(p,q), we wish to construct a binary search tree whose expected search cost if smallest.
Define e[i,j] as the expected cost of searching an optimal binary search tree containing the keys k[i]...k[j].
Define w[i,j] as the sum of probabilities in the subtrees.
Define root[i,j] as the recording the root of the subtree
w[i,j] = w[i,j-1] + p[j] + q[j]
e[i,j] = e[i,r-1] + e[r+1,j] + w[i,j]	if i <= j
       = q[i-1] 						if j = i -1
 */
package main

import (
	"fmt"
	"math"
)

func main()  {
	p := []float32{0,0.15,0.10,0.05,0.10,0.20}
	q := []float32{0.05,0.10,0.05,0.05,0.05,0.10}
	e, s := optimalBST(p,q,5)
	fmt.Println(e,s)
}

func optimalBST(p, q []float32, n int) ([][]float32,[][]int) {
	e, w, root := make([][]float32,n+2), make([][]float32,n+2), make([][]int,n+1)
	for i := 0; i <= n + 1; i++ {
		e[i], w[i] = make([]float32,n+1), make([]float32,n+1)
	}
	for i := 0; i <= n; i++ {
		root[i] = make([]int,n+1)
	}
	for i := 1; i <= n + 1; i++ {
		e[i][i-1], w[i][i-1] = q[i-1],q[i-1]
	}
	for l := 1; l <= n; l++ { // calculate e[i,i+l]
		for i := 1; i <= n - l + 1; i++ {
			j := i + l - 1
			e[i][j] = math.MaxInt32
			w[i][j] = w[i][j-1] + p[j] + q[j]
			for r := i; r <= j; r++ { // find the optimal root
				t := e[i][r-1] + e[r+1][j] + w[i][j]
				if t < e[i][j] {
					e[i][j], root[i][j] = t, r
				}
			}
		}
	}
	return e,root
}