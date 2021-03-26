package main

import (
	"fmt"
	"strconv"
)

func main() {
	//为了说明类型,我采用了显性的变量定义方法,实际开发中更多的是用“:=”自动获取类型变量类型
	var mystr string = "Hello!"
	var mystrP *string = &mystr

	fmt.Println(mystrP, *mystrP)
	fmt.Println(strconv.ParseInt("0xFFFF", 0, 0))
}
