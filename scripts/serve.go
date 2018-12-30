// +build ignore

package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "a", "localhost:8080", "address for server.")
}

func main() {
	flag.Parse()
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("."))))
}
