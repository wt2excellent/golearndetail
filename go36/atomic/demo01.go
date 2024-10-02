// ------------------------------------------------------------
// package atomic
// @file      : demo01.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/10/2 20:44
// ------------------------------------------------------------
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	a := int32(0)
	a = atomic.AddInt32(&a, int32(23))
	fmt.Println(a)
	swapBool := atomic.CompareAndSwapInt32(&a, int32(23), int32(45))
	if swapBool {
		fmt.Printf("a swap val is : %v\n", a)
	}

	// 验证，通过AddUint32执行原子减法
	fmt.Println("===========================")
	b := uint32(18)
	delat := int32(-3)
	b = atomic.AddUint32(&b, uint32(delat))
	fmt.Println("b sub 3 result is : ", b)
	// 计算机系统中补码表示为：源码 求反 + 1.
	// ^uint32(-(-3) - 1)表示的补码与int32(-3)表示的补码相同的
	b = atomic.AddUint32(&b, ^uint32(-(-3)-1))
	fmt.Println("b sub 3 result is : ", b)
}
