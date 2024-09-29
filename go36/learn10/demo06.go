package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		time.AfterFunc(time.Second*3, func() {
			close(ch)
		})
		// 方式一：使用break配合标签实现
	loop:
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					fmt.Println("The ch is closed.")
					break loop // 使用return配合标签退出for循环
				}
				fmt.Println("The ch case is selected.")
			}
		}
		fmt.Println("The end of ch.")
	}()

	go func() {
		ch1 := make(chan int, 3)
		ch1 <- 5
		ch1 <- 6
		ch1 <- 7
		time.AfterFunc(3*time.Second, func() {
			close(ch1)
		})

		for {
			select {
			case _, ok := <-ch1:
				if !ok {
					fmt.Println("The ch1 is closed.")
					goto loop1 // 使用goto配合标签实现退出select和for的结合使用
				}
				fmt.Println("The ch1 case is selected.")
			}
		}
	loop1:
		fmt.Println("The end of ch1.")
	}()
	// 等待10防止主协程退出后所有子协程死亡
	time.Sleep(time.Second * 10)
	fmt.Println("The end of main.")
}
