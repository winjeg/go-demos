package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Usage:\n./proxy remote_addr:remote_port  local_port")
	listenAddr := fmt.Sprintf(":%s", os.Args[2])
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s", &method)

	proxyUrl := os.Args[1]
	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", proxyUrl)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(b[:n])
	}
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}
