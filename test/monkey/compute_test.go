// ------------------------------------------------------------
// package monkey
// @file      : compute_test.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/10/4 15:24
// ------------------------------------------------------------
package monkey

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"reflect"
	"testing"
)

// 对函数进行mock
func Test_compute(t *testing.T) {
	patches := gomonkey.ApplyFunc(networkCompute, func(a, b int) (int, error) {
		return 8, nil
	})
	defer patches.Reset()
	sum, err := compute(1, 2)
	if err != nil {
		t.Error("Error occurred:", err)
	}
	fmt.Printf("sum : %d\n", sum)
	if sum != 8 {
		t.Error("Error occurred in sum:", err)
	}
}

// 对结构体中的方法进行mock
func Test_Compute_NetworkCompute(t *testing.T) {
	var compute *Compute
	patches := gomonkey.ApplyMethod(reflect.TypeOf(compute), "NetworkCompute", func(_ *Compute, a, b int) (int, error) {
		return 10, nil
	})
	defer patches.Reset()
	compute = &Compute{}
	sum, err := compute.Compute(3, 2)
	if err != nil {
		t.Error("Error occurred:", err)
	}
	fmt.Printf("sum : %d\n", sum)
	if sum != 10 {
		t.Error("Error occurred in sum:", err)
	}
}

var num = 5

// 对变量进行mock
func Test_GlobalVal(t *testing.T) {
	patches := gomonkey.ApplyGlobalVar(&num, 10)
	defer patches.Reset()
	if num != 10 {
		t.Error("Error occurred in num:mock failure", num)
	}
}
