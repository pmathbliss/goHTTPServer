package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there, I love %s!</h1>", r.URL.Path[1:])
}

func main() {
	var portValue string
	flag.StringVar(&portValue, "port", "8081", "Port for server")
	flag.Parse()
	fmt.Printf("Start server on port %s", portValue)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+portValue, nil))
}
