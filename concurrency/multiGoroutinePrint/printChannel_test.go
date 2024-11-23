// ------------------------------------------------------------
// package multiGoroutinePrint
// @file      : printChannel_test.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/11/23 11:11
// ------------------------------------------------------------
package multiGoroutinePrint

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestMultiGoroutinePrintWithChan(t *testing.T) {
	// 创建无缓冲通道用于通信
	channelA := make(chan struct{})
	channelB := make(chan struct{})
	channelC := make(chan struct{})

	// 启动协程 A
	go func() {
		for {
			<-channelA // 等待通道A的信号
			fmt.Printf("A_1\n")
			channelB <- struct{}{} // 发送信号到通道B
		}
	}()

	// 启动协程 B
	go func() {
		for {
			<-channelB // 等待通道B的信号
			fmt.Printf("B_2\n")
			channelC <- struct{}{} // 发送信号到通道C
		}
	}()

	// 启动协程 C
	go func() {
		for {
			<-channelC // 等待通道C的信号
			fmt.Printf("C_3\n")
			channelA <- struct{}{} // 发送信号到通道A
		}
	}()

	// 开始打印，从A启动
	channelA <- struct{}{}

	// 让主线程阻塞，避免程序退出
	select {}
}

func TestMultiGoroutinePrintWithCond(t *testing.T) {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	current := 1 // 当前打印的线程标志

	go func() {
		for {
			mu.Lock()
			for current != 1 {
				cond.Wait() // 等待条件满足
			}
			fmt.Println("A_1")
			current = 2
			cond.Broadcast() // 唤醒其他协程
			mu.Unlock()
		}
	}()

	go func() {
		for {
			mu.Lock()
			for current != 2 {
				cond.Wait()
			}
			fmt.Println("B_2")
			current = 3
			cond.Broadcast()
			mu.Unlock()
		}
	}()

	go func() {
		for {
			mu.Lock()
			for current != 3 {
				cond.Wait()
			}
			fmt.Println("C_3")
			current = 1
			cond.Broadcast()
			mu.Unlock()
		}
	}()

	// 防止主线程退出
	select {}
}

func TestMultiGoroutinePrintWithCAS(t *testing.T) {
	a := int32(0)

	go func() {
		for {
			for atomic.LoadInt32(&a) != 0 {
				//time.Sleep(time.Millisecond * 1000)
			}
			fmt.Printf("A_1\n")
			atomic.StoreInt32(&a, 1)
		}
	}()
	go func() {
		for {
			for atomic.LoadInt32(&a) != 1 {
				//time.Sleep(time.Millisecond * 1000)
			}
			fmt.Printf("B_2\n")
			atomic.StoreInt32(&a, 2)
		}
	}()
	go func() {
		for {
			for atomic.LoadInt32(&a) != 2 {
				//time.Sleep(time.Millisecond * 1000)
			}
			fmt.Printf("C_3\n")
			atomic.StoreInt32(&a, 0)
		}
	}()

	select {}
}
