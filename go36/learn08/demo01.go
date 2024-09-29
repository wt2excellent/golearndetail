package main

import (
	"container/list"
	"fmt"
)

/*
*
Go语言链表详解
*/
func main() {
	//Go语言链表使用
	l := list.New()
	ele4 := l.PushBack(4)
	ele1 := l.PushFront(1)
	l.InsertBefore(3, ele4)
	l.InsertAfter(2, ele1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 声明一个链表但未初始化，随后使用链表的延迟初始化技术在插入元素时进行初始化（lazyInit）
	l2 := list.List{}
	fmt.Println(l2)
	l2.PushBack(3)
	fmt.Println(l2)
}
