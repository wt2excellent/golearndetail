// ------------------------------------------------------------
// package main
// @file      : saletickets.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/9/30 11:21
// ------------------------------------------------------------
package main

import (
	"fmt"
	"time"
)

var (
	saleTickets = 100
	currentNum  = 0
)

func main() {
	for i := range 10 {
		go saleTicket(i)
	}
	time.Sleep(5 * time.Second)
}

// 售票函数
func saleTicket(i int) {
	for currentNum < saleTickets {
		fmt.Printf("current goroutine is : %v, is slaing ticket num is %v\n", i, currentNum)
		currentNum += 1
	}
}
