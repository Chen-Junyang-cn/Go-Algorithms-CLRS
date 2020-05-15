/*
In open addressing,all elements occupy the hash table itself.
Using a[][0] represents the key, a[][1] represents the slot is nil or not
 */
package main

import (
	"fmt"
)

type table struct {
	t []int
}

var m int;//the length of the array
func main()  {
	fmt.Scan(&m)
	test := make([][2]int, m)
	hashInsert(test,60)
	fmt.Println(test,hashSearch(test,60))
}

func hashInsert(t [][2]int, k int) int {

	for i := 0; i < len(t); i++ {
		j := h(k,i);
		if t[j][1] == 0 {
			t[j][0], t[j][1] = k, 1 //insert the key and the pointer does not point to the nil
			return j
		}
	}
	fmt.Println("hash table overflow")
	return 0
}

func hashSearch(t [][2]int ,k int) int {
	for i := 0; i < len(t); i++ {
		j := h(k,i)
		if t[j][1] == 0 {
			fmt.Println("the key isn't in the table")
			return -1
		}
		if t[j][0] == k {//find the value
			return j
		}
	}
	return 0
}

func h(k, i int) int { //auxiliary hash function : linear probing
	return (k + i) % m
}