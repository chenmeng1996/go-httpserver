package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:6666")
	if err != nil {
		panic(err)
	}
	fd, err := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := fd.Accept()
		if err != nil {
			continue
		}
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	bs := make([]byte, 0)
	_, _ = conn.Read(bs)
	fmt.Println(string(bs))

	_, _ = conn.Write([]byte("hello"))
	_ = conn.Close()
}
