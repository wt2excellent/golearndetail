// ------------------------------------------------------------
// package main
// @file      : voteselect_channel.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/9/30 10:36
// ------------------------------------------------------------
package main

func main() {

	count := 0
	finished := 0
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		// 匿名函数，创建共 10 个线程
		go func() {
			ch <- requestVote()
		}()
	}
	// 这里实现并不完美，如果count >= 5了，主线程不会再监听channel，导致其他还在运行的子线程会阻塞在往channel写数据的步骤。
	// 但是这里主线程退出后子线程们也会被销毁，影响不大。但如果是在一个长期运行的大型工程中，这里就存在泄露线程leaking threads
	for count < 5 && finished != 10 {
		// 主线程在这里等待
		v := <-ch
		if v {
			count++
		}
		finished++
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
}
