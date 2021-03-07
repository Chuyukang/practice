package main

// a demo follow the tutorial: https://studygolang.com/articles/9467

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func newTimeHandlerFunc(format string) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		chineseLoc := time.FixedZone("CST", +8*60*60)
		tm := time.Now().In(chineseLoc).Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}


func main() {
	fmt.Println("Hello World!")

	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/foo", rh)
	
	mux.HandleFunc("/time", newTimeHandlerFunc(time.RFC1123))
	mux.HandleFunc("/time/rfc1123", newTimeHandlerFunc(time.RFC1123))
	mux.HandleFunc("/time/rfc3339", newTimeHandlerFunc(time.RFC3339))

	log.Println("Listening...")
	
	http.ListenAndServe(":3000", mux)
}