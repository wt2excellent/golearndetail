package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "greeting object.")
}
func main() {
	flag.Parse()
	// fmt.Printf("Hello %s\n", name)
	hello(name)
}
