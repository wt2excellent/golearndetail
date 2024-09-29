// ------------------------------------------------------------
// package main
// @file      : demo03.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/9/29 23:00
// ------------------------------------------------------------
package main

func test01() {
	l := 1
	a := make([]int, l, l)
	for i := 0; i < l; i++ {
		a[i] = i
	}
}

func main() {
	test01()
}
