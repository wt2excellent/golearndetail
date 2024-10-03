// ------------------------------------------------------------
// package main
// @file      : atomicValue.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/10/2 21:44
// ------------------------------------------------------------
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// 创建atomic.Value类型的值
	box := atomic.Value{}
	// 错误示例：构建一个切片并直接存储
	slice := []int{1, 2, 3}
	box.Store(slice)
	fmt.Println(box.Load().([]int))
	fmt.Println(box.Load().([]int))

	// 正确的存储，由于切片是引用类型，使用方式一进行存储的话在另一个goroutine中也可以改变存储的底层的slice的值，使用下述方法
	// 直接创建一个本地的切片，并复制原有切片，指针和外界没有关系。
	store := func(v []int) {
		replica := make([]int, len(v))
		copy(replica, v)
		box.Store(replica)
	}
	store(slice)
	fmt.Println(box.Load().([]int))

	// 调用atomic.Value的Swap方法
	oldSlice := (box.Swap([]int{3, 5})).([]int)
	fmt.Printf("old slice: %v\n", oldSlice)
	fmt.Printf("new slice:  %v\n", box.Load().([]int))
	// 调用atomic.Value的CAS方法，下面的方法是实现不了的，由于切片是引用类型不支持比较
	box.CompareAndSwap([]int{3, 5}, []int{9, 10})
	fmt.Println(box.Load().([]int))
}
