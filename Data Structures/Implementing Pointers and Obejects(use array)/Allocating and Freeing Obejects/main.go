/*
Each list element is an object that occupies a contiguous subarray of length 3 within the array.
The three attributes key, next, and prev correspond to the offsets 0, 1, and 2.
Multiple array : array[0][0] means the head of the list.
                 column 0, 1, 2 means key, next and prev.
Free list: keeps the free objects in a singly linked list,and acts like a stack.
			And using 0 to present "nil"
Insert and delete are auxiliary functions
 */
package main

import (
	"fmt"
	"math/rand"
)

var free = 4 //The head of the free list is held in the global variable
var listhead = 7

func main()  {
	/*test := [][3]int{{0,0,0},{0,0,0},{4,3,5},{1,0,2},{0,8,0},{16,2,7},{0,1,0},{9,5,4},{0,6,0}}
	//let the free-list head start at 1, list head start at 2
	insert(test,25)
	fmt.Println(test[4],listhead,free,test)
	freeObeject(test,5)
	fmt.Println(test,free,listhead)*/
	free = 1
	listhead = 2
	test2 := [][3]int{{0,0,0},{0,4,6},{1,7,3},{4,2,5},{0,8,1},{16,3,7},{0,1,8},{9,5,2},{0,6,4}}
	//test2 is an circular list, and so is free list
	transpose(test2,2,7)
	transpose(test2,3,5) //sort the list
	//[[0 0 0] [0 4 6] [9 3 7] [16 5 2] [0 8 1] [4 7 3] [0 1 8] [1 2 5] [0 6 4]]
	compactify(test2)
	fmt.Println(test2)
}

func compactListSearch(a [][3]int, n, k int) int {
	/*all keys are distinct and the compact list is sorted.
	 T(n) = O(sqrt(n)) */
	i := a[listhead][1]
	for i != 0 && a[i][0] < k { // i != nil and key[i] < k
		j := rand.Intn(n) //[0,n)
		if a[i][0] < a[j][0] && a[j][0] <= k {
			i = j
			if a[i][0] == k { // find the key
				return i
			}
		}
		i = a[i][1] // i = i.next
	}
	if i == 0 || a[i][0] > k { // i==nil and key[i] > k
		return 0
	} else {
		return i
	}
}

func compactify(a [][3]int)  { //the list must be sorted
	transpose(a,listhead,1)
	if free == 1 {
		free = listhead
	}
	listhead = 1
	l := a[listhead][1] //l = list.head.next//list.next != nil
	for i := 2; l != 1; i++ { //if the list.next != list.head
		transpose(a,l,i)
		if free == i {
			free = l
		}
		l = a[i][1]//l = l.next, but notice that l exchange with i
		 fmt.Println(l,i)
	}
}

func transpose(a [][3]int, index1, index2 int)  {
	a[a[index1][2]][1], a[a[index2][2]][1] = a[a[index2][2]][1], a[a[index1][2]][1] //swap index.prev.next
	a[index1][2], a[index2][2] = a[index2][2], a[index1][2]		//swap index.prev
	a[a[index1][1]][2], a[a[index2][1]][2] = a[a[index2][1]][2], a[a[index1][1]][2] //swap index.next.prev
	a[index1][1], a[index2][1] = a[index2][1], a[index1][1]		//swap index.next
	a[index1][0], a[index2][0] = a[index2][0], a[index1][0]		//swap index.key*/
}

func insert(a [][3]int, k int)  {
	pos := allocate(a)  //find the allocated index
	a[pos][0] = k 		//setting key
	a[pos][1] = listhead //pos.next = head.next
	a[pos][2] = 0 		//pos.prev = nil
	a[listhead][2] = pos //head.next.prev = pos
	listhead = pos 		//head.next = pos
}

func allocate(a [][3]int) int {
	if free == 0 { //the list is empty
		fmt.Println("out of the space")
		return 0
	}
	x := free
	free = a[x][1] //free = x.next
	return x //returns index x
}

func delete(a [][3]int, pos int)  {
	a[a[pos][2]][1] = a[pos][1] 	//pos.prev.next = pos.next
	a[a[pos][1]][2] = a[pos][2]		//pos.next.prev = pos.prev
	a[pos][0], a[pos][2] = 0, 0
}

func freeObeject(a [][3]int, x int)  {
	delete(a,x)
	a[x][1] = free //x.next = free
	free = x
}