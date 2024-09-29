package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
使用sync.Mutex解决多协程同时对共享数据进行读写操作，避免竟态条件
*/
func main() {
	m := map[int]string{
		1: "haha",
	}
	var mutex sync.Mutex // 创建一个互斥锁

	// 启动读协程
	go read(m, &mutex)

	// 等待一秒钟，确保读协程已经开始运行
	time.Sleep(time.Second)

	// 启动写协程
	go write(m, &mutex)

	// 等待足够长的时间，以便读写协程可以运行
	time.Sleep(30 * time.Second)

	// 打印最终的map
	fmt.Println(m)
}

func read(m map[int]string, mutex *sync.Mutex) {
	for {
		mutex.Lock() // 在读取之前加锁
		value := m[1]
		mutex.Unlock() // 读取完毕后解锁
		fmt.Println("Read:", value)
		time.Sleep(1 * time.Second)
	}
}

func write(m map[int]string, mutex *sync.Mutex) {
	for {
		mutex.Lock() // 在写入之前加锁
		m[1] = "write"
		mutex.Unlock() // 写入完毕后解锁
		fmt.Println("Write:", m[1])
		time.Sleep(1 * time.Second)
	}
}
