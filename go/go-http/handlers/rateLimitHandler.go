package handlers

import "net/http"
import "golang.org/x/time/rate"

// MakeRateLimitHandler make a rate limited handler
// limit: rps, b: max value of token
func MakeRateLimitHandler(next http.HandlerFunc, limit rate.Limit, b int) http.HandlerFunc {
	bucket := rate.NewLimiter(limit, b)
	return func(w http.ResponseWriter, r *http.Request) {
		if bucket.Allow() {
			next(w, r)
		} else {
			w.WriteHeader(429)
			w.Write([]byte("You have exhausted your API Request Quota"))
		}
	}
}
