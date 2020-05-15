/*
We place the edge weight attribute in the node of the adjacency list.
The graph is represented by an adjacency list.
Bellman-Ford algorithm returns a boolean value indicating whether or not there is
a negative-weight cycle that is reachable from the source.
 */
package main

import (
	"fmt"
	"math"
)

type node struct {
	number int
	d int 	// an upper bound on the weight of a shortest path from source to the node
	w int
	pre *node	// the node's predecessor
	next *node 	// the node's adjacency node
}

func main()  {
	list := []node{}
	list = adjacencyListInitial(list)
	BellmanFord(list,&list[0])
	fmt.Println(list)
}

func BellmanFord(list []node, s *node) bool { // T(n) = O(VE)
	initializeSingleSource(list,s)
	for i := 0; i < len(list) - 1; i++ {
		for j := 0; j < len(list); j++ { // for each edge(u,v) belongs to G.E
			for k := list[j].next; k != nil; k = k.next {
				relax(&list[j],&list[k.number],k.w)
			}
		}
	}
	for j := 0; j < len(list); j++ { // check for a negative-weight cycle and return the appropriate boolean value.
		for k := list[j].next; k != nil; k = k.next {
			if list[j].d > list[k.number].d + k.w {
				return false
			}
		}
	}
	return true
}

func initializeSingleSource(list []node, s *node)  {
	for i := 0; i < len(list); i++ {
		list[i].d = math.MaxInt32
	}
	s.d = 0
}

func relax(u, v *node, weight int)  {
	if v.d > u.d + weight {
		v.d = u.d + weight
		v.pre = u
	}
}

func w(u, v *node) int { // search u's adjacency list for the edge weight of u and v
	for i := u.next; i != nil; i = i.next {
		if i.number == v.number {
			return i.w
		}
	}
	return 0
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