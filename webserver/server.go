package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "pong")
	})

	handler.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		vars := request.URL.Query()
		fmt.Println(vars)
		writer.Write([]byte("OK"))
	})

	handler.HandleFunc("/post", func(writer http.ResponseWriter, request *http.Request) {
		body := struct {
			Name string
			Age  int
		}{}
		bs, _ := ioutil.ReadAll(request.Body)
		json.Unmarshal(bs, &body)
		fmt.Println(body)
		writer.Write([]byte("OK"))
	})

	handler.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		js, _ := json.Marshal(SUCCESS)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(js)
	})

	handler.HandleFunc("/xml", func(writer http.ResponseWriter, request *http.Request) {
		js, _ := xml.MarshalIndent(SUCCESS, "", "  ")
		writer.Header().Set("Content-Type", "application/xml")
		writer.Write(js)
	})

	handler.HandleFunc("/download", func(writer http.ResponseWriter, request *http.Request) {
		fp := filepath.Join("http_server", "images", "image.jpg")
		fn := filepath.Base(fp)
		// 返回的文件直接下载
		writer.Header().Set("Content-Disposition", "attachment; filename="+fn)
		// 没有设置Content-Disposition，默认在浏览器展示，展示不了的按照二进制下载
		http.ServeFile(writer, request, fp)
	})

	handler.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		fp := filepath.Join("http_server", "templates", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx := map[string]interface{}{
			"Name": "cm",
		}

		//ctx := struct {
		//	Name string
		//}{
		//	Name: "cm",
		//}

		if err := tmpl.Execute(writer, ctx); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	_ = http.ListenAndServe(":20000", handler)
}
