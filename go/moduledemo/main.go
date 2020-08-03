package main

import (
	"fmt"
	"moduledemo/mypackage"
	"moduledemo/util"
)

func main() {
	fmt.Println("main")
	mypackage.Show()
	mypackage.UtilShow()
	fmt.Println(util.Sum(1, 2, 3, 4))
}
