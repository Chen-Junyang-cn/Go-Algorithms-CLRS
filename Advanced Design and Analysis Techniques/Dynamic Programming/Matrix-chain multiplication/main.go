/*
We are given a sequence(chain) <A1,A2,...,An> of n matrices to be multiplied, and we wish to compute the product A1A2..An/
m[i,j] = m[i,k] + m[k+1,j] +p(i-1)p(k)p(j) (i < j)
       = 0  i==j
*/
package main

import (
	"fmt"
	"math"
)

func main()  {
	test := [][]int{{1,2}}
	test2 := [][]int{{1},{2}}
	fmt.Println(matrixMultiply(test,test2))
	p := []int{5,10,3,12,5,50,6}
	var m, s [][]int = matrixChainOrder(p)
	fmt.Println(m,s)
	printOptimalParens(s,1,6)
	var m2, s2 [][]int = memoizedMatrixChain(p,1,6)
	fmt.Println(m2,s2)
}

func printOptimalParens(s [][]int, i, j int)  {
	if i == j {
		fmt.Print("A",i)
	} else {
		fmt.Print("(")
		printOptimalParens(s,i,s[i][j])
		printOptimalParens(s,s[i][j]+1,j)
		fmt.Print(")")
	}
}

func memoizedMatrixChain(p []int, i, j int) ([][]int, [][]int) {
	n := len(p) - 1
	m, s := make([][]int,n+1), make([][]int,n)
	for i := 0; i < len(m); i++ {
		m[i] = make([]int,n+1)
		for j := 0; j < len(m); j++ {
			m[i][j] = math.MaxInt32 //indicate that the entry has yet to be filled in
		}
	}
	for i := 0; i < len(s); i++ {
		s[i] = make([]int,n+1)
	}
	lookUpChain(m,s,p,i,j)
	return m,s
}

func lookUpChain(m, s [][]int, p []int, i, j int) int {
	if m[i][j] < math.MaxInt32 { // the entry has been filled in.
		return m[i][j]
	}
	if i == j {
		m[i][j] = 0
	} else {
		for k := i; k <= j - 1; k++ {
			q := lookUpChain(m,s,p,i,k) + lookUpChain(m,s,p,k+1,j) + p[i-1]*p[k]*p[j]
			if q < m[i][j] {
				m[i][j] = q
				s[i][j] = k
			}
		}
	}
	return m[i][j]
}

func matrixChainOrder(p []int) ([][]int, [][]int) { // p is a dimensional array of matrix chains
	n := len(p) - 1
	m, s := make([][]int,n+1), make([][]int,n)
	/* uses table m for storing the m[i,j] costs and
	table s that records which index of k achieved the optimal cost in computing m[o,j] */
	for i := 0; i < len(m); i++ {//The array element initializes to 0, so it's a minimum
		m[i] = make([]int,n+1) 	 //so we have already set m[i,i] = 0
	}
	for i := 0; i < len(s); i++ {
		s[i] = make([]int,n+1)
	}
	for l := 2; l <= n; l++ { //l is the chain length
		for i := 1; i <= n - l + 1; i++ { // compute the minimum costs for chains of length l
			j := i + l -1
			m[i][j] = math.MaxInt32
			for k := i; k <= j - 1; k++ { // k is the split index
				q := m[i][k] + m[k+1][j] + p[i-1] * p[k] * p[j]
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}
	return m,s
}

func matrixMultiply(a, b [][]int) [][]int {
	c := make([][]int,0) // let c be a new a.rows * b.columns matrix
	for i := 0; i < len(a); i++ { // a.rows
		temp := make([]int,0)
		for j := 0; j < len(b[0]); j++ { //b.columns
			sum := 0
			for k := 0; k < len(a[0]); k++ { // a.columns = b.rows
				sum += a[i][k] * b[k][j]
			}
			temp = append(temp,sum)
		}
		c = append(c,temp)
	}
	return c
}