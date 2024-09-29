package main

import (
	"flag"
	"testhello/go36/learn03/q2/lib"
	// "testhello/learn03/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "greeting object.")
}
func main() {
	flag.Parse()
	// fmt.Printf("Hello %s\n", name)
	// lib_demo02.Hello(name)
	lib_tews.Hello(name)
	// lib_my.Hello(name)
}
