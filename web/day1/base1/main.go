package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello",helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "URL.path=%q", req.URL.Path)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {

		fmt.Fprint(w, "Header[%q]=%q\n", k, v)
	}
}
