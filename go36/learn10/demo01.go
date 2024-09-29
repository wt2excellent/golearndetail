package main

import (
	"fmt"
)

func main() {
	// 使用make声明并初始化一个int型的带缓冲的通道，并将其容量设置为3
	ch1 := make(chan int, 3)
	// 使用make声明并初始化一个int型的不带缓冲的通道，其容量为0
	ch2 := make(chan int)
	// 声明一个int型的通道
	var ch3 chan int
	// 使用接送操作符 <- 向通道ch1中发送int型数据1
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()
	// 使用接送操作符 <- 从通道ch1中读取数据，下面的短变量表达式左边有两个变量num, ok:其中ok用于判断通道ch1是否关闭，
	// 当ok == ture时表示通道没有关闭，可以读取数据并将其保存到变量nums中，当ok == false时表示通道关闭，不能读取数据。
	num, ok := <-ch1
	if ok {
		fmt.Println(num)
	}
	close(ch1)
	close(ch2)
	close(ch3)
}
