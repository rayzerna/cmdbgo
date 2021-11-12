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

// Main function
func main() {
	// Static dir
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Static page
	http.HandleFunc("/", index)

	// Restful API
	http.HandleFunc("/api/model", control.Model)
	http.HandleFunc("/api/registry", control.Sighup)
	fmt.Println("Running at port 3000 ...")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
