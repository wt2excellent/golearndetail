// ------------------------------------------------------------
// package main
// @file      : voteselect_mutex_cond.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/9/30 10:23
// ------------------------------------------------------------
package main

import "sync"

func main() {

	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	//cond := sync.Cond{L: &mu}
	for i := 0; i < 10; i++ {
		// 匿名函数，创建共 10 个线程
		go func() {
			vote := requestVote() // 一个内部sleep随机时间，最后返回true的函数，模拟投票
			// 临界区加锁
			mu.Lock()
			// Broadcast 唤醒所有等待条件变量 c 的 goroutine，无需锁保护；Signal 只唤醒任意 1 个等待条件变量 c 的 goroutine，无需锁保护。
			// 这里只有一个waiter，所以用Signal或者Broadcast都可以
			defer cond.Broadcast()
			// 推迟到基本block结束后执行，这里即函数执行结束后 自动执行解锁操作。利用defer语言，一般在声明加锁后，立即defer声明推迟解锁
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++

		}()
	}
	mu.Lock()
	for count < 5 || finished != 10 {
		// 如果条件不满足，则在制定的条件变量上wait。内部原子地进入sleep状态，并释放与条件变量关联的锁。当条件变量得到满足时，这里内部重新获取到条件变量关联的锁，函数返回。
		cond.Wait()
		// 使用cond.Wait的目的防止CPU空转，使用time.sleep()无法控制合适的休眠时间
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
	mu.Unlock()
}
