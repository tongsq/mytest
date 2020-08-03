package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.c5game.com/api/product/sale.json?id=2705689&page=1&sort=1&key=1523539522")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response error ", err)
	}
	fmt.Println(string(body))
}
