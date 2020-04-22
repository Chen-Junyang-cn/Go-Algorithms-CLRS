package main

import "fmt"

func main()  {
	test := []int{78,17,39,26,72,94,21,12,23,68}
	fmt.Println(bucketSort(test))
}

func bucketSort(a []int) [][]int {
	b, n, max := make([][]int,len(a)), len(a), findMaxNumber(a)
	for i := 0; i < len(a); i++ { 	//puts them into the sorted bucket
		index := (a[i] * (n - 1)) / max
		b[index] = append(b[index],a[i])
	}
	for i := 0; i < len(b); i++ { 	//puts them into the proper order in the bucket
		insertionSort(b[i])
	}
	return b
}

func insertionSort(arr []int){
	for i := 1; i < len(arr); i++  {
		key := arr[i]				//Insert A[i] into the sorted sequence A[1...i-1]
		j := i - 1
		for j >= 0 && key < arr[j]{//ascending sort
			arr[j+1]=arr[j]
			j--
		}
		arr[j+1] = key
	}
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