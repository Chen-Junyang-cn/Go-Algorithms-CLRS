/*
输入长度为n的数组 A[0, …, n-1], 其元素数值代表每个人的能力。
假定初始时助手的能力为0，且每次遇到能力高于当前助手能力的人，便辞掉当前助手，并雇佣该人.
每次雇佣需要花费费用ch，求总费用。
品均总为费用：O(cn*lnn)

对雇佣问题进行修改，现在我们只雇佣一次：
选择一个正整数k < n,面试前k个应聘者然后拒绝他们，再雇佣其后第一个比前面所有应聘者分数都高的人，如果没有这个人，就雇佣最后一个人。
k=n/e
*/
package main

import (
	"fmt"
	"math"
)

func main()  {
	test := []int{1,2,3,4}
	fmt.Println(hireAssistant(test))
	fmt.Println(onlineHiring(test))
}

func hireAssistant(a []int) (int,int) {
	best, cost := 0, 0 // candidate 0 is a least-qualified dummy candidate
	for i := 0; i < len(a); i++ {
		if a[i] > best {
			best = i
			cost++
		}
	}
	return best, cost
}

func onlineHiring(a []int) int {
	bestCore := math.MinInt32
	k := len(a)/2
	for i := 0; i <= k; i++ {
		if a[i] > bestCore {
			bestCore = a[i]
		}
	}
	for i := k+1; i < len(a); i++ {
		if a[i] > bestCore {
			bestCore = a[i]
			return i
		}
	}
	return len(a)-1
}