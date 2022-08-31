package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello tiam.com"))
}

func main() {
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8888",
	}
	fmt.Println("server startup...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
