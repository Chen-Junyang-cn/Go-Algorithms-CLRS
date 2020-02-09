package main

import "fmt"

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

func main(){
	test := []int{5,2,4,6,1,3}
	insertionSort(test)
	fmt.Println(test)
}
