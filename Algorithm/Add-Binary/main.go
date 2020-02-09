package main

import "fmt"

func addBinary(a []int, b []int) [5]int {	//adding two n-bit binary integers
	var c [5]int
	carry := 0
	var i int
	for i = len(a);i > 1 ; i-- {		//Notice that the top digit is on the left
		c[i] = (a[i-1] + b[i-1] + carry) % 2
		carry = (a[i-1] + b[i-1] + carry) / 2
	}
	c[i] = carry
	return c
}

func main(){
	a := []int{0,1,1,1}				//represent 7 in binary
	b := []int{0,0,1,1}				//represent 3 in binary
	fmt.Println(addBinary(a,b))		//print 10
}