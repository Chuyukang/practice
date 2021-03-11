package service

import (
	"golang.org/x/time/rate"
	"strings"
	"sync"
	"time"
)

// duck type? no need to declare or import interface
// import "go-http/handlers"

// a demo follow the tutorial on https://www.alexedwards.net/blog/how-to-rate-limit-http-requests

type cacheLine struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type PricingPlanServiceImpl struct {
	// the reason why not use sync.Mutex is
	// the latter is optimized for certain usage, which is not compatible
	mu    sync.Mutex
	cache map[string]*cacheLine
}

func NewPricingPlanService() *PricingPlanServiceImpl {
	s := PricingPlanServiceImpl{cache: make(map[string]*cacheLine)}
	go s.cleanCacheLines()
	return &s
}

func (s *PricingPlanServiceImpl) ResolvePricingPlan(apikey string) *rate.Limiter {
	s.mu.Lock()
	defer s.mu.Unlock()

	// final result
	var limiter *rate.Limiter

	c, exists := s.cache[apikey]
	tick := time.Now()
	if !exists { // add new entry
		switch {
		case strings.HasPrefix(apikey, "PX001-"):
			limiter = rate.NewLimiter(0.025, 1)
		case strings.HasPrefix(apikey, "BX001-"):
			limiter = rate.NewLimiter(0.05, 1)
		default:
			limiter = rate.NewLimiter(0.02, 1)
		}
		s.cache[apikey] = &cacheLine{limiter, tick}
	} else { // update entry
		limiter = c.limiter
		s.cache[apikey].lastSeen = tick
	}

	return limiter
}

// cleanCacheLine periodically remove unused lines
// the sleep time and timeout should be set appropriately
func (s *PricingPlanServiceImpl) cleanCacheLines() {
	for {
		time.Sleep(time.Minute)

		s.mu.Lock()
		for k, c := range s.cache {
			if time.Since(c.lastSeen) > 3*time.Minute {
				delete(s.cache, k)
			}
		}
		s.mu.Unlock()
	}
}
