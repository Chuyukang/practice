package main

// a demo follow the tutorial: https://studygolang.com/articles/9467

import (
	"fmt"
	"go-http/handlers"
	"go-http/service"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/foo", rh)

	// require PricingPlanService get a bucket for rate limiting
	var s handlers.PricingPlanService
	s = service.NewPricingPlanService()

	timeHandler := handlers.MakeRateLimitHandler(
		handlers.NewTimeHandlerFunc(time.RFC1123),
		s,// wire service
	)
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/time/rfc1123", handlers.NewTimeHandlerFunc(time.RFC1123))
	mux.HandleFunc("/time/rfc3339", handlers.NewTimeHandlerFunc(time.RFC3339))

	log.Println("Listening...")

	log.Fatal(http.ListenAndServe(":3000", mux))
}
