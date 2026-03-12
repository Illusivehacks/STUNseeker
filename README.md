derrickseeker

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)](https://golang.org)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/yourusername/stunseeker)

> A high-performance STUN server discovery and NAT traversal analysis tool

Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Logic & Architecture](#logic--architecture)
- [Usage Examples](#usage-examples)
- [Configuration](#configuration)
- [API Reference](#api-reference)
- [Contributing](#contributing)
- [License](#license)

---

Introduction

StunSeeker is a powerful command-line tool and library designed to discover, test, and analyze STUN (Session Traversal Utilities for NAT) servers. It helps developers and network administrators:

- 🔍 Discover public STUN servers
- 🧪 Test STUN server responsiveness and latency
- 🌐 Determine NAT type and behavior
- 📊 Analyze network traversal capabilities
- 🔧 Debug WebRTC and VoIP connectivity issues

What is STUN?

STUN (Session Traversal Utilities for NAT) is a standardized protocol ([RFC 5389](https://tools.ietf.org/html/rfc5389)) that enables devices behind NAT (Network Address Translation) to discover their public IP address and port mapping. This is essential for peer-to-peer communication in applications like:

- WebRTC video/audio calls
- VoIP (Voice over IP) systems
- Online gaming
- P2P file sharing
- Real-time collaboration tools

---

Features

Feature	Description	
Server Discovery	Automatically discovers public STUN servers from multiple sources	
Latency Testing	Measures round-trip time (RTT) to STUN servers	
NAT Type Detection	Identifies Full Cone, Restricted Cone, Port Restricted, and Symmetric NAT	
Parallel Scanning	High-performance concurrent server testing	
JSON Output	Machine-readable output for integration with CI/CD pipelines	
Interactive Mode	Real-time visualization of STUN tests	
Custom Servers	Test private or specific STUN server endpoints	
Cross-Platform	Works on Linux, macOS, and Windows	

---

Installation

Using Go

```bash
go install github.com/yourusername/stunseeker@latest
```

Using Homebrew (macOS/Linux)

```bash
brew tap yourusername/stunseeker
brew install stunseeker
```

From Source

```bash
git clone https://github.com/yourusername/stunseeker.git
cd stunseeker
go build -o stunseeker cmd/stunseeker/main.go
```

Docker

```bash
docker pull yourusername/stunseeker:latest
docker run --rm yourusername/stunseeker --help
```

---

Quick Start

Basic Usage

```bash
# Discover and test all known public STUN servers
stunseeker discover

# Test a specific STUN server
stunseeker test stun.l.google.com:19302

# Detect your NAT type
stunseeker nat-type

# Find the fastest STUN server
stunseeker discover --fastest --limit 5
```

Smoke Test / Quick Verification

```bash
# Build & Run (Windows / Linux / macOS)
go build -o stunseeker cmd/stunseeker/main.go
./stunseeker      # Linux/macOS
.\stunseeker.exe  # Windows PowerShell
```

Erwartete Ausgabe (Minimal-Stub)

```
Top 5 STUN Servers:
    stun.l.google.com:19302 - 12 latency
    stun1.l.google.com:19302 - 15 latency

NAT Type: Full Cone NAT
Public Endpoint: 203.0.113.45:52413
```

Example Output

```
╔══════════════════════════════════════════════════════════════╗
║                    StunSeeker v1.0.0                         ║
║           STUN Server Discovery & Analysis Tool              ║
╚══════════════════════════════════════════════════════════════╝

[✓] Discovered 47 STUN servers
[✓] Testing server responsiveness...

┌─────────────────────────────────────────────────────────────┐
│ Top 5 Fastest STUN Servers                                  │
├──────────────────────────────┬──────────┬─────────┬─────────┤
│ Server                       │ Latency  │ Status  │ NAT     │
├──────────────────────────────┼──────────┼─────────┼─────────┤
│ stun.l.google.com:19302      │ 12ms     │ ✓ OK    │ Full    │
│ stun1.l.google.com:19302     │ 15ms     │ ✓ OK    │ Full    │
│ stun.voipbuster.com          │ 28ms     │ ✓ OK    │ Restrict│
│ stun.ekiga.net               │ 34ms     │ ✓ OK    │ Full    │
│ stun.ideasip.com             │ 41ms     │ ✓ OK    │ Full    │
└──────────────────────────────┴──────────┴─────────┴─────────┘

NAT Type: Full Cone NAT ✓
Public IP: 203.0.113.45:52413
```

---

Logic & Architecture

Core Logic Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                        StunSeeker Core                          │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐      │
│  │   Discovery  │───▶│   Testing    │───▶│   Analysis   │      │
│  │   Engine     │    │   Engine     │    │   Engine     │      │
│  └──────────────┘    └──────────────┘    └──────────────┘      │
│         │                   │                   │               │
│         ▼                   ▼                   ▼               │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐      │
│  │ STUN Server  │    │ UDP/TCP      │    │ NAT Type     │      │
│  │ Registry     │    │ Probes       │    │ Classifier   │      │
│  └──────────────┘    └──────────────┘    └──────────────┘      │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

STUN Protocol Implementation

```go
// pkg/stun/message.go

// STUN Message Structure (RFC 5389)
//  0                   1                   2                   3
//  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |0 0|     STUN Message Type     |         Message Length        |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                         Magic Cookie                          |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                                                               |
// |                     Transaction ID (96 bits)                  |
// |                                                               |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

const (
    // STUN Magic Cookie (RFC 5389)
    MagicCookie = 0x2112A442
    
    // Message Types
    BindingRequest         = 0x0001
    BindingResponse        = 0x0101
    BindingErrorResponse   = 0x0111
)

// Attribute Types
const (
    MappedAddress     = 0x0001
    ResponseAddress   = 0x0002
    ChangeRequest     = 0x0003
    SourceAddress     = 0x0004
    ChangedAddress    = 0x0005
    Username          = 0x0006
    Password          = 0x0007
    MessageIntegrity  = 0x0008
    ErrorCode         = 0x0009
    UnknownAttributes = 0x000A
    ReflectedFrom     = 0x000B
    XorMappedAddress  = 0x0020
)
```

NAT Type Detection Algorithm

```go
// pkg/nat/detector.go

// NATType represents the type of NAT detected
type NATType int

const (
    NATUnknown NATType = iota
    NATOpenInternet
    NATFullCone
    NATRestricted
    NATPortRestricted
    NATSymmetric
    NATSymmetricUDPFirewall
    NATBlocked
)

// DetectNATType performs the RFC 5780 NAT type detection tests
func (d *Detector) DetectNATType(ctx context.Context, serverAddr string) (NATType, *Result, error) {
    // Test I: Check if we can get a binding response
    // Send Binding Request to primary server
    resp1, err := d.sendBindingRequest(ctx, serverAddr, serverAddr)
    if err != nil {
        return NATBlocked, nil, fmt.Errorf("no response from server: %w", err)
    }
    
    mappedAddr1 := resp1.GetXorMappedAddress()
    localAddr := d.conn.LocalAddr().(*net.UDPAddr)
    
    // If mapped address equals local address, we're not behind NAT
    if mappedAddr1.IP.Equal(localAddr.IP) && mappedAddr1.Port == localAddr.Port {
        // Test II: Check for UDP firewall
        resp2, err := d.sendBindingRequestWithChangeIP(ctx, serverAddr, true, false)
        if err != nil {
            return NATSymmetricUDPFirewall, nil, nil
        }
        _ = resp2
        return NATOpenInternet, nil, nil
    }
    
    // Test II: Check if NAT is Full Cone
    // Send request to primary server asking for response from different IP
    resp2, err := d.sendBindingRequestWithChangeIP(ctx, serverAddr, true, true)
    if err == nil && resp2 != nil {
        return NATFullCone, &Result{
            PublicIP:   mappedAddr1.IP,
            PublicPort: mappedAddr1.Port,
        }, nil
    }
    
    // Test III: Check if NAT is Symmetric
    // Send request to secondary server
    secondaryAddr := d.getSecondaryAddress(serverAddr)
    resp3, err := d.sendBindingRequest(ctx, secondaryAddr, secondaryAddr)
    if err != nil {
        return NATUnknown, nil, err
    }
    
    mappedAddr2 := resp3.GetXorMappedAddress()
    
    // If mapped addresses differ, it's Symmetric NAT
    if !mappedAddr1.IP.Equal(mappedAddr2.IP) || mappedAddr1.Port != mappedAddr2.Port {
        return NATSymmetric, &Result{
            PublicIP:   mappedAddr1.IP,
            PublicPort: mappedAddr1.Port,
        }, nil
    }
    
    // Test IV: Check if NAT is Port Restricted
    // Send request asking for response from same IP but different port
    resp4, err := d.sendBindingRequestWithChangeIP(ctx, serverAddr, false, true)
    if err != nil {
        return NATPortRestricted, &Result{
            PublicIP:   mappedAddr1.IP,
            PublicPort: mappedAddr1.Port,
        }, nil
    }
    _ = resp4
    
    return NATRestricted, &Result{
        PublicIP:   mappedAddr1.IP,
        PublicPort: mappedAddr1.Port,
    }, nil
}
```

Server Discovery Logic

```go
// pkg/discovery/discovery.go

// DiscoveryEngine finds and validates STUN servers
type DiscoveryEngine struct {
    sources    []Source
    tester     *testing.Engine
    timeout    time.Duration
    workers    int
}

// Source represents a STUN server source
type Source interface {
    Fetch(ctx context.Context) ([]Server, error)
    Name() string
}

// Discover finds and tests STUN servers from all sources
func (e *DiscoveryEngine) Discover(ctx context.Context, opts DiscoverOptions) ([]ServerResult, error) {
    // Fetch servers from all sources concurrently
    serverChan := make(chan Server, 100)
    var wg sync.WaitGroup
    
    for _, source := range e.sources {
        wg.Add(1)
        go func(src Source) {
            defer wg.Done()
            servers, err := src.Fetch(ctx)
            if err != nil {
                log.Printf("Failed to fetch from %s: %v", src.Name(), err)
                return
            }
            for _, s := range servers {
                select {
                case serverChan <- s:
                case <-ctx.Done():
                    return
                }
            }
        }(source)
    }
    
    // Close channel when all sources complete
    go func() {
        wg.Wait()
        close(serverChan)
    }()
    
    // Test servers concurrently
    resultChan := make(chan ServerResult, 100)
    var testWg sync.WaitGroup
    
    // Start worker pool
    for i := 0; i < e.workers; i++ {
        testWg.Add(1)
        go func() {
            defer testWg.Done()
            for server := range serverChan {
                result := e.tester.Test(ctx, server)
                select {
                case resultChan <- result:
                case <-ctx.Done():
                    return
                }
            }
        }()
    }
    
    // Close results when testing completes
    go func() {
        testWg.Wait()
        close(resultChan)
    }()
    
    // Collect results
    var results []ServerResult
    for result := range resultChan {
        if opts.OnlyResponsive && !result.Responsive {
            continue
        }
        results = append(results, result)
    }
    
    // Sort by latency
    sort.Slice(results, func(i, j int) bool {
        return results[i].Latency < results[j].Latency
    })
    
    if opts.Limit > 0 && len(results) > opts.Limit {
        results = results[:opts.Limit]
    }
    
    return results, nil
}
```

Concurrent Testing Engine

```go
// pkg/testing/engine.go

// Engine tests STUN server connectivity and performance
type Engine struct {
    timeout     time.Duration
    retryCount  int
    retryDelay  time.Duration
}

// Test performs comprehensive testing of a STUN server
func (e *Engine) Test(ctx context.Context, server Server) ServerResult {
    result := ServerResult{
        Server:   server,
        TestedAt: time.Now(),
    }
    
    // Resolve address
    addr, err := net.ResolveUDPAddr("udp", server.Address)
    if err != nil {
        result.Error = fmt.Errorf("resolve failed: %w", err)
        return result
    }
    
    // Create UDP connection
    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        result.Error = fmt.Errorf("dial failed: %w", err)
        return result
    }
    defer conn.Close()
    
    // Set timeout
    deadline, ok := ctx.Deadline()
    if !ok {
        deadline = time.Now().Add(e.timeout)
    }
    conn.SetDeadline(deadline)
    
    // Build STUN binding request
    req := stun.NewMessage(stun.BindingRequest)
    req.AddAttribute(stun.AttrSoftware, []byte("StunSeeker/1.0"))
    req.Encode()
    
    // Measure latency with multiple samples
    var latencies []time.Duration
    for i := 0; i < e.retryCount; i++ {
        start := time.Now()
        
        // Send request
        _, err = conn.Write(req.Bytes())
        if err != nil {
            result.Error = fmt.Errorf("send failed: %w", err)
            return result
        }
        
        // Receive response
        buf := make([]byte, 1500)
        n, err := conn.Read(buf)
        if err != nil {
            if i < e.retryCount-1 {
                time.Sleep(e.retryDelay)
                continue
            }
            result.Error = fmt.Errorf("receive failed: %w", err)
            return result
        }
        
        latency := time.Since(start)
        latencies = append(latencies, latency)
        
        // Parse response
        resp := &stun.Message{}
        if err := resp.Decode(buf[:n]); err != nil {
            result.Error = fmt.Errorf("parse failed: %w", err)
            return result
        }
        
        // Verify transaction ID matches
        if !bytes.Equal(req.TransactionID, resp.TransactionID) {
            result.Error = fmt.Errorf("transaction ID mismatch")
            return result
        }
        
        // Extract mapped address
        xorAddr, err := resp.GetXorMappedAddress()
        if err != nil {
            // Fallback to MAPPED-ADDRESS
            mappedAddr, err := resp.GetMappedAddress()
            if err != nil {
                result.Error = fmt.Errorf("no mapped address: %w", err)
                return result
            }
            result.MappedAddress = mappedAddr
        } else {
            result.MappedAddress = xorAddr
        }
        
        time.Sleep(e.retryDelay)
    }
    
    // Calculate average latency
    var total time.Duration
    for _, l := range latencies {
        total += l
    }
    result.Latency = total / time.Duration(len(latencies))
    result.Responsive = true
    
    return result
}
```

---

Usage Examples

Command-Line Interface

```bash
# Discover all public STUN servers
stunseeker discover

# Find fastest servers with JSON output
stunseeker discover --format json --limit 10

# Test specific servers
stunseeker test stun.l.google.com:19302 stun1.l.google.com:19302

# Detect NAT type with verbose output
stunseeker nat-type --verbose

# Continuous monitoring mode
stunseeker monitor --interval 30s --servers stun.l.google.com:19302

# Export results to file
stunseeker discover --output stun-servers.json
```

Library Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/yourusername/stunseeker/pkg/discovery"
    "github.com/yourusername/stunseeker/pkg/nat"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // Create discovery engine
    engine := discovery.NewEngine(
        discovery.WithTimeout(5*time.Second),
        discovery.WithWorkers(10),
    )
    
    // Add default sources
    engine.AddSource(discovery.NewPublicServerSource())
    engine.AddSource(discovery.NewDNSSource("stun.%s"))
    
    // Discover servers
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
    
    // Detect NAT type
    detector := nat.NewDetector()
    natType, result, err := detector.DetectNATType(ctx, "stun.l.google.com:19302")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("\nNAT Type: %s\n", natType)
    fmt.Printf("Public Endpoint: %s:%d\n", result.PublicIP, result.PublicPort)
}
```

---

Configuration

Configuration File

```yaml
# ~/.stunseeker/config.yaml

# Default timeout for STUN requests
timeout: 5s

# Number of retry attempts
retries: 3

# Number of concurrent workers
workers: 20

# Default output format (table, json, csv)
format: table

# Custom STUN servers
servers:
  - address: stun.custom.com:3478
    region: us-east
    priority: 1
  - address: stun.backup.com:3478
    region: eu-west
    priority: 2

# Discovery sources
sources:
  public_registry: true
  dns_discovery: true
  custom_list: /path/to/custom-servers.txt

# Monitoring settings
monitor:
  enabled: false
  interval: 60s
  alert_threshold: 100ms
```

Environment Variables

Variable	Description	Default	
`STUNSEEKER_TIMEOUT`	Request timeout	`5s`	
`STUNSEEKER_WORKERS`	Concurrent workers	`20`	
`STUNSEEKER_FORMAT`	Output format	`table`	
`STUNSEEKER_CONFIG`	Config file path	`~/.stunseeker/config.yaml`	

---

API Reference

REST API (Server Mode)

```bash
# Start API server
stunseeker server --port 8080
```

Endpoints

Method	Endpoint	Description	
`GET`	`/health`	Health check	
`GET`	`/api/v1/servers`	List discovered servers	
`POST`	`/api/v1/test`	Test specific server	
`GET`	`/api/v1/nat-type`	Get NAT type	
`GET`	`/api/v1/metrics`	Prometheus metrics	

Example API Request

```bash
# Test a STUN server via API
curl -X POST http://localhost:8080/api/v1/test \
  -H "Content-Type: application/json" \
  -d '{"server": "stun.l.google.com:19302"}'
```

Response:

```json
{
  "server": "stun.l.google.com:19302",
  "responsive": true,
  "latency_ms": 12,
  "mapped_address": "203.0.113.45:52413",
  "tested_at": "2024-01-15T10:30:00Z"
}
```

---

NAT Types Explained

NAT Type	Description	P2P Compatibility	
Open Internet	No NAT, public IP	✓ Excellent	
Full Cone	Any external host can send packets	✓ Excellent	
Restricted Cone	Must first send packet to external host	✓ Good	
Port Restricted	Must first send packet to specific IP:port	△ Moderate	
Symmetric	Different mapping for each destination	✗ Poor (needs TURN)	
UDP Blocked	Firewall blocks UDP	✗ Requires TURN	

---

Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

Development Setup

```bash
# Clone repository
git clone https://github.com/yourusername/stunseeker.git
cd stunseeker

# Install dependencies
go mod download

# Run tests
go test ./...

# Run linter
golangci-lint run

# Build locally
go build -o stunseeker cmd/stunseeker/main.go
```

---

License

MIT License - see [LICENSE](LICENSE) file for details.

---

Acknowledgments

- [RFC 5389](https://tools.ietf.org/html/rfc5389) - STUN Protocol
- [RFC 5780](https://tools.ietf.org/html/rfc5780) - NAT Behavior Discovery
- [RFC 7064](https://tools.ietf.org/html/rfc7064) - STUN URI
- [pion/stun](https://github.com/pion/stun) - Go STUN library inspiration

---

Related Projects

- [pion/stun](https://github.com/pion/stun) - Go STUN library
- [coturn](https://github.com/coturn/coturn) - TURN/STUN server
- [stunserver](https://github.com/jselbie/stunserver) - STUN server implementation

---
