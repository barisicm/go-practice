package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
	// TODO: DUMP WEB REQUEST
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hi there, i love %s!", r.URL.Path[1:])
}


