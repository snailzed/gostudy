package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	addr := ":8080"
	http.Handle("/", gethHandler(indexHandler))
	http.Handle("/temp", gethHandler(templateHanlder))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//
func gethHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("error:", err)
			}
		}()
		handler(writer, request)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello")
	if err != nil {
		log.Println("write error:", err)
	}
}

func templateHanlder(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	t := template.Must(template.ParseFiles("./index.html"))
	err := t.Execute(w, data)
	if err != nil {
		fmt.Println("temp error")
	}
}
