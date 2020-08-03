package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func readFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return ioutil.ReadAll(file)
}

func deferIt2() {
    for i := 1; i < 5; i++ {
        defer fmt.Print(i)
    }
}

func deferIt3() {
    f := func(i int) int {
        fmt.Printf("%d ",i)
        return i * 10
    }
    for i := 1; i < 5; i++ {
        defer fmt.Printf("%d ", f(i))
    }
}

func main() {
	res,_ := readFile("/tmp/aa.txt")
	str := string(res)
	fmt.Println(str)
	deferIt2()
	deferIt3()
}