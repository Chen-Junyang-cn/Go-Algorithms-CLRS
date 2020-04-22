/*
important property:numbers with the value appear in the output array in the same order as they do in the input array.
 */
package main

import "fmt"

func main()  {
	test := []int{2,5,3,0,2,3,0,3}
	fmt.Println(countingSort(test,findMaxNumber(test)))
}

func countingSort(a []int, k int) ([]int,[]int) {
	res, c := make([]int,len(a)), make([]int,k+1)
	for j := 0; j < len(a); j++ {
		c[a[j]] = c[a[j]] + 1 	//contains the number of elements equal to each position.
	}
	for i := 1; i < len(c); i++ {
		c[i] = c[i] + c[i-1] 	//contains the number of elements less than or equal to each position.
	}
	for j := len(a)-1; j >= 0; j-- {
		res[c[a[j]]-1] = a[j] 	//c[a[j]]-1 is the correct final position of a[j] in the output array
		c[a[j]]--
	}
	return c,res
}

func findMaxNumber(a []int) int {
	max := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}