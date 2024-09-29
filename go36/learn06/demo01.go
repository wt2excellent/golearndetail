package main

import "fmt"

var container = []string{"Beijing", "Shanghai"}

func main() {
	fmt.Println(container)
	container = []string{"Hello", "Hi"}
	fmt.Printf("variable redeclare %s\n", container)
	container := map[int]string{1: "Beijing", 2: "Shanghai"}
	strings, ok := interface{}(container).([]string)
	if ok {
		fmt.Println("Container type is []string...")
		fmt.Println(strings)
	} else {
		fmt.Println("Container type is map...")
		fmt.Printf("strings is %v\n", strings)
	}
	fmt.Println(container)
}
