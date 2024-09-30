// ------------------------------------------------------------
// package main
// @file      : utils.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/9/30 10:21
// ------------------------------------------------------------
package main

import "math/rand"

func requestVote() bool {
	intn := rand.Intn(10)
	if intn <= 5 {
		return true
	}
	return false
}
