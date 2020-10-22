package client

import "context"

// Producer is an example interface which we'll use for mock generation
type Producer interface {
	Produce(ctx context.Context, objects []string) error
}
