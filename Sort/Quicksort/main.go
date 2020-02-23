package main

import "fmt"

func main()  {
	test := []int{13,19,9,5,12,8,7,4,21,2,6,11}
	quicksort(test,0,len(test)-1)
	fmt.Println(test)
}

func quicksort(a []int, p,r int)  {
	if p < r {
		q := partition(a,p,r)
		quicksort(a,p,q-1)
		quicksort(a,q,r)
	}
}

func partition(a []int, p,r int) int { //T(n)=O(n)
	key, i := a[r], p - 1
	for j := p; j < r; j++ {
		if a[j] <= key {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i+1
}