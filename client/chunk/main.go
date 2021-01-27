package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:8888")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	buf := make([]byte, 1024)
	var s string
	for {
		n, err := res.Body.Read(buf)
		if n == 0 || err != nil {
			// 等待数据
			continue
		}
		s += string(buf[:n])
		fmt.Print(s)
		s = ""
	}

}
