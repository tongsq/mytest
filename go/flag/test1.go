package main

import (
	"flag"
	"fmt"
)

var name string
var quit bool

func init() {
	flag.StringVar(&name, "Name", "Tom", "input Name")
	flag.BoolVar(&quit, "q", false, "is Quit")
}
func main() {
	flag.Parse()
	if quit {
		return
	}
	fmt.Println(name)
	fmt.Printf("%#v\n", flag.Lookup("Name"))
	fmt.Println(flag.Lookup("Name").Value)
}
