package handlers

import (
	"net/http"
)
import "golang.org/x/time/rate"

// This is a demo application following the tutorial on
// https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
// also refer: https://www.zhihu.com/question/265433666
// try to use manual dependency injection

type PricingPlanService interface {
	ResolvePricingPlan(apikey string) *rate.Limiter
}

// MakeRateLimitHandler make a rate limited handler
// limit: rps, b: max value of token
// default hardcoded three type of limitation, select by apikey header
func MakeRateLimitHandler(next http.HandlerFunc, s PricingPlanService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apikey := r.Header.Get("X-api-key")
		// bucket, a global object, maintaining info across request
		bucket := s.ResolvePricingPlan(apikey)

		if bucket.Allow() {
			next(w, r)
		} else {
			w.WriteHeader(429)
			w.Write([]byte("You have exhausted your API Request Quota"))
		}
	}
}
