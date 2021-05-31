package main

import (
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	Address = "127.0.0.1:6379"
	Network = "tcp"
)

func Conn(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetRequest(args []string) []byte {
	req := []string{
		"*" + strconv.Itoa(len(args)),
	}

	for _, arg := range args {
		req = append(req, "$"+strconv.Itoa(len(arg)))
		req = append(req, arg)
	}

	str := strings.Join(req, "\r\n")
	return []byte(str + "\r\n")
}

func main() {
	//args := os.Args[1:]
	//if len(args) <= 0 {
	//	log.Fatalf("os.Args <= 0")
	//}

	redisConn, err := Conn(Network, Address)
	if err != nil {
		log.Fatalf("Conn error : %v", err)
	}
	defer redisConn.Close()

	w, err := redisConn.Write([]byte("*3\\r\\n$3\\r\\nSET\\r\\n$5\\r\\nmykey\\r\\n$7\\r\\nmyvalue\\r\\n"))
	if err != nil {
		log.Fatalf("Conn Write err : %v", err)
	}
	log.Println("write success", w)
	time.Sleep(time.Second * 3)
	command := make([]byte, 1024)
	n, err := redisConn.Read(command)
	if err != nil {

		log.Fatalf("read err : %v", err)
	}
	println("reply", n, command)
}
