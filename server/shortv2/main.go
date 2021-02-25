package shortv2

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", Hello)

	server := http.Server{
		Handler:      http.DefaultServeMux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	listen, err := net.Listen("tcp4", server.Addr)
	if err != nil {
		panic(err)
	}
	/*
		默认长连接。当下面两种情况可以设置成短连接：
		1. 服务端资源受限
		2. 服务端将要关闭
	*/
	server.SetKeepAlivesEnabled(false)
	_ = server.Serve(listen)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "hello")
}
