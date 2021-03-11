package main

// a demo follow the tutorial: https://studygolang.com/articles/9467

import (
	"fmt"
	"go-http/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/foo", rh)

	timeHandler := handlers.MakeRateLimitHandler(
		handlers.NewTimeHandlerFunc(time.RFC1123),
		0.05,
		1,
	)
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/time/rfc1123", handlers.NewTimeHandlerFunc(time.RFC1123))
	mux.HandleFunc("/time/rfc3339", handlers.NewTimeHandlerFunc(time.RFC3339))

	log.Println("Listening...")

	log.Fatal(http.ListenAndServe(":3000", mux))
}
