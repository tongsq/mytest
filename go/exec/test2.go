package main

import (
	"bytes"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("ls", "-l", "/var/log")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdout = os.Stdout
	_ = cmd.Start()
	_ = cmd.Wait()
	time.Sleep(time.Second * 3)
	//err := cmd.Run()
	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}
}
