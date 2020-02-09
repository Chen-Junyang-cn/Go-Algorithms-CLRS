package main

import(
	"fmt"
)

func linearSearch(arr []int, target int) int{	//Returning the position of the value in the array
	for i := 0; i < len(arr); i++{
		if arr[i] == target{
			return i
		}
	}
	return -1	//There is no value in the array
}

func main(){
	test := []int{1,2,3,4,5}
	fmt.Println(linearSearch(test,3),linearSearch(test,10))
}
