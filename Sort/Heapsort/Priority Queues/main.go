/*
Implement the operations of a max-priority queue.
*/
package main

import (
	"fmt"
	"math"
)

func main()  {
	test := []int{15,13,9,5,12,8,7,4,0,6,2,1}
	fmt.Println(heapExtractMax(test,len(test)-1))
}

/* It takes as an input the key of the new element to be inserted into max-heap A.*/
func maxHeapInsert(a []int,key int) []int {
	a = append(a,math.MinInt32)
	heapIncreaseKey(a,len(a)-1,key)
	return a
}
/* find a proper place for the newly increased key.*/
func heapIncreaseKey(a []int,i,key int)  { //T(n)=O(lgn)
	if key < a[i] { 	//since the max-heap property now holds
		fmt.Errorf("new key is smaller than current key")
	}
	for i > 0 && a[parent(i)] < key { // traverses this path, it repeatedly compares an element to its parent
		a[i], i = a[parent(i)], parent(i)
	}
	a[i] = key
}

/*view the max*/
func heapMaximum(a []int) int {
	return a[0]
}

/*extract the max number*/
func heapExtractMax(a []int, heapSize int) (int, []int) { //T(n)=O(lgn)
	if heapSize < 0 {
		fmt.Errorf("heap underflow")
	}
	max := a[0]
	a[0], a = a[heapSize], a[0:heapSize]
	heapSize--
	maxHeapify(a,0,heapSize)
	return max,a
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

func left(i int) int {
	return 2*i+1
}

func right(i int) int {
	return 2*i+2
}

func parent(i int) int {
	return (i-1)/2
}