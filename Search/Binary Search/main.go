package main

import "fmt"

func iterativeBinarySearch(a []int, target, low, high int) int {
	for low <= high {
		mid := (low + high) / 2
		if target == a[mid]{
			return mid				//Returning the position of the value in the array
		}
		if target > a[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1						//There is no value in the array
}

func recursiveBinarySearch(a []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if target == a[mid] {
		return mid
	} else if target > a[mid] {
		return recursiveBinarySearch(a,target,mid+1,high)
	} else {
		return recursiveBinarySearch(a,target,low,mid-1)
	}
}

func main()  {
	test := []int{1,2,3,4,5}
	fmt.Println(iterativeBinarySearch(test,3, 0 , len(test) - 1),recursiveBinarySearch(test,10,0,len(test) - 1))
}