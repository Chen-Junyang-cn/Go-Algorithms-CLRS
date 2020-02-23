/*
heapSort : T(n) = theta(nlgn)
maxHeapify : T(n) <= T(2n/3) + theta(1), so T(n) = O(lgn)
buildMaxHeap : T(n) = O(n)
*/
package main

import "fmt"

func main()  {
	test := []int{4,1,3,2,16,9,10,14,8,7}
	heapsort(test)
	fmt.Println(test)
}

func heapsort(a []int)  {
	buildMaxHeap(a)
	heapSize := len(a)
	for i := len(a)-1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapSize--
		maxHeapify(a,0,heapSize)
	}

}

func buildMaxHeap(a []int)  { 					//build the heap. two subtrees of the node are both max-heaps.
	for i := (len(a)-1)/2; i >= 0; i-- {
		maxHeapify(a,i,len(a))
	}
}

func maxHeapify(a []int,i, heapSize int)  {
	l, r, largest := left(i), right(i), 0
	if l < heapSize && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i { 							//the subtree rooted at node i isn't the largest element
		a[i], a[largest] = a[largest], a[i] 	//a[i] is swapped with a[largest]
		maxHeapify(a,largest,heapSize)  		//let its children to satisfy the max-heap property
	}
}
/*The following function calculates the position of the node in the array*/
func left(i int) int {
	return 2*i+1
}

func right(i int) int {
	return 2*i+2
}

func parent(i int) int {
	return (i-1)/2
}