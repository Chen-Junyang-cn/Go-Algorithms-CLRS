package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++  {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {		//ascending sort
			if arr[j] < arr[minIndex]{
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i] 	//Exchange two sets of values using multiple assignments
		/*temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp*/
	}
}

func main(){
	test := []int{5,2,4,6,1,3}
	selectionSort(test)
	fmt.Println(test)
}
