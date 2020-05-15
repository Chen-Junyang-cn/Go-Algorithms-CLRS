/*
Using adjacency-list representation of graph to implement BFS.
list[u] consists of all the vertices adjacent to u in graph.
 */
package main

import "fmt"

type node struct {
	key int
	next *node
	d int 		 //和根节点距离
	color string //节点的颜色
	pre *node 	 //节点的祖先
}

type queue struct { //由一组灰色顶点组成，代表即将遍历的节点
	array [5]int   //最大可缓存5个节点
	head, tail int
}

func main()  {
	list := []node{}
	list = adjacencyListInitial(list)
	q := queue{
		array: [5]int{},
		head:  0,
		tail:  0,
	}
	BFS(list,&list[0],&q)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i],list[i].next,list[i].next.next)
	}
	printPath(list,&list[0],&list[4])
}

func BFS(list []node, s *node, q *queue)  { // s是根节点（源顶点），即从s节点开始遍历
	for i := 0; i < len(list); i++ {
		list[i].color = "white" //将颜色调为白色，其祖先和键值在创建的时候已经赋值nil和0
	}
	s.color = "gray" // s已经遍历过
	q.head, q.tail = 0, 0  // 将队列置空
	enqueueTail(q,s.key)
	for !empty(q) { // 队列非空
		u := dequeueHead(q)
		for v := list[u].next; v != nil; v = v.next { // 对与u相邻的节点v，这个节点v实际作用是指向每条链的头节点
			if list[v.key].color == "white" {	//v.key指向的是那个节点,所有的操作应该在list所含的节点上
				list[v.key].color = "gray"  	//节点颜色变灰色
				list[v.key].d = list[u].d + 1
				list[v.key].pre = &list[u]
				fmt.Println("进栈的节点：",v.key)
				enqueueTail(q,v.key)
			}
		}
		list[u].color = "black" //邻近u的节点已经被发现了
	}
}

func printPath(list []node, s, v *node)  {
	if v == s {
		fmt.Println(s)
	} else if v.pre == nil {
		fmt.Println("no path from",s,"to",v,"exists")
	} else {
		printPath(list,s,v.pre)
		fmt.Println(v)
	}
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

func enqueueTail(q *queue, key int)  {
	if full(q) {
		return
	}
	q.array[q.tail] = key
	if q.tail == len(q.array) - 1 {
		q.tail = 0
	} else {
		q.tail++
	}
}

func dequeueHead(q *queue) int {
	if empty(q) {
		return 0
	}
	key := q.array[q.head]
	if q.head == len(q.array) - 1 {
		q.head = 0
	} else {
		q.head++
	}
	return key
}

func full(q *queue) bool {
	if q.head == q.tail + 1 {
		return true
	}
	return false
}

func empty(q *queue) bool {
	if q.head == q.tail {
		return true
	}
	return false
}