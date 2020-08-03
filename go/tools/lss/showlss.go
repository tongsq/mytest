package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	INDENT = "  "
)

var basePath string

func init() {
	flag.StringVar(&basePath, "p", "", "path you want to show")
}
func showDirs(bashPath string, prefix string, showAll bool) error {
	base, err := os.Open(bashPath)
	if err != nil {
		return err
	}
	items, err := base.Readdir(-1)
	if err != nil {
		return err
	}
	var dirList []string
	var dirNameList []string
	for _, item := range items {
		fileInfo := item.(os.FileInfo)
		if fileInfo.IsDir() {
			dirPath := fileInfo.Name()
			dirList = append(dirList, path.Join(bashPath, dirPath))
			dirNameList = append(dirNameList, dirPath)
			continue
		}
		if strings.HasPrefix(fileInfo.Name(), ".") && !showAll {
			continue
		}
		fmt.Printf("%s\n", prefix+fileInfo.Name())
	}
	for k, dirItem := range dirList {
		fmt.Printf("%s\\\n", prefix+dirNameList[k])
		err := showDirs(dirItem, INDENT+prefix, true)
		if err != nil {
			return err
		}
	}
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
	err := showDirs(basePath, INDENT, true)
	if err != nil {
		fmt.Println("lss error: ", err)
	}
}
