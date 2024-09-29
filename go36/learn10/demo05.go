package main

import (
	"fmt"
	"math/rand"
)

/*
*
channel和select的联合使用
*/
func main() {
	// 创建一个多维通道
	chs := []chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 创建随机数
	index := rand.Intn(3)
	fmt.Printf("index is: %v\n", index)
	chs[index] <- 1
	select {
	case elem := <-chs[0]:
		fmt.Printf("通道0被选中,元素为：%v\n", elem)
	case elem := <-chs[1]:

		fmt.Printf("通道1被选中,元素为：%v\n", elem)
	case elem := <-chs[2]:
		fmt.Printf("通道2被选中,元素为：%v\n", elem)
	default:
		fmt.Println("error.")
	}
}
