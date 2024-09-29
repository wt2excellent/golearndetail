package main

import "fmt"

/*
*数组和切片进阶用法
 */
func main() {
	slice1 := make([]int, 5)
	slice2 := make([]int, 5, 8)
	fmt.Printf("slice1 len is : %d, cap is : %d\n", len(slice1), cap(slice1))
	fmt.Printf("slice2 len is : %d, cap is : %d\n", len(slice2), cap(slice2))

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice4 := slice3[3:7]
	fmt.Printf("slice3 elements is : %d\n", slice3)
	fmt.Printf("slice4 elements is : %d\n", slice4)
	slice4 = append(slice4, 10)
	fmt.Printf("slice3 elements is : %d\n", slice3)
	fmt.Printf("slice4 elements is : %d\n", slice4)
	//fmt.Printf("slice3 len is : %d, cap is : %d\n", len(slice3), cap(slice3))
	//fmt.Printf("slice4 len is : %d, cap is : %d\n", len(slice4), cap(slice4))
}
