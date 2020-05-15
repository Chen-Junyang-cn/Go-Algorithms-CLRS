/*
Define u.d as the discovery time and u.f as the finishing time.
Using adjacency-list representation of graph to implement DFS.
list[u] consists of all the vertices adjacent to u in graph.
 */
package main

import "fmt"

type node struct {
	key int		// using key as the vertex index in the link list
	color string
	d, f int 	// record the time
	next, pre *node
}

var time int	// resets the global time counter

func main()  {
	list := []node{}
	list = adjacencyListInitial(list)
	time = 0
	DFS(list)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i],list[i].next,list[i].next.next)
	}
}

func DFS(list []node)  {
	for i := 0; i < len(list); i++ { //initial
		list[i].color = "white"
	}
	for u := 0; u < len(list); u++ { 	// check each vertex in list
		if list[u].color == "white" {	// when a white vertex is found, visit it using DFSVisit
			DFSVisit(list,&list[u])
		}
	}
}

func DFSVisit(list []node, u *node)  {
	time++ 	//vertex u has been just discovered
	u.d = time
	u.color = "gray"
	for v := u.next; v != nil; v = v.next { // explore edge(u,v)
		if list[v.key].color == "white" {
			list[v.key].pre = u
			DFSVisit(list,&list[v.key])
		}
	}
	u.color = "black" 	// blacken u; it is finished
	time++
	u.f = time
}

func adjacencyListInitial(list []node) []node {
	var temp int
	fmt.Println("你要输入的邻接表的顶点个数：")
	fmt.Scan(&temp)
	for i := 0; i < temp; i++ {
		head := new(node)//初始化邻接表元素，并把这个元素当作邻接表头
		head.key = i 	 //键值为i，相当于这个节点的编号
		list = append(list,*head)
		fmt.Println("你想输入邻接表的边的个数")
		var j int
		fmt.Scan(&j)
		for k := 0; k < j; k++ {
			fmt.Println("输入邻接表的元素（头插法）")
			insert(&list[i])
		}
	}
	return list
}

func insert(head *node)  { //头插法
	x := new(node)
	fmt.Scan(&x.key)
	x.next = head.next
	head.next = x
}