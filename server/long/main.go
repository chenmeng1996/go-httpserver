package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", Hello)
	_ = http.ListenAndServe(":7777", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "hello")
	time.Sleep(4 * time.Second)
	_, _ = fmt.Fprint(w, "hello again")
}
