package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/s2a0v2a4/LocationWACall/pkg/discovery"
	"github.com/s2a0v2a4/LocationWACall/pkg/nat"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	engine := discovery.NewEngine(
		discovery.WithTimeout(5*time.Second),
		discovery.WithWorkers(10),
	)
	engine.AddSource(discovery.NewPublicServerSource())

	results, err := engine.Discover(ctx, discovery.DiscoverOptions{
		OnlyResponsive: true,
		Limit:          5,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Top 5 STUN Servers:")
	for _, r := range results {
		fmt.Printf("  %s - %v latency\n", r.Server.Address, r.Latency)
	}

	detector := nat.NewDetector()
	natType, result, err := detector.DetectNATType(ctx, "stun.l.google.com:19302")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nNAT Type: %s\n", natType)
	fmt.Printf("Public Endpoint: %s:%d\n", result.PublicIP, result.PublicPort)
}
