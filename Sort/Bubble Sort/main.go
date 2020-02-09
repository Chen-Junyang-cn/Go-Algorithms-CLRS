package main

import "fmt"

func bubbleSort(arr []int){
	for i := 1; i < len(arr); i++ {
		for j := i - 1;j >= 0 && arr[j] > arr [j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func main(){
	test := []int{5,2,4,6,1,3}
	bubbleSort(test)
	fmt.Println(test)
}
