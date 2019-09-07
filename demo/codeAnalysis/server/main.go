package main

import (
	"io"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		io.WriteString(w, "HI")
	})
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		panic(err)
	}
}
