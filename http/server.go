package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello \n")
}

func header(w http.ResponseWriter, req *http.Request) {
	for name, header := range req.Header {
		for _, h := range header {
			fmt.Fprintf(w, "%v:%v \n", name, h)
		}
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", header)

	http.ListenAndServe("127.0.0.1:7000", nil)

}
