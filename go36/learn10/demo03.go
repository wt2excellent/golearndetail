package main

import "fmt"

func main() {
	sendChan := getIntChan()
	for elem := range sendChan {
		fmt.Println(elem)
	}
}

// 获取一个单向发送通道
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}
