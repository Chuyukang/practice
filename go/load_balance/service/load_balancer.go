package service

import (
	"math/rand"
	"sync"
)

// LoadBalancer a LoadBalancer support multi policy
type LoadBalancer interface {
	GetBackend() string
}

type UpstreamFetcher interface {
	FetchUpstream() []string
}

// RoundRobinLB RodRobin policy LB
type RoundRobinLB struct {
	lastTarget int
	fetcher    UpstreamFetcher
	mu         sync.Mutex
}

// NewRoundRobinLB construct a RoundRobinLB object
func NewRoundRobinLB(fetcher UpstreamFetcher) LoadBalancer {
	lb := RoundRobinLB{lastTarget: -1, fetcher: fetcher}
	return &lb
}

// GetBackend select a backend from upstreams
func (lb *RoundRobinLB) GetBackend() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	upstreams := lb.fetcher.FetchUpstream()

	n := len(upstreams)
	target := (lb.lastTarget + 1) % n
	lb.lastTarget = target

	return upstreams[target]
}

type RandomLB struct {
	fetcher UpstreamFetcher
	mu      sync.Mutex
}

func NewRandomLB(fetcher UpstreamFetcher) LoadBalancer {
	return &RandomLB{fetcher: fetcher}
}

func (lb *RandomLB) GetBackend() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	upstreams := lb.fetcher.FetchUpstream()
	n := len(upstreams)

	target := rand.Intn(n)

	return upstreams[target]
}
