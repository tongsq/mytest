package main

import (
	"log"
	"os"
)

func main() {
	env := os.Environ()
	envCmd := []string{}
	for _, e := range env {
		envCmd = append(envCmd, e)
	}
	envCmd = append(envCmd, "name1=aaa")
	log.Printf("%#v\n", envCmd)
}
