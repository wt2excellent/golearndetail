package main

/*
*
验证通道何时触发panic
*/
func main() {
	var ch chan int
	close(ch)

	ch1 := make(chan int)
	close(ch1)
	close(ch1)

}
