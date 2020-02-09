/*
C(ij)=sigma(k=1...n)[A(ik)*B(kj)]
T(n)=theta(n^3)
*/
package main

import "fmt"

func main()  {
	test1 := [][]int{{1,2},{3,4}}
	test2 := [][]int{{1,2},{3,4}}
	fmt.Println(squareMatrixMultiply(test1,test2))
}

func squareMatrixMultiply(a,b [][]int) ([][]int) {
	var c [][]int
	for i := 0; i < len(a); i++ {
		temp := make([]int,0,0)
		for j := 0; j < len(a); j++ {
			sum := 0
			for k := 0; k < len(a); k++ {
				sum += a[i][k] * b[k][j]
			}
			temp = append(temp,sum)
		}
		c = append(c,temp)
	}
	return c
}

func squareMatrixMultiplyRecursive(a,b [][]int) ([][]int) {

}