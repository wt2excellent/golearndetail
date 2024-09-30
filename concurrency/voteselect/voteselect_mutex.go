// ------------------------------------------------------------
// package main
// @file      : voteselect_mutex.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/9/30 10:14
// ------------------------------------------------------------
package main

import (
	"sync"
)

func main() {

	count := 0
	finished := 0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		// 匿名函数，创建共 10 个线程
		go func() {
			vote := requestVote() // 一个内部sleep随机时间，最后返回true的函数，模拟投票
			// 临界区加锁
			mu.Lock()
			// 推迟到基本block结束后执行，这里即函数执行结束后 自动执行解锁操作。利用defer语言，一般在声明加锁后，立即defer声明推迟解锁
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
		}()
	}
	for {
		mu.Lock()
		if count > 5 || finished == 10 {
			// 不能在此处解锁，下面仍然需要对count变量进行判断
			//mu.Unlock()
			break
		}
		mu.Unlock()
		// wait
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
	mu.Unlock()
}
