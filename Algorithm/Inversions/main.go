/*
Let A[1..n] be an array of n distinct numbers.
If i < j and A[i] > A[j],
then the pair (i,j) is called an inversion of A.
*/
package main

import (
	"fmt"
	"math"
)

func main()  {
	test := []int{2,3,8,6,1}
	fmt.Println(countInversions(test,0,len(test)-1))
}

func countInversions(a []int, l,r int) int {
	if l < r {
		mid := (l + r) / 2
		left := countInversions(a,l,mid)
		right := countInversions(a,mid+1,r)
		inversions := mergeInversions(a,l,mid,r) + left +right
		return inversions		//returns the number of inversions in the elements of A[p..r]
	}
	return 0
}

func mergeInversions(a []int, l,mid,r int) int {
	n1 := mid - l
	n2 := r - mid
	L,R := make([]int,0,0),make([]int ,0,0)
	for i := 0; i <= n1; i++ {
		L = append(L,a[l + i])
	}
	for i := 0; i < n2; i++ {
		R = append(R,a[mid + i + 1])
	}
	L = append(L,math.MaxInt32)
	R = append(R,math.MaxInt32)
	i,j,inversions := 0,0,0
	for k := l; k <= r; k++ {			//merge
		if L[i] <= R[j]{
			a[k] = L[i] 				//may not be needed
			i++
		} else {
			inversions += n1 - i + 1 	//Summing these up gives the total number of inversions in A.
			a[k] = R[j]
			j++
		}
	}
	return inversions
}