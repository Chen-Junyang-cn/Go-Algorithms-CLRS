/*
list[u] consists of all the vertices adjacent to u in graph.
The adjacent list includes the property weight.
 */
package main

import (
	"fmt"
	"math"
)

type node struct {
	number int	 //编号指向节点的位置
	key int
	next *node
	w int 		 //权值
	color string //节点的颜色
	pre *node 	 //节点的祖先
}

func main()  {
	list := []node{}
	list = adjacencyListInitial(list)
	MSTPrim(list)
	/*queue := make([]node,len(list))
	copy(queue,list)
	for i := 2; i >= 0; i-- {
		queue,min := heapExtractMin(queue,i)
		fmt.Println(min,queue,len(queue))
	}*/
	fmt.Println(list)
}

func MSTPrim(list []node)  { 	// The root node as the "search" node
	for i := 0; i < len(list); i++ {
		list[i].key = 20 + i 			// let the key different but bigger than the max weight
		list[i].color = "white"
	}
	list[0].key = 0
	queue := make([]node,len(list))
	copy(queue,list)
	cnt := 0 			// calculate the elements that are out of the queue
	for len(queue) != 0 {
		var u *node
		queue, u = heapExtractMin(queue,len(queue)-1)
		fmt.Println(queue)
		cnt++
		list[u.number].color = "black" 			// the node is out of the queue
		for v := u.next; v != nil; v = v.next { // for each v belongs to adj[u]
			if list[v.number].color == "white" && v.w < list[v.number].key {
				list[v.number].pre = &list[u.number]
				list[v.number].key = v.w
				queue[v.number-cnt].key = v.w // make elements in different data structures one to one
			}
		}
	}
}

func adjacencyListInitial(list []node) []node {
	var temp int
	fmt.Println("你要输入的邻接表的顶点个数：")
	fmt.Scan(&temp)
	for i := 0; i < temp; i++ {
		head := new(node)		//初始化邻接表元素，并把这个元素当作邻接表头
		head.number = i 	 	//这个节点的编号
		list = append(list,*head)
		fmt.Println("你想输入邻接表的边的个数")
		var j int
		fmt.Scan(&j)
		for k := 0; k < j; k++ {
			fmt.Println("输入邻接的元素（头插法）和这条边的权值")
			insert(&list[i])
		}
	}
	return list
}

func insert(head *node)  { //头插法
	x := new(node)
	fmt.Scan(&x.number,&x.w)
	x.next = head.next
	head.next = x
}

func buildMinHeap(a []node)  { 					//build the heap. two subtrees of the node are both max-heaps.
	for i := (len(a))/2-1; i >= 0; i-- {
		minHeapify(a,i,len(a)-1)
	}
}

func heapExtractMin(a []node, heapSize int) ([]node,*node) { //T(n)=O(lgn)
	if heapSize < 0 {
		fmt.Errorf("heap underflow")
		return nil,nil
	}
	min := a[0]
	a[0], a = a[heapSize], a[0:heapSize]
	heapSize--
	minHeapify(a,0,heapSize)
	return a,&min
}

func minHeapify(a []node,i, heapSize int)  {
	l, r, min := left(i), right(i), math.MaxInt32
	/*if l > heapSize || r > heapSize {
		return
	}*/
	if l <= heapSize {
		if a[l].key > a[i].key {
			min = i
		} else {
			min = l
		}
		fmt.Println(min)
	}
	if r <= heapSize {
		if a[r].key < a[min].key {
			min = r
		}
	}
	/*if l <= heapSize && a[l].key > a[i].key {
		min = i
	} else {
		min = l
	}
	if r <= heapSize && a[r].key < a[min].key {
		min = r
	}*/
	if min != i && min <= heapSize {
		a[i], a[min] = a[min], a[i]
		minHeapify(a,min,heapSize)
	}
}

func left(i int) int {
	return 2*i+1
}

func right(i int) int {
	return 2*i+2
}

func parent(i int) int {
	return (i-1)/2
}