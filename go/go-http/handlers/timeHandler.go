package handlers

import (
	"net/http"
	"time"
)

func NewTimeHandlerFunc(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chineseLoc := time.FixedZone("CST", +8*60*60)
		tm := time.Now().In(chineseLoc).Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}
