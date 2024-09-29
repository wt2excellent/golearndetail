package main

//// 潜在类型
//type MyString string
//
//// 类型别名
//type MyString2 = string
//
//func main() {
//	var name MyString
//	name = "zhang san"
//	var copyName string
//	/**
//	潜在类型相同的不同类型的值之间是可以进行类型转换的。
//	即使两个类型的潜在类型相同，但这两个类型对应的变量之间也不能进行判等或者比较，也不能进行赋值，只能进行类型转换
//	copyName = name 不允许的操作
//	*/
//	copyName = string(name)
//	fmt.Println(name)
//	fmt.Println(copyName)
//	fmt.Println("=========================")
//	var name2 MyString2
//	name2 = "lisi"
//	var copyName2 string
//	copyName2 = name2
//	fmt.Println(name2)
//	fmt.Println(copyName2)
//}
