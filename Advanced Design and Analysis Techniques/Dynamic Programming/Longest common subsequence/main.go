/*
In the longest-common-subsequence problem, we are given two sequences
X = <x1, x2, ..., xm > and Y = <y1, y2, ..., yn> and wish to find a maximum-
length common subsequence of X and Y .
Define c[i,j] to be the length of an LCS of the sequences Xi, Yi.
c[i,j] = c[i-1,j-1] + 1 		i,j > 0 and xi == yi
	   = max(c[i-1,j],c[i,j-1]) i,j > 0 and xi != yi
	   = 0						i = 0 or j = 0
*/
package main

import "fmt"

func main()  {
	x, y := []string{"A","B","C","B","D","A","B"}, []string{"B","D","C","A","B","A"}
	b, c := LCSLength(x,y)
	fmt.Println(c)
	printLCS(b,x,len(x),len(y))
}

func printLCS(b [][]string, x []string, i, j int)  { // i,j is the length of x,y + 1
	if i == 0 || j == 0 {
		return
	}
	if b[i][j] == "upleft" {
		printLCS(b,x,i-1,j-1)
		fmt.Print(x[i-1])
	} else if b[i][j] == "up" {
		printLCS(b,x,i-1,j)
	} else {
		printLCS(b,x,i,j-1)
	}
}

func LCSLength(x, y []string) ([][]string,[][]int) { //bottom-up
	m, n := len(x), len(y)
	b, c := make([][]string,m+1), make([][]int,m+1) // table b helps us construct an optimal solution.
	for i := 0; i <= m; i++ {
		b[i], c[i] = make([]string,n+1), make([]int,n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] { 	// The index of the two sequences starts at 0
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = "upleft"
			} else if c[i-1][j] >= c[i][j-1] { // go up
				c[i][j] = c[i-1][j]
				b[i][j] = "up"
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = "left"
			}
		}	
	}
	return b,c
}