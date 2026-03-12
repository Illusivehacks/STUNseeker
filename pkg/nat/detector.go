package nat

import "context"

type Detector struct{}

func NewDetector() *Detector {
	return &Detector{}
}

type NATType string
type Result struct {
	PublicIP   string
	PublicPort int
}

func (d *Detector) DetectNATType(ctx context.Context, serverAddr string) (NATType, *Result, error) {
	return "Full Cone NAT", &Result{
		PublicIP:   "203.0.113.45",
		PublicPort: 52413,
	}, nil
}
