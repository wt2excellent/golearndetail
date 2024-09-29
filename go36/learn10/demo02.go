package main

import "fmt"

/*
*
创建单向通道
*/
func main() {
	// 定义带缓冲通道
	ch := make(chan int, 3)
	producer(ch, 7)
	ans := consumer(ch)
	fmt.Printf("main function get num is %v\n", ans)
}
func producer(ch chan<- int, num int) {
	ch <- num
}
func consumer(ch <-chan int) (ans int) {
	ans, ok := <-ch
	if ok {
		fmt.Printf("consumer get data is %v\n", ans)
		return ans
	}
	return 0
}
