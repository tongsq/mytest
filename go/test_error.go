package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"bytes"
)

func main() {
	file, err := os.Open("/tmp/aa.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	br := bufio.NewReader(file)
	var buf bytes.Buffer
	for {
		ba, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error: %s\n", err)
			break
		}
		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}
	}
	str := buf.String()
	fmt.Println(str)
}