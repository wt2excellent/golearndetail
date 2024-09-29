package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 类型转换范围限定演练
	srcNum := int16(-255)
	dstNum := int8(srcNum)
	fmt.Printf("srcNum:%d, dstNum:%d\n", srcNum, dstNum)
	/**
	Go语言中负数以补码的形式存在，补码：源码求反+1
	-255 ：1111111100000001
	从16位转为8位，需要高位截断变为00000001，由于最高位是0所以表示正数，所以是1
	*/

	/**
	Go中有效的Unicode代码点是什么？？？
	在将一个整数值转换为字符串时，这个整数应该是一个有效的 Unicode 代码点的值。Unicode 是一个字符编码标准，它为世界上大多数的文字系统提供了一个唯一的码位。每个 Unicode 代码点都对应一个字符。
	在计算机中，字符通常以字节的形式存储，而每个字节可以表示 0 到 255 之间的整数值。当一个整数超出了这个范围，或者它不是一个有效的 Unicode 代码点时，尝试将它转换为字符串可能会导致无法正确显示该字符，从而出现替代字符，通常是 ""（一个黑色菱形，中间有一个问号）。
	例如，在 UTF-8 编码中，一个字符可能由一到四个字节表示。如果一个整数对应于一个超出常用 Unicode 字符范围的值（比如大于 0x10FFFF），或者它是一个用于表示字符属性的码点（比如一些控制字符），那么它可能无法被正确地转换为一个可打印的字符。
	在 Go 语言中，如果你使用 string() 函数将一个整数值转换为字符串，并且该值超出了有效的 Unicode 代码点范围，你可能会得到一个替代字符。为了避免这个问题，你应该确保转换的整数值在有效的 Unicode 范围内（通常是 0 到 0x10FFFF）。
	*/
	fmt.Println(string(65))
	fmt.Println(string(37))
	fmt.Println(string(-1))

	/**
	正确的将整数转为字符串应该使用Go中提供的转换工具
	如：strconv.Itoa()\strconv.FormatInt()
	*/

	num := -1
	fmt.Println(strconv.Itoa(num))
	fmt.Println(strconv.FormatInt(int64(num), 10))
}
