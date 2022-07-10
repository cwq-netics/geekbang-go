package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"week_9/protocol"
)

//2.实现一个从 socket connection 中解码出 goim 协议的解码器

var (
	proto protocol.Decoder
)

func main() {
	proto = protocol.NewGoImProto()
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("dial err: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("connect success")
	fmt.Printf("send msg\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		_, _ = conn.Write([]byte(text))
		read(conn)
	}
}

func read(conn net.Conn) {
	rr := bufio.NewReader(conn)
	req := proto.ReadTCP(rr)
	if req == nil {
		fmt.Printf("illegal req")
		return
	}
	result, err := req.Body()
	if err != nil {
		fmt.Printf("err when read: %v", err)
	}
	fmt.Println(result)
}
