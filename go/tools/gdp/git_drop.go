package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const (
	INDENT = "  "
)

var basePath string

func init() {
	flag.StringVar(&basePath, "p", "", "path you want to show")
}

func gitDrop() error {
	err := execCommand("git", "add", ".")
	if err != nil {
		return err
	}
	err = execCommand("git", "stash")
	if err != nil {
		return err
	}
	err = execCommand("git", "stash", "drop")
	if err != nil {
		return err
	}
	return nil
}

func execCommand(name string, options ...string) error {
	c := exec.Command(name, options...)
	c.Dir = basePath
	var result []byte
	var err error
	if result, err = c.Output(); err != nil {
		fmt.Println("exec fail", err)
		return err
	}
	fmt.Print(string(result))
	return nil
}
func main() {
	flag.Parse()
	if basePath == "" {
		defaultPath, err := os.Getwd()
		if err != nil {
			fmt.Println("GetWd Error:", err)
			return
		}
		basePath = defaultPath
	}
	fmt.Printf("%s : \n", basePath)
	err := gitDrop()
	if err != nil {
		fmt.Println("lss error: ", err)
	}
}
