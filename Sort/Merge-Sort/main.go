package main

import (
	"fmt"
	"math"
)

func main() {
	test := []int{3,41,52,26,38,57,9,49}
	mergeSort(test,0,len(test) - 1)//末尾的值注意要减一符合数组的习惯
	fmt.Println(test)
}

func mergeSort(a[]int ,p,r int)(){
	if p < r {
		q := (p + r) / 2
		mergeSort(a,p,q)
		mergeSort(a,q + 1,r)
		merge(a,p,q,r)
	}
}

func merge(a []int,p,q,r int) () {
	n1 := q - p
	n2 := r - q
	L,R := make([]int,0,0),make([]int ,0,0)
	for i := 0; i <= n1; i++ {
		L = append(L,a[p + i])
	}
	for i := 0; i < n2; i++ {
		R = append(R,a[q + i + 1])
	}
	L = append(L,math.MaxInt32)
	R = append(R,math.MaxInt32)
	/*Here, we use max as the sentinel value, so that whenever the max is exposed,
	it cannot be the smaller value unless both piles have their sentinel cards exposed.*/
	i,j := 0,0
	for k := p; k <= r; k++ {		//merge
		if L[i] < R[j]{
			a[k] = L[i]
			i++
		} else {
			a[k] = R[j]
			j++
		}
	}
}