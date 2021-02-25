package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Println(err)
	}

	// transport设置短连接
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	httpClient := &http.Client{Timeout: time.Duration(6) * time.Second, Transport: tr}
	resp, err := httpClient.Do(req)

	respBuf, err := ioutil.ReadAll(resp.Body)
	fmt.Println(len(respBuf))
}
