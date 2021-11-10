package main

import (
	"cmdbgo/control"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Index
func index(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(writer, nil)
}

// 处理GET请求
func handleGet(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	fmt.Println(request.Method)

	// 第一种方式
	// id := query["id"][0]

	// 第二种方式
	id := query.Get("id")

	fmt.Printf("GET: id=%s\n", id)

	fmt.Fprintf(writer, `{"code":0}`)
}

// Main function
func main() {
	// Static dir
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Restful API
	http.HandleFunc("/", index)
	http.HandleFunc("/model", control.Model)
	fmt.Println("Running at port 3000 ...")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
