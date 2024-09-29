// Package mallocStackOrHeap ------------------------------------------------------------
// package mallocStackOrHeap
// @file      : demo02.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/9/29 22:55
// ------------------------------------------------------------
package main

func test() *int {
	a := 10
	return &a
}

func main() {
	_ = test()
}
