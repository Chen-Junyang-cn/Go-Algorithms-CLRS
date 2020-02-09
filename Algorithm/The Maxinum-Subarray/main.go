/*
Use divide and conquer,to find the largest subarray : T(n)=theta(nlgn)

Use brute-force method to solve the problem : T(n)=theta(n^2)

Use linear-time algorithm : T(n)=theta(n)
method:
Knowing a maximum subarray A[1..j], extend the answer to find a maximum subarray ending at index j + 1 by using the following observation:
a maximum subarray A[i..j + 1], for some 1 <= i <= j + 1.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	test := []int{1,2,-3,1}
	fmt.Println(findMaximumSubarray(test,0,len(test)-1))
	fmt.Println(bruteForce(test))
	fmt.Println(iterativeFindMaximumSubarray(test))
}

func iterativeFindMaximumSubarray(a []int) (int,int,int) {
	maxSum, sum := math.MinInt32, math.MinInt32
	low, high, currentLow := 0, 0, 0
	for j := 0; j < len(a); j++ { 	//start form the left,towards to the right
		currentHigh := j			// keeping track of the maximum subarray seen so far
		if sum > 0 {
			sum += a[j]
		} else {
			sum, currentLow = a[j], j
		}
		if sum > maxSum {
			low, high, maxSum = currentLow, currentHigh, sum
		}
	}
	return low,high,maxSum
}

func bruteForce(a []int) (int,int,int) {
	maxSum, maxLeft, maxRight := math.MinInt32, 0, 0
	for i := 0; i < len(a); i++ {
		sum := a[i]
		for j := i + 1; j < len(a); j++ {
			sum += a[j]
			if sum > maxSum {
				maxSum, maxLeft, maxRight = sum, i, j
			}
		}
	}
	return maxLeft,maxRight,maxSum
}

func findMaxCrossingSubway(a []int,low,mid,high int)(int,int,int) {
	leftSum := math.MinInt32
	maxLeft, maxRight,sum := 0,0,0
	for i := mid; i >= low; i-- {
		sum += a[i]
		if sum > leftSum {
			leftSum, maxLeft = sum, i
		}
	}
	rightSum := math.MinInt32
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += a[i]
		if sum > rightSum {
			rightSum, maxRight = sum, i
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}

func findMaximumSubarray(a []int, low,high int) (int,int,int) {
	if low == high {
		return low,high,a[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := findMaximumSubarray(a,low,mid)
		rightLow, rightHigh, rightSum := findMaximumSubarray(a,mid+1,high)
		crossLow, crossHigh, crossSum := findMaxCrossingSubway(a,low,mid,high)
		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}