package main

// a demo follow the tutorial: https://studygolang.com/articles/9467

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"go-http/handlers"
)

func main() {
	fmt.Println("Hello World!")

	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/foo", rh)
	
	mux.HandleFunc("/time", handlers.NewTimeHandlerFunc(time.RFC1123))
	mux.HandleFunc("/time/rfc1123", handlers.NewTimeHandlerFunc(time.RFC1123))
	mux.HandleFunc("/time/rfc3339", handlers.NewTimeHandlerFunc(time.RFC3339))

	log.Println("Listening...")
	
	http.ListenAndServe(":3000", mux)
}