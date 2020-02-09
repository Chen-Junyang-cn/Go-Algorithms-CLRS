/*
求多项式A(x) = (an)*x^n + [a(n-1)]*x^(n-1) +...+ a1*x + a0的值
*/
package main

import "fmt"

func naivePE(a []int, x int) int { 			/*朴素多项式求值*/
	y := 0
	for k := 0; k <= len(a) - 1; k++ {
		temp := 1
		for i := 1; i <= k; i++ {
			temp *= x
		}
		y += a[k] * temp
	}
	return y
}

func hornerRole(a []int, x int) int {		/*霍纳法则*/
	y := 0
	for k := len(a) - 1; k >= 0; k-- {
		y = a[k] + x * y
	}
	return y
}

func main()  {
	test := []int{5,2}
	fmt.Println(naivePE(test,1),naivePE(test,2))
	fmt.Println(hornerRole(test,1),hornerRole(test,2))
}