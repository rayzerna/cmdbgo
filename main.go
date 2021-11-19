package main

import (
	"cmdbgo/control"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/urfave/negroni"
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

	// Auth API
	http.HandleFunc("/api/login", control.LoginHandler)
	http.HandleFunc("/api/registry", control.SighupHandler)

	// Restful API
	http.Handle("/api/model", negroni.New(
		negroni.HandlerFunc(control.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(control.Model)),
	))
	http.Handle("/api/item", negroni.New(
		negroni.HandlerFunc(control.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(control.Item)),
	))

	fmt.Println("Running at port 3000 ...")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
