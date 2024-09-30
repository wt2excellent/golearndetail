// ------------------------------------------------------------
// package main
// @file      : voteselect.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/9/30 09:36
// ------------------------------------------------------------
package main

func main() {

	count := 0
	finished := 0

	for i := 0; i < 10; i++ {
		// 匿名函数，创建共 10 个线程
		go func() {
			vote := requestVote() // 一个内部sleep随机时间，最后返回true的函数，模拟投票
			if vote {
				count++
			}
			finished++
		}()
	}
	for count < 5 && finished != 10 {
		// wait
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
}
