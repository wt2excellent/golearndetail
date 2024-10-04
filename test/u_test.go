// ------------------------------------------------------------
// package test
// @file      : u_test.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/10/4 11:56
// ------------------------------------------------------------
package test

import (
	"fmt"
	"testing"
)

// 其他的测试框架，测试期望，更复杂的期望，当我没法去调用gorm进行查询的时候

// 测试函数名称必须要大写
func TestName(t *testing.T) {
	fmt.Println("my first test.")
	//t.Error("my first error.")
}
