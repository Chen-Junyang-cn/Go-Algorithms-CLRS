/*
function MRQ : worst-case stack depth is theta(lgn)
T(n) = O(nlgn)

tail recursive : The second recursive call in Quciksort is not really necessary,
we can avoid it by using an iterative control structure.
 */
package main

import "fmt"

func main()  {
	test := []int{13,19,9,5,12,8,7,4,21,2,6,11}
	modifyRecursiveQuicksort(test,0,len(test)-1)
	fmt.Println(test)
}

func tailRecursiveQuicksort(a []int, p,r int)  {
	for p < r {
		q := partition(a,p,r)
		tailRecursiveQuicksort(a,p,q-1)
		p = q + 1
	}
}

func modifyRecursiveQuicksort(a []int, p,r int)  {
	for p < r {
		q := partition(a,p,r)
		if q < (p + r) / 2 {
			modifyRecursiveQuicksort(a,p,q-1)
			p = q + 1
		} else {
			modifyRecursiveQuicksort(a,q+1,r)
			r = q -1
		}
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