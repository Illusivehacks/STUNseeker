package discovery

import (
	"context"
	"time"
)

type Engine struct{}

func NewEngine(options ...interface{}) *Engine {
	return &Engine{}
}

type ServerResult struct {
	Server  struct{ Address string }
	Latency int
}

type DiscoverOptions struct {
	OnlyResponsive bool
	Limit          int
}

func (e *Engine) AddSource(source interface{}) {}

func (e *Engine) Discover(ctx context.Context, opts DiscoverOptions) ([]ServerResult, error) {
	return []ServerResult{
		{Server: struct{ Address string }{Address: "stun.l.google.com:19302"}, Latency: 12},
		{Server: struct{ Address string }{Address: "stun1.l.google.com:19302"}, Latency: 15},
	}, nil
}

func WithTimeout(timeout time.Duration) interface{} {
	return timeout
}

func WithWorkers(workers int) interface{} {
	return workers
}

func NewPublicServerSource() interface{} {
	return struct{}{}
}
