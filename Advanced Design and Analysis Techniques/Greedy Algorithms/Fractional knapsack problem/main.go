/*
The setup is the same as 0-1 Knapsack problem, but the thief can take fractions of items
We assume that the cost performance of items is sorted in descending order.
 */
package main

import "fmt"

func main()  {
	v, w := []int{0,60,100,120}, []int{1,10,20,30}
	fmt.Println(greedy(v,w,3,50))
}

func greedy(v, w []int, n, weight int) []float32 {
	r := make([]float32,n+1)
	p := make([]float32,len(v)) //p[i] is the price ratio of the ith item
	for i := 0; i < len(p); i++ {
		p[i] = float32(v[i] / w[i])
	}
	for i := 1; i <= n && weight >= 0; i++ {
		if w[i] < weight {
			r[i] = p[i] * float32(w[i]) + r[i-1]
			weight -= w[i]
		} else {
			r[i] = p[i] * float32(weight) + r[i-1]
		}
	}
	return r
}