package main

import (
	"fmt"
	"net/http"
	"wsjgin"
)

func main() {
	r := wsjgin.Default()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":3456")
}